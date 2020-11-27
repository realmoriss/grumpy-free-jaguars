package main

import (
	"net/http"

	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/nosurf"
	adapter "github.com/gwatts/gin-adapter"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	content "server/endpoint/content"
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

	router.HTMLRender = ginview.Default()

	userGroup := router.Group("/user")
	userEndpoint := user.NewEndpoint(userGroup, db)

	authorized := router.Group("/content")

	authorized.Use(userEndpoint.AuthRequired())
	{
		content.NewEndpoint(authorized, db)
	}

	router.GET("/", func(c *gin.Context) {
		currentUser := userEndpoint.CurrentUser(c)

		if currentUser != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/content/browse")
		} else {
			c.HTML(200, "index", gin.H{
				"title": "Home",
			})
		}
	})

	http.ListenAndServe(":3000", nosurf.New(router))
	router.Run(":3000")
}
