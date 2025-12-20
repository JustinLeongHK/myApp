package app

import (
	"log"
	"net/http"

	"github.com/JustinLeongHK/myApp/internal/config"
)

func Run() {
	cfg := config.Load()
	router := NewRouter(cfg)

	log.Printf("Starting server  on :%s", "8080")
	log.Fatal(http.ListenAndServe(":"+"8080", router))
}
