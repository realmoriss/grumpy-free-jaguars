package model

import (
	"bytes"
	"context"
	"errors"
	"image/png"
	"io"
	"io/ioutil"
	"server/caff"

	"gorm.io/gorm"
)

var (
	parser *caff.CaffParser
)

func init() {
	newParser, err := caff.NewParser("/mnt/WinStorage/Work/grumpy-free-jaguars/libcaff/build/caff")
	if err != nil {
		panic("unable to open caff parser!")
	}
	parser = newParser
}

type CaffContent struct {
	gorm.Model
	User       *User
	UserID     uint
	Title      string `gorm:"not null"`
	RawFile    []byte `gorm:"->;<-:create;not null"`
	PreviewPng []byte `gorm:"->;<-:create;not null"`
}

type CaffPreview struct {
	gorm.Model
	User       *User
	UserID     uint
	Title      string
	PreviewPng []byte
}

var (
	ErrParserUnavailable = errors.New("parser is unavailable")
)

func ParseCaff(file io.Reader) (*CaffContent, error) {
	if parser == nil {
		return nil, ErrParserUnavailable
	}

	originFile, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, ErrParserUnavailable
	}

	rawImg, err := parser.ParseCAFF(context.Background(), bytes.NewReader(originFile))
	if err != nil {
		return nil, err
	}

	var imgBuf bytes.Buffer
	err = png.Encode(&imgBuf, rawImg)
	if err != nil {
		return nil, err
	}

	return &CaffContent{
		RawFile:    originFile,
		PreviewPng: imgBuf.Bytes(),
	}, err
}
