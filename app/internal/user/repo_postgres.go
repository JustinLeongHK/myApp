package user

import (
	"context"
	"database/sql"
)

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) CreateUser(ctx context.Context, email string) (*User, error) {

	u := &User{
		Email: email,
	}

	err := r.db.QueryRowContext(ctx, `INSERT INTO users (email) VALUES ($1) RETURNING id`, email).Scan(&u.ID)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *postgresRepository) GetTime(ctx context.Context) (string, error) {

	var now string

	err := r.db.QueryRowContext(ctx, `SELECT NOW()`).Scan(&now)

	if err != nil {
		return "", err
	}

	return now, nil
}
