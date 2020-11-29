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

	router.GET("/preview/:id/download", func(c *gin.Context) {
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
		if len(titleSlug) == 0 {
			titleSlug = c.Param("id")
		}

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

	router.POST("/delete/caff/:id", func(c *gin.Context) {
		caffID := c.Param("id")
		user := middleware.CurrentUser(c)

		switch {
		case user == nil:
			// should never happen due to the auth middleware, but let's be sure we do not dereference any nils.
			c.Redirect(http.StatusSeeOther, "/")
			return
		case model.IsAdministrator(*user) == false:
			util.HtmlWithContext(c, http.StatusUnauthorized, "unauth", gin.H{})
		}

		caffIDNum, err := strconv.ParseUint(caffID, 10, 32)
		if err != nil {
			renderPreview(c, contentManager.db, http.StatusBadRequest, caffID)
		}

		var pic model.CaffContent

		result := contentManager.db.Find(&pic, caffIDNum)
		if result.Error != nil {
			renderPreview(c, contentManager.db, http.StatusInternalServerError, caffID)
			return
		}

		contentManager.db.Delete(pic)

		c.Redirect(http.StatusSeeOther, "/content/browse")
	})

	router.POST("/delete/comment/:cid", func(c *gin.Context) {
		caffID := c.Param("id")
		commentID := c.Param("cid")

		user := middleware.CurrentUser(c)

		switch {
		case user == nil:
			// should never happen due to the auth middleware, but let's be sure we do not dereference any nils.
			c.Redirect(http.StatusSeeOther, "/")
			return
		case model.IsAdministrator(*user) == false:
			util.HtmlWithContext(c, http.StatusUnauthorized, "unauth", gin.H{})
		}

		commentIDNum, err := strconv.ParseUint(commentID, 10, 32)
		if err != nil {
			renderPreview(c, contentManager.db, http.StatusBadRequest, caffID)
		}

		var comment model.Comment
		result := contentManager.db.Find(&comment, commentIDNum)
		if result.Error != nil {
			renderPreview(c, contentManager.db, http.StatusInternalServerError, caffID)
			return
		}

		contentManager.db.Delete(comment)

		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/content/preview/%d", comment.CaffContentID))
	})
}
