package content

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"

	"server/model"
)

var (
	logger *log.Logger
)

func init() {
	logger = log.New(os.Stdout, "content: ", log.LstdFlags|log.LUTC|log.Lmsgprefix)
}

type ContentEndpoint struct {
	db *gorm.DB
}

func NewEndpoint(router gin.IRouter, db *gorm.DB) *ContentEndpoint {
	db.AutoMigrate(&model.CaffContent{})
	db.AutoMigrate(&model.Comment{})

	self := &ContentEndpoint{db: db}

	self.addBrowseEndpoints(router)
	self.addUploadEndpoints(router)
	self.addPreviewEndpoints(router)

	return self
}
