package user

import (
	"net/http"
	"server/middleware"
	"server/util"

	"server/model"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	sessionUserId = "user-id"
)

func (userManager UserEndpoint) GetCurrentUserFromSession(c *gin.Context) *model.User {
	var user model.User

	sess := sessions.Default(c)
	userId := sess.Get(sessionUserId)
	if userId != nil {
		uid := userId.(uint)
		result := userManager.db.First(&user, uid)
		if result.Error != nil {
			return nil
		}
		return &user
	}

	return nil
}

func (userManager UserEndpoint) SetCurrentUser(c *gin.Context, user *model.User) {
	sess := sessions.Default(c)

	switch {
	case user != nil:
		sess.Set(sessionUserId, user.ID)
		logger.Printf("%#v", c)
	default:
		sess.Delete(sessionUserId)
	}

	err := sess.Save()
	if err != nil {
		logger.Println("set-current-user:", err)
	}
}

func renderLogin(c *gin.Context, status int) {
	util.HtmlWithContext(c, status, "login", gin.H{
		"title": "Login",
	})
}

func (userManager UserEndpoint) addLoginEndpoints(router gin.IRouter) {
	router.GET("/login", func(c *gin.Context) {
		user := middleware.CurrentUser(c)

		if user == nil {
			renderLogin(c, http.StatusOK)
			return
		}

		c.Redirect(http.StatusTemporaryRedirect, "/")
	})

	router.POST("/login", func(c *gin.Context) {
		user := middleware.CurrentUser(c)

		switch {
		case user != nil:
			c.Redirect(http.StatusSeeOther, "/")
			return
		}

		var provided struct {
			Username string `form:"username" binding:"required"`
			Password string `form:"password" binding:"required"`
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

		userManager.SetCurrentUser(c, user)

		c.Redirect(http.StatusSeeOther, "/")
	})

	router.POST("/logout", func(c *gin.Context) {
		userManager.SetCurrentUser(c, nil)
		c.Redirect(http.StatusSeeOther, "/")
	})
}

func (userManager UserEndpoint) Login(username, password string) (*model.User, error) {
	var user model.User
	result := userManager.db.Model(&user).Where("username = ?", username).First(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, model.CheckPasswordsMatch(password, user.PasswordHash)
}
