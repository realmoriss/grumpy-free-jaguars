package content

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/nosurf"
)

func renderUpload(c *gin.Context, status int) {
	c.HTML(status, "upload", gin.H{
		"title":      "Upload",
		"csrf_token": nosurf.Token(c.Request),
	})
}

func (contentManager ContentEndpoint) addUploadEndpoints(router gin.IRouter) {
	router.GET("/upload", func(c *gin.Context) {
		renderUpload(c, http.StatusOK)
	})
}
