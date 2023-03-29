package handlers

import (
	"auth/db"
	"auth/models"
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

// Login handles POST requests to the /login endpoint
func Login(w http.ResponseWriter, r *http.Request) {
	var creds models.Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the user with the given email and password exists in the database
	user, err := db.GetUserByEmailAndPassword(creds.Email, creds.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(Response{Message: "Invalid credentials"})
		return
	}

	// Generate a JWT token for the user
	token, err := GenerateToken(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write the response
	json.NewEncoder(w).Encode(Response{Message: token})
}
