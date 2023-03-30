package db

// db/get_user.go
import (
	"database/sql"
	"fmt"
	"github.com/bytecats/auth/models"
)

// GetUserByID retrieves a user with the given ID from the database
func GetUserByID(id uint) (models.User, error) {
	// Prepare the SELECT statement
	stmt, err := DB.Prepare("SELECT id, name, email_address, password FROM users WHERE id = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	// Execute the SELECT statement
	row := stmt.QueryRow(id)

	// Scan the result into a User struct
	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.EmailAddress, &user.Password)
	if err == sql.ErrNoRows {
		return models.User{}, fmt.Errorf("user not found")
	} else if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// GetUserByEmail retrieves a user with the given email from the database
func GetUserByEmail(email string) (models.User, error) {
	// Prepare the SELECT statement
	stmt, err := DB.Prepare("SELECT id, name, email_address, password FROM users WHERE email_address = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	// Execute the SELECT statement
	row := stmt.QueryRow(email)

	// Scan the result into a User struct
	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.EmailAddress, &user.Password)
	if err == sql.ErrNoRows {
		return models.User{}, fmt.Errorf("user not found")
	} else if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func GetUserByEmailAndPassword(email, password string) (models.User, error) {
	// Prepare the SELECT statement
	stmt, err := DB.Prepare("SELECT id, name, email_address, password FROM users WHERE email_address = ? AND password = ?")
	if err != nil {
		return models.User{}, err
	}
	defer stmt.Close()

	// Execute the SELECT statement
	row := stmt.QueryRow(email, password)

	// Scan the result into a User struct
	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.EmailAddress, &user.Password)
	if err == sql.ErrNoRows {
		return models.User{}, fmt.Errorf("user not found")
	} else if err != nil {
		return models.User{}, err
	}

	return user, nil

}
