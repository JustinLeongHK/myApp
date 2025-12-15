package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

func Home(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	fmt.Fprintf(w, "Welcome! DB is ready to use.\n")
}

func Hello(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	// Example query
	row := db.QueryRow("SELECT NOW()")
	var now string
	row.Scan(&now)

	fmt.Fprintf(w, "Hello! Database time is: %s\n", now)
}
