package routes

import (
	"database/sql"
	"net/http"

	"github.com/JustinLeongHK/myApp/handlers"
	_ "github.com/lib/pq"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.Home(w, r, db)
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		handlers.Hello(w, r, db)
	})

	return mux
}
