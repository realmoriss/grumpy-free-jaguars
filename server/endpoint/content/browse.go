package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/nosurf"
)

func renderBrowse(c *gin.Context, status int) {
	c.HTML(status, "browse", gin.H{
		"title":      "Browse",
		"csrf_token": nosurf.Token(c.Request),
	})
}

func (contentManager ContentEndpoint) addBrowseEndpoints(router gin.IRouter) {
	router.GET("/browse", func(c *gin.Context) {
		renderBrowse(c, http.StatusOK)
	})
}
