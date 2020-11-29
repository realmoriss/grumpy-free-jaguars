package user

import (
	"errors"
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
	logger = log.New(os.Stdout, "user: ", log.LstdFlags|log.LUTC|log.Lmsgprefix)
}

type UserEndpoint struct {
	db *gorm.DB
}

var (
	ErrPasswordDoNotMatch       = errors.New("Passwords did not match")
	ErrPasswordInsecure         = errors.New("Password must be at least 8 characters long") // TODO: at least not empty, perhaps?
	ErrRegistrationUnsuccessful = errors.New("Registration failed with the specified username")
)

func NewEndpoint(router gin.IRouter, db *gorm.DB) *UserEndpoint {
	db.AutoMigrate(&model.User{})

	self := &UserEndpoint{db: db}

	self.addLoginEndpoints(router)
	self.addRegisterEndpoints(router)

	return self
}
