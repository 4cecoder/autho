package routing

import (
	"auth/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Profile endpoint
	r.HandleFunc("/profile/{id}", handlers.Profile).Methods("GET")

	// Register endpoint
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	// Login endpoint
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Start the server
	http.ListenAndServe(":8080", r)
}
