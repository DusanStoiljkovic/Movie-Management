package service

import (
	"context"
	"movie-management/internal/models"
	"movie-management/internal/repository"
)

type GenreService struct {
	repo *repository.GenreRepository
}

func NewGenreService(repo *repository.GenreRepository) *GenreService {
	return &GenreService{repo: repo}
}

func (s *GenreService) AddGenre(ctx context.Context, genre *models.Genre) error {
	return s.repo.AddGenre(ctx, genre)
}
