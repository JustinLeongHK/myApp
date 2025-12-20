package db

import (
	"context"

	"github.com/JustinLeongHK/myApp/internal/model"
)

func (r *PostgresRepo) SaveUser(ctx context.Context, u model.User) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO users (id, name) VALUES ($1, $2)`,
		u.ID,
		u.Name,
	)
	return err
}
