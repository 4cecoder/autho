package db

// db/insert.go

import (
	"github.com/byte-cats/autho/models"
)

// InsertUser inserts a new user into the database
func InsertUser(user models.User) error {
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
