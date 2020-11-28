package content

import (
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/util"
)

func renderPreview(c *gin.Context, db *gorm.DB, status int, id string) {
	var content model.CaffPreview

	dbRes := db.Model(&model.CaffContent{}).Preload("User").First(&content, id)
	if dbRes.Error != nil {
		c.String(http.StatusOK, dbRes.Error.Error())
		return
	}

	util.HtmlWithContext(c, status, "preview", gin.H{
		"title": "Preview",
		"valid": dbRes.Error == nil,
		"image": content,
	})
}

func (contentManager ContentEndpoint) addPreviewEndpoints(router gin.IRouter) {
	router.GET("/preview/:id", func(c *gin.Context) {
		renderPreview(c, contentManager.db, http.StatusOK, c.Param("id"))
	})
}
