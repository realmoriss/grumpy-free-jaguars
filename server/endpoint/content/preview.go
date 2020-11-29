package content

import (
	"fmt"
	"net/http"
	"server/middleware"
	"server/model"
	"strconv"

	"github.com/gosimple/slug"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"server/util"
)

func renderPreview(c *gin.Context, db *gorm.DB, status int, id string) {
	var content model.CaffPreview

	contentResult := db.Model(&model.CaffContent{}).Preload("User").First(&content, id)
	if contentResult.Error != nil {
		c.String(http.StatusNotFound, contentResult.Error.Error())
		return
	}

	var comments []model.Comment
	commentsResult := db.Model(&model.Comment{}).Where("caff_content_id = ?", id).Order("updated_at desc").Preload("User").Find(&comments)
	if commentsResult != nil {

	}

	util.HtmlWithContext(c, status, "preview", gin.H{
		"title":    "Preview",
		"valid":    contentResult.Error == nil,
		"image":    content,
		"comments": comments,
	})
}

func (contentManager ContentEndpoint) addPreviewEndpoints(router gin.IRouter) {
	router.GET("/preview/:id", func(c *gin.Context) {
		renderPreview(c, contentManager.db, http.StatusOK, c.Param("id"))
	})

	router.GET("/preview/:id/:caff", func(c *gin.Context) {
		type CaffTitle struct {
			Title string
		}

		var title CaffTitle

		contentResult := contentManager.db.Model(&model.CaffContent{}).Where("id = ?", c.Param("id")).Scan(&title)
		if contentResult.Error != nil {
			c.String(http.StatusNotFound, contentResult.Error.Error())
			return
		}

		titleSlug := slug.Make(title.Title)
		actualURL := fmt.Sprintf("/content/preview/%s/%s.caff", c.Param("id"), titleSlug)
		if c.Request.URL.Path != actualURL {
			c.Redirect(http.StatusFound, actualURL)
			return
		}

		var content model.CaffContent

		contentResult = contentManager.db.Model(&model.CaffContent{}).First(&content, c.Param("id"))
		if contentResult.Error != nil {
			c.String(http.StatusNotFound, contentResult.Error.Error())
			return
		}

		c.Data(http.StatusOK, "application/octet-stream", content.RawFile)
	})

	router.POST("/preview/:id", func(c *gin.Context) {
		caffID := c.Param("id")
		caffIDNum, err := strconv.ParseUint(caffID, 10, 32)
		if err != nil {
			renderPreview(c, contentManager.db, http.StatusBadRequest, caffID)
		}

		var provided struct {
			Comment string `form:"comment" binding:"required"`
		}

		if err := c.ShouldBind(&provided); err != nil {
			renderPreview(c, contentManager.db, http.StatusBadRequest, caffID)
			return
		}

		comment := model.Comment{
			Text:          provided.Comment,
			CaffContentID: uint(caffIDNum),
			UserID:        middleware.CurrentUser(c).ID,
		}

		result := contentManager.db.Create(&comment)
		if result.Error != nil {
			renderPreview(c, contentManager.db, http.StatusInternalServerError, caffID)
			return
		}

		c.Redirect(http.StatusSeeOther, c.Request.RequestURI)
	})
}
