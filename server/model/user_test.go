package model

import (
	"testing"
)

func TestHashing(t *testing.T) {
	pw := HashPassword("password")

	t.Run("TestHashing", func(t *testing.T) {
		if CheckPasswordsMatch("password", pw) != nil {
			t.Errorf("Password matching doesn't work for 'password' and %v", pw)
		}
		if CheckPasswordsMatch("notpassword", pw) == nil {
			t.Errorf("False correct password 'notpassword' for %v", pw)
		}
	})
}

func TestValueScan(t *testing.T) {
	pw := HashPassword("password")

	t.Run("TestValueScan", func(t *testing.T) {
		value, err := pw.Value()
		if err != nil {
			t.Error(nil)
		}
		var pw2 PasswordHash
		pw2.Scan(value)
		if CheckPasswordsMatch("password", pw2) != nil {
			t.Errorf("Password serialization broken for 'password' and %v", pw)
		}
	})
}

func TestUserAdmin(t *testing.T) {
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
