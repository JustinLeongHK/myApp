package user

import (
	"context"

	"github.com/JustinLeongHK/myApp/internal/model"
)

// A repository interface defines what operations you expect from your data storage, without saying how they are implemented.

type Repository interface {
	SaveUser(ctx context.Context, user model.User) error
	GetTime(ctx context.Context) (string, error)
}
