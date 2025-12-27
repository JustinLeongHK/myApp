package user

import (
	"context"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(ctx context.Context, user User) (*User, error) {
	return s.repo.CreateUser(ctx, user.Email)
}

func (s *Service) GetDBTime(ctx context.Context) (string, error) {
	return s.repo.GetTime(ctx)
}
