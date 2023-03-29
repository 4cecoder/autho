package main

import (
	"log"
	"net/http"

	"auth/db"
)

func main() {
	defer db.DB.Close()

	log.Fatal(http.ListenAndServe(":8081", nil))
}
