package content

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"server/util"
)

func renderBrowse(c *gin.Context, status int) {
	util.HtmlWithContext(c, status, "browse", gin.H{
		"title": "Browse",
	})
}

func (contentManager ContentEndpoint) addBrowseEndpoints(router gin.IRouter) {
	router.GET("/browse", func(c *gin.Context) {
		renderBrowse(c, http.StatusOK)
	})
}
