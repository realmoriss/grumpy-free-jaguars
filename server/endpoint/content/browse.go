package content

import (
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/util"
)

func renderBrowse(c *gin.Context, db *gorm.DB, status int) {
	var content []model.CaffPreview

	dbRes := db.Model(&model.CaffContent{}).Preload("User").Find(&content)

	util.HtmlWithContext(c, status, "browse", gin.H{
		"title":   "Browse",
		"valid":   dbRes.Error == nil,
		"content": content,
	})
}

func (contentManager ContentEndpoint) addBrowseEndpoints(router gin.IRouter) {
	router.GET("/browse", func(c *gin.Context) {
		renderBrowse(c, contentManager.db, http.StatusOK)
	})
}
