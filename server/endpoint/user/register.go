package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/nosurf"

	"server/model"
)

func (userManager UserEndpoint) Register(user model.User) error {
	result := userManager.db.Create(&user)
	return result.Error
}

func (userManager UserEndpoint) addRegisterEndpoints(router gin.IRouter) {
	// TODO: continue cleanup
	router.GET("/register", func(c *gin.Context) {
		// CSRF example
		tok := nosurf.Token(c.Request)
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"csrf_token": tok,
		})
	})

	router.POST("/register", func(c *gin.Context) {
		var requested struct {
			Username        string `form:"username"`
			Password        string `form:"password"`
			PasswordConfirm string `form:"password_confirm"`
		}

		fail := func(err error) {
			logger.Println(err)
			msg := "failed to serve request"
			status := http.StatusBadRequest

			switch err {
			case ErrPasswordInsecure:
				fallthrough
			case ErrPasswordDoNotMatch:
				msg = err.Error()
			}

			c.String(status, msg)
			return
		}

		err := c.ShouldBind(&requested)
		if err != nil {
			fail(err)
			return
		}

		switch {
		case requested.Password != requested.PasswordConfirm:
			fail(ErrPasswordDoNotMatch)
			return

		case len(requested.Password) < 1:
			fail(ErrPasswordInsecure)
			return
		}

		err = userManager.Register(model.User{
			Username:     requested.Username,
			PasswordHash: model.HashPassword(requested.Password),
		})
		if err != nil {
			fail(err)
			return
		}

		c.String(http.StatusOK, "Hello, %s!", c.PostForm("username"))
	})
}
