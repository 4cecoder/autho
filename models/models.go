package models

// models/models.go
import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// User represents a user account
type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name" validate:"required"`
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"-"`
}

type Credentials struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Password     string `json:"password" validate:"required"`
}

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// Validate checks if a user struct is valid
func (u *User) Validate() error {
	// Check if the name is empty or contains invalid characters
	if u.Name == "" || strings.ContainsAny(u.Name, "!@#$%^&*()_+-=") {
		return errors.New("invalid name")
	}

	// Check if the email address is empty or invalid
	if u.EmailAddress == "" || !strings.Contains(u.EmailAddress, "@") {
		return errors.New("invalid email address")
	}

	return nil
}
