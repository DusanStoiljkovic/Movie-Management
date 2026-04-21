package service

import (
	"context"
	"errors"
	"movie-management/internal/models"
	"movie-management/internal/repository"
)

type GenreService struct {
	repo *repository.GenreRepository
}

func NewGenreService(repo *repository.GenreRepository) *GenreService {
	return &GenreService{repo: repo}
}

func (s *GenreService) GetAll(ctx context.Context) ([]models.Genre, error) {
	genres, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, errors.New("Problem with fetching genres from db")
	}

	return genres, nil
}

func (s *GenreService) AddGenre(ctx context.Context, genre *models.Genre) error {
	return s.repo.AddGenre(ctx, genre)
}

func (s *GenreService) DeleteGenreByID(ctx context.Context, genreIDs []int) error {
	return s.repo.DeleteGenreByID(genreIDs)
}
