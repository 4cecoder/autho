package handlers

// handlers/register.go

import (
	"encoding/json"
	"github.com/byte-cats/autho/db"
	"github.com/byte-cats/autho/models"
	"net/http"
)

// RegisterRequest represents the request body for the Register handler
type RegisterRequest struct {
	Name         string `json:"name"`
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

// RegisterResponse represents the response body for the Register handler
type RegisterResponse struct {
	Token string `json:"token"`
}

// Register registers a new user
func Register(w http.ResponseWriter, r *http.Request) {
	// Decode the request body into a RegisterRequest struct
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate the request body
	if request.Name == "" || request.EmailAddress == "" || request.Password == "" {
		http.Error(w, "name, email_address, and password are required", http.StatusBadRequest)
		return
	}

	// Check if the email is already registered
	_, err = db.GetUserByEmail(request.EmailAddress)
	if err == nil {
		http.Error(w, "email_address is already registered", http.StatusConflict)
		return
	}

	// Hash the password
	hashedPassword, err := models.HashPassword(request.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new User object
	user := models.User{
		Name:         request.Name,
		EmailAddress: request.EmailAddress,
		Password:     hashedPassword,
	}

	// Insert the user into the database
	err = db.InsertUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Generate a JWT token for the user
	token, err := GenerateToken(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a RegisterResponse object and encode it as JSON
	response := RegisterResponse{
		Token: token,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
