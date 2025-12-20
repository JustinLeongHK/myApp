package user

import (
	"context"

	"github.com/JustinLeongHK/myApp/internal/model"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, user model.User) error {
	return s.repo.SaveUser(ctx, user)
}

func (s *Service) GetDBTime(ctx context.Context) (string, error) {
	return s.repo.GetTime(ctx)
}
