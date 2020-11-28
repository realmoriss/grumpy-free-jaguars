package content

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"server/middleware"
	"server/model"
	"server/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func renderUpload(c *gin.Context, status int, err error) {
	util.HtmlWithContext(c, status, "upload", gin.H{
		"title": "Upload",
		"error": err,
	})
}

var (
	ErrNotLoggedIn   = errors.New("User not logged in")
	ErrUnableToParse = errors.New("Unable to parse image")
)

func (contentManager ContentEndpoint) addUploadEndpoints(router gin.IRouter) {
	router.GET("/upload", func(c *gin.Context) {
		renderUpload(c, http.StatusOK, nil)
	})

	router.POST("/upload", func(c *gin.Context) {
		var provided struct {
			Title string                `form:"title"`
			File  *multipart.FileHeader `form:"file"`
		}

		if err := c.ShouldBind(&provided); err != nil {
			renderUpload(c, http.StatusBadRequest, err)
			return
		}

		file, err := provided.File.Open()
		if err != nil {
			renderUpload(c, http.StatusBadRequest, err)
			return
		}

		currentUser := middleware.CurrentUser(c)
		if currentUser == nil {
			renderUpload(c, http.StatusUnauthorized, ErrNotLoggedIn)
			return
		}

		caff, err := model.ParseCaff(file)
		if err != nil {
			renderUpload(c, http.StatusBadRequest, ErrUnableToParse)
			return
		}

		caff.User = currentUser
		caff.Title = provided.Title

		result := contentManager.db.Clauses(clause.OnConflict{
			UpdateAll: true,
		}).Create(&caff)
		if result.Error != nil {
			renderUpload(c, http.StatusBadRequest, result.Error)
			return
		}

		c.Redirect(http.StatusSeeOther, fmt.Sprintf("/content/preview/%d", caff.ID))
	})
}
