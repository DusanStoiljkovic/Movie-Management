package service

import (
	"context"
	"errors"
	dto "movie-management/internal/dto/user"
	"movie-management/internal/mapper"
	"movie-management/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo      *repository.UserRepository
	genreRepo *repository.GenreRepository
}

func NewUserService(repo *repository.UserRepository, genreRepo *repository.GenreRepository) *UserService {
	return &UserService{
		repo:      repo,
		genreRepo: genreRepo,
	}
}

func (s *UserService) Register(ctx context.Context, req *dto.RequestRegisterUser) (*dto.ResponseRegisterUser, error) {

	if req.Email == "" || req.Password == "" {
		return nil, errors.New("email and password required")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("Problem with hashing")
	}

	userModel := mapper.ToUserModelFromRegister(req, string(hashPassword))

	if err := s.repo.CreateUser(ctx, userModel); err != nil {
		return nil, err
	}

	res := mapper.ToRegisterResponse(userModel)

	return res, nil
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*dto.ResponseLoginUser, error) {
	response, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return mapper.ToLoginResponse(response), nil

}

func (s *UserService) Login(ctx context.Context, email, password string) (*dto.ResponseLoginUser, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return mapper.ToLoginResponse(user), nil
}

func (s *UserService) GetAllUsers(ctx context.Context, role string) ([]dto.ResponseRegisterUser, error) {
	if role != "admin" {
		return nil, errors.New("forbidden")
	}

	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("Users not found")
	}

	response := []dto.ResponseRegisterUser{}

	for _, u := range users {
		response = append(response, *mapper.ToRegisterResponse(&u))
	}

	return response, nil
}

func (s *UserService) AssignFavGenres(ctx context.Context, userID int, genreIDs []int) (*dto.UserResponse, error) {
	genres, err := s.genreRepo.GetGenresByIDs(ctx, genreIDs)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	updatedUser, err := s.repo.AssignGenresToUser(ctx, user, genres)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return mapper.ToUserResponse(updatedUser, genres), nil

}
