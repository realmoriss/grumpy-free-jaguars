package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/nosurf"
    "github.com/gwatts/gin-adapter"

    user "server/endpoint/user"
)

func main() {
    // TODO: Setup db.
    db := struct{}{}

	router := gin.Default()
	router.Use(adapter.Wrap(nosurf.NewPure))

    router.LoadHTMLGlob("templates/*")

    {
        g := router.Group("/user")
        user.NewEndpoint(g, db)
    }
    http.ListenAndServe(":3000", nosurf.New(router))
	router.Run(":3000")
}
