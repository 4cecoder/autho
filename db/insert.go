package db

// db/insert.go

import (
	"github.com/bytecats/auth/models"
)

// InsertUser inserts a new user into the database
func InsertUser(user models.User) error {
	// Prepare the INSERT statement
	stmt, err := DB.Prepare("INSERT INTO users (name, email_address, password) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the INSERT statement
	_, err = stmt.Exec(user.Name, user.EmailAddress, user.Password)
	if err != nil {
		return err
	}

	return nil
}
