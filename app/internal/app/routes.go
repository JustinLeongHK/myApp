package app

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/JustinLeongHK/myApp/internal/config"
	"github.com/JustinLeongHK/myApp/internal/user"
	_ "github.com/lib/pq"
)

func NewRouter(cfg *config.Config) http.Handler {
	sqlDB := initDB(cfg)
	repo := user.NewPostgresRepository(sqlDB)
	service := user.NewService(repo)
	handler := user.NewHandler(service)

	mux := http.NewServeMux()
	registerRoutes(mux, handler)
	return mux
}

func initDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("failed to connect to DB: %v", err))
	}
	return db
}

func registerRoutes(mux *http.ServeMux, h *user.Handler) {
	mux.HandleFunc("/users", h.CreateUser)
	mux.HandleFunc("/hello", h.GetDBTimeHandler)
}
