package handlers

import (
	"encoding/json"
	"github.com/byte-cats/autho/db"
	"github.com/byte-cats/autho/models"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	// You should replace this with the actual user ID or email, possibly from a session or JWT
	userID := uint(1)

	user, err := db.GetUserByID(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error retrieving user profile"))
		return
	}

	userProfile := models.UserProfile{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.EmailAddress,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userProfile)
}
