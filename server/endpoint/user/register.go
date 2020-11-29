package user

import (
	"net/http"
	"server/util"

	"server/model"

	"github.com/gin-gonic/gin"
)

func (userManager UserEndpoint) Register(user model.User) error {
	result := userManager.db.Create(&user)
	return result.Error
}

func renderRegister(c *gin.Context, status int, err error) {
	util.HtmlWithContext(c, status, "register", gin.H{
		"title": "Register",
		"error": err,
	})
}

func (userManager UserEndpoint) addRegisterEndpoints(router gin.IRouter) {
	// TODO: continue cleanup
	router.GET("/register", func(c *gin.Context) {
		// CSRF example
		renderRegister(c, http.StatusOK, nil)
	})

	router.POST("/register", func(c *gin.Context) {
		var requested struct {
			Username        string `form:"username" binding:"required"`
			Password        string `form:"password" binding:"required"`
			PasswordConfirm string `form:"password_confirm" binding:"required"`
		}

		fail := func(err error) {
			renderRegister(c, http.StatusBadRequest, err)
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

		case len(requested.Password) < 8:
			fail(ErrPasswordInsecure)
			return
		}

		err = userManager.Register(model.User{
			Username:     requested.Username,
			PasswordHash: model.HashPassword(requested.Password),
		})
		if err != nil {
			fail(ErrRegistrationUnsuccessful)
			return
		}

		c.Redirect(http.StatusSeeOther, "/")
	})
}
