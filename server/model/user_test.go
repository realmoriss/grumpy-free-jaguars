package model

import (
	"testing"
)

func TestHashing() {
	pw := HashPassword("password")

	t.Run("TestHashing", func(t *testing.T) {
		if !(CheckPasswordsMatch("password", pw)) {
			t.Errorf("Password matching doesn't work for 'password' and %v", pw)
		}
		if CheckPasswordsMatch("notpassword", pw) {
			t.Errorf("False correct password 'notpassword' for %v", pw)
		}
	})
}

func TestValueScan() {
	pw := HashPassword("password")

	t.Run("TestValueScan", func(t *testing.T) {
		value, err := pw.Value()
		if err != nil {
			t.Error(nil)
		}
		var pw2 PasswordHash
		pw2.Scan(value)
		if !(CheckPasswordsMatch("password", pw2)) {
			t.Errorf("Password serialization broken for 'password' and %v", pw)
		}
	})
}

func TestUserAdmin() {
	user := User{
		Username:     "name",
		PasswordHash: HashPassword("xxx"),
		IsAdmin:      false,
	}

	t.Run("TestUserAdmin", func(t *testing.T) {
		if IsAdministrator(user) {
			t.Errorf("A non-admin user evaluated as admin")
		}
		user.IsAdmin = true
		if !IsAdministrator(user) {
			t.Errorf("An admin user evaluated as non-admin")
		}
	})
}
