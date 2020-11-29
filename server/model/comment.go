package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CaffContent   *CaffContent
	CaffContentID uint `gorm:"not null"`
	User          *User
	UserID        uint   `gorm:"not null"`
	Text          string `gorm:"not null"`
}
