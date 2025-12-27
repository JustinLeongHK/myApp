package user

import (
	"context"
)

// A repository interface defines what operations you expect from your data storage, without saying how they are implemented.

type Repository interface {
	CreateUser(ctx context.Context, email string) (*User, error)
	GetTime(ctx context.Context) (string, error)
}
