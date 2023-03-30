package serve

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func Serve(r *mux.Router) {
	// get godot env variable api port and set it to port variable
	godotenv.Load(".env")
	port := os.Getenv("API_PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
