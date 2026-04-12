package service

import (
	"context"
	"errors"
	"movie-management/internal/models"
	"movie-management/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, user *models.User) error {
	if user.Email == "" || user.Password == "" {
		return errors.New("email and password required")
	}

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}
