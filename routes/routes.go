package routes

import (
	"net/http"

	"github.com/JustinLeongHK/myApp/handlers"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/hello", handlers.Hello)

	return mux
}
