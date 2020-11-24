package user

import (
    "log"
    "os"
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/nosurf"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "user: ", log.LstdFlags|log.LUTC|log.Lmsgprefix)
}

type UserEndpoint struct {
	db interface{}
}

func NewEndpoint(router gin.IRouter, db interface{}) *UserEndpoint {
	counter := 0

	router.GET("/profile", func(c *gin.Context) {
		// TODO: Replace with actual functionality
		counter++
		c.String(http.StatusOK, "Hello, visitor %d!", counter)
	})

	router.GET("/register", func(c *gin.Context) {
		// CSRF example
		tok := nosurf.Token(c.Request)
		c.HTML(http.StatusOK, "register.tmpl", gin.H{
			"csrf_token": tok,
		})
	})

	router.POST("/register", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, %s!", c.PostForm("username"))
	})

	return &UserEndpoint{
		db: db,
	}
}
