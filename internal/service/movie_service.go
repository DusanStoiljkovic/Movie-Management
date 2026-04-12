package service

import (
	"context"
	"errors"
	"movie-management/internal/models"
	"movie-management/internal/repository"
)

type MovieService struct {
	repo *repository.MovieRepository
}

func NewMovieService(repo *repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) CreateMovie(ctx context.Context, movie *models.Movie) error {
	if movie.Title == "" {
		return errors.New("title is required")
	}

	if movie.Year < 1888 {
		return errors.New("invalid year")
	}

	return s.repo.CreateMovie(ctx, movie)
}

func (s *MovieService) GetMovieByID(ctx context.Context, id uint) (*models.Movie, error) {
	return s.repo.GetMovieByID(ctx, id)
}

func (s *MovieService) GetMovies(
	ctx context.Context,
	limit, offset int,
	sort string,
	genre string,
	minYear, maxYear int,
	minRating float64,
) ([]models.Movie, error) {

	// default pagination
	if limit == 0 {
		limit = 10
	}
	return s.repo.GetMovies(ctx, limit, offset, sort, genre, minYear, maxYear, minRating)
}

func (s *MovieService) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	if movie.ID == 0 {
		return errors.New("movie ID is required")
	}

	return s.repo.UpdateMovie(ctx, movie)
}

func (s *MovieService) DeleteMovie(ctx context.Context, id uint) error {
	return s.repo.DeleteMovie(ctx, id)
}

func (s *MovieService) AddGenres(ctx context.Context, movieID uint, genres []models.Genre) error {
	return s.repo.AddGenresToMovie(ctx, movieID, genres)
}
