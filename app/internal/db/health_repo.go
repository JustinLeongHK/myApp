package db

import (
	"context"
	"fmt"
)

func (r *PostgresRepo) GetTime(ctx context.Context) (string, error) {
	var now string
	err := r.db.QueryRowContext(ctx, `SELECT NOW()`).Scan(&now)
	if err != nil {
		return "", fmt.Errorf("get time: %w", err)
	}
	return now, nil
}
