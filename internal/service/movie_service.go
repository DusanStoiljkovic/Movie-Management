package service

import (
	"context"
	"errors"
	dto "movie-management/internal/dto/movie"
	"movie-management/internal/models"
	"movie-management/internal/repository"
)

type MovieService struct {
	repo *repository.MovieRepository
}

func NewMovieService(repo *repository.MovieRepository) *MovieService {
	return &MovieService{repo: repo}
}

func (s *MovieService) CreateMovie(ctx context.Context, req *dto.RequestMovie) error {
	if req.Title == "" {
		return errors.New("title is required")
	}

	if req.Year < 1888 {
		return errors.New("invalid year")
	}

	movie := models.Movie{
		Title:  req.Title,
		Year:   req.Year,
		Rating: req.Rating,
	}

	if len(req.GenreIDs) > 0 {
		genres, err := s.repo.GetGenresByIDs(ctx, req.GenreIDs)
		if err != nil {
			return err
		}
		movie.Genres = genres
	}

	return s.repo.CreateMovie(ctx, &movie)
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
