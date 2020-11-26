package model

import (
	"database/sql/driver"
	"encoding/hex"
	"gorm.io/gorm"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model

	Username     string       `gorm:"unique;not null"`
	PasswordHash PasswordHash `gorm:"not null"`
}

type PasswordHash struct {
	hashBytes []byte
}

func (p PasswordHash) Value() (driver.Value, error) {
	return hex.EncodeToString(p.hashBytes), nil
}

func (p *PasswordHash) Scan(value interface{}) error {
	asStr, err := driver.String.ConvertValue(value)
	if err != nil {
		return nil
	}

	p.hashBytes, err = hex.DecodeString(asStr.(string))
	return err
}

func HashPassword(pw string) PasswordHash {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	if err != nil {
		panic(err) //TODO
	}
	return PasswordHash{hash}
}

func CheckPasswordsMatch(provided string, pw PasswordHash) error {
	return bcrypt.CompareHashAndPassword(pw.hashBytes, []byte(provided))
}
