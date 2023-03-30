package routing

import (
	"github.com/byte-cats/autho/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	r := mux.NewRouter()

	// Profile endpoint
	r.HandleFunc("/profile/{id}", handlers.Profile).Methods("GET")

	// Register endpoint
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	// Login endpoint
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	return r
}
