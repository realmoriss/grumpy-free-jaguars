package main

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/nosurf"
	"github.com/gwatts/gin-adapter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	user "server/endpoint/user"
)

func main() {
	// TODO: Setup db.
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
	  panic("failed to connect database")
	}

	cookieStore := cookie.NewStore([]byte("TODO-take-secret-from-elsewhere"))

	router := gin.Default()
	router.Use(adapter.Wrap(nosurf.NewPure))
	router.Use(sessions.Sessions("login_state", cookieStore))

	router.LoadHTMLGlob("templates/*")

	{
		g := router.Group("/user")
		user.NewEndpoint(g, db)
	}

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{ }) // TODO
	})

	http.ListenAndServe(":3000", nosurf.New(router))
	router.Run(":3000")
}
