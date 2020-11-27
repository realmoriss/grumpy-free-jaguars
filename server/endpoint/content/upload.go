package content

import (
	"net/http"
	"server/util"

	"github.com/gin-gonic/gin"
)

func renderUpload(c *gin.Context, status int) {
	util.HtmlWithContext(c, status, "upload", gin.H{
		"title": "Upload",
	})
}

func (contentManager ContentEndpoint) addUploadEndpoints(router gin.IRouter) {
	router.GET("/upload", func(c *gin.Context) {
		renderUpload(c, http.StatusOK)
	})
}
