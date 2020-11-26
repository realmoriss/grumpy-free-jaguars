package model

import (
	"time"

	"gorm.io/gorm"
)

type CaffContent struct {
	gorm.Model

	UserID    uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Personally, I think this syntax for "column permission" is a bit retarted.
	RawFile     []byte `gorm:"->;<-:create;not null"`
	PreviewWebp []byte `gorm:"->;<-:create;not null"`
}
