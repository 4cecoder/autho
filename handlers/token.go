package handlers

import (
	"auth/models"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Load the .env file on package initialization
func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

// Claims represents the JWT claims
type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

// GenerateToken generates a JWT token for the given user
func GenerateToken(user models.User) (string, error) {
	// Set the expiration time of the token to 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims with the user ID and expiration time
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Read the secret key from the environment variable
	jwtKey := []byte(os.Getenv("COMPANY_SUPER_SECRET"))

	// Create the JWT token with the claims and secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	// Return the JWT token
	return tokenString, nil
}
