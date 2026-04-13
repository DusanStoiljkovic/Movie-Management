package service

import (
	"context"
	"errors"
	"movie-management/internal/models"
	"movie-management/internal/repository"

	"golang.org/x/crypto/bcrypt"
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

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Problem with hashing")
	}

	user.Password = string(hashPassword)

	return s.repo.CreateUser(ctx, user)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.GetUserByEmail(ctx, email)
}

func (s *UserService) Login(ctx context.Context, email, password string) (*models.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, role string) ([]models.User, error) {
	if role != "admin" {
		return nil, errors.New("forbidden")
	}

	return s.repo.GetAll(ctx)
}
