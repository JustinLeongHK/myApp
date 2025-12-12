package main

import (
	"log"

	"net/http"

	"github.com/JustinLeongHK/myApp/routes"
)

func main() {
	mux := routes.NewRouter()

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
