package main

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/nosurf"
	adapter "github.com/gwatts/gin-adapter"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"server/model"

	"server/endpoint/content"
	"server/endpoint/user"
	"server/middleware"
)

func main() {
	// TODO: Setup db.
	db, err := gorm.Open(sqlite.Open("./data/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	cookieStore := cookie.NewStore([]byte("TODO-take-secret-from-elsewhere"))

	router := gin.Default()

	csrfHandler := func() gin.HandlerFunc {
		next, wrapper := adapter.New()
		nsHandler := nosurf.New(next)
		nsHandler.SetBaseCookie(http.Cookie{
			Path: "/",
			HttpOnly: true,
		})
		nsHandler.SetFailureHandler(http.HandlerFunc((func(rw http.ResponseWriter, req *http.Request) {
			http.Error(rw, "failed to verify CSRF token", http.StatusBadRequest)
		})))

		return wrapper(nsHandler)
	}()

	router.Use(csrfHandler)
	router.Use(sessions.Sessions("login_state", cookieStore))

	router.HTMLRender = ginview.Default()

	userGroup := router.Group("/user")
	userEndpoint := user.NewEndpoint(userGroup, db)
	userFromSession := func(c *gin.Context) *model.User {
		return userEndpoint.GetCurrentUserFromSession(c)
	}
	router.Use(middleware.WithUser(userFromSession))

	authorized := router.Group("/content")

	// avoid import cycle so user endpoint package can defer tasks to middleware

	authorized.Use(middleware.AuthRequired())
	{
		content.NewEndpoint(authorized, db)
	}

	router.GET("/", func(c *gin.Context) {
		currentUser := middleware.CurrentUser(c)
		if currentUser != nil {
			c.Redirect(http.StatusTemporaryRedirect, "/content/browse")
			return
		}

		c.HTML(200, "index", gin.H{
			"title": "Home",
		})
	})

	http.ListenAndServe(":3000", router)
	router.Run(":3000")
}
