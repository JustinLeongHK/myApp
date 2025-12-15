package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"net/http"

	"github.com/JustinLeongHK/myApp/routes"
	_ "github.com/lib/pq"
)

func main() {
	// Build DSN from environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// Connect to Postgres
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Ping to make sure the database is reachable
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	// Pass the db to your router (so handlers can use it)
	mux := routes.NewRouter(db)

	log.Println("Server starting on :8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
