package user

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/nosurf"

	"server/model"
)

const sessionUserId = "user-id"

func (userManager UserEndpoint) CurrentUser(c *gin.Context) *model.User {
	var user model.User

	sess := sessions.Default(c)
	userId := sess.Get(sessionUserId)
	if userId != nil {
		uid := userId.(uint)
		userManager.db.First(&user, uid)
		return &user
	}

	return nil
}

func (userManager UserEndpoint) SetCurrentUser(c *gin.Context, user model.User) {
	sess := sessions.Default(c)
	sess.Set(sessionUserId, user.ID)
	sess.Save()
}

func renderLogin(c *gin.Context, status int) {
	c.HTML(status, "login", gin.H{
		"title":      "Login",
		"csrf_token": nosurf.Token(c.Request),
	})
}

func (userManager UserEndpoint) addLoginEndpoints(router gin.IRouter) {
	router.GET("/login", func(c *gin.Context) {
		user := userManager.CurrentUser(c)

		if user == nil {
			renderLogin(c, http.StatusOK)
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	router.POST("/login", func(c *gin.Context) {
		user := userManager.CurrentUser(c)

		switch {
		case user != nil:
			c.Redirect(http.StatusTemporaryRedirect, "/")
			return
		}

		var provided struct {
			Username string `form:"username"`
			Password string `form:"password"`
		}

		if err := c.ShouldBind(&provided); err != nil {
			renderLogin(c, http.StatusBadRequest)
			return
		}

		user, err := userManager.Login(provided.Username, provided.Password)

		if err != nil {
			renderLogin(c, http.StatusUnauthorized)
			return
		}

		userManager.SetCurrentUser(c, *user)

		c.Redirect(http.StatusFound, "/")
	})
}

func (userManager UserEndpoint) Login(username, password string) (*model.User, error) {
	var user model.User
	result := userManager.db.Model(&user).Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, model.CheckPasswordsMatch(username, user.PasswordHash)
}

func (userManager UserEndpoint) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := userManager.CurrentUser(c)
		if user == nil {
			c.HTML(http.StatusUnauthorized, "unauth", gin.H{
				"title": "Unauthorized",
			})
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			c.Next()
		}
	}
}
