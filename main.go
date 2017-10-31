package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// our main function
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	router := mux.NewRouter()
	log.Printf("Starting server at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
