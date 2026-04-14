package repository

import (
	"context"
	"movie-management/internal/models"

	"gorm.io/gorm"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{db: db}
}

func (r *MovieRepository) GetGenresByIDs(ctx context.Context, ids []int) ([]models.Genre, error) {
	var genres []models.Genre

	err := r.db.WithContext(ctx).Where("id IN ?", ids).Find(&genres).Error

	return genres, err
}

func (r *MovieRepository) CreateMovie(ctx context.Context, movie *models.Movie) error {
	return r.db.WithContext(ctx).Create(movie).Error
}

func (r *MovieRepository) GetMovieByID(ctx context.Context, id int) (*models.Movie, error) {
	var movie models.Movie

	err := r.db.WithContext(ctx).
		Preload("Genres").
		First(&movie, id).Error

	if err != nil {
		return nil, err
	}

	return &movie, nil

}

func (r *MovieRepository) GetMovies(
	ctx context.Context,
	limit, offset int,
	sort string,
	genre string,
	minYear, maxYear int,
	minRating float64,
) ([]models.Movie, error) {
	var movies []models.Movie

	query := r.db.WithContext(ctx).Model(&models.Movie{}).
		Preload("Genres")

	if genre != "" {
		query = query.Joins("JOIN movie_genres mg On mg.movie_id = movies.id").
			Joins("JOIN genres g ON g.id = mg.genre_id").
			Where("g.name = ?", genre)
	}

	if minYear != 0 && maxYear != 0 {
		query = query.Where("year BETWEEN ? AND ?", minYear, maxYear)
	}

	if minRating != 0 {
		query = query.Where("raing >= ?", minRating)
	}

	if sort != "" {
		query = query.Order(sort)
	}

	if offset >= 0 {
		query = query.Offset(offset)
	}

	err := query.Find(&movies).Error
	return movies, err
}

func (r *MovieRepository) UpdateMovie(ctx context.Context, movie *models.Movie) error {
	return r.db.WithContext(ctx).Save(movie).Error
}

func (r *MovieRepository) DeleteMovie(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Movie{}, id).Error
}

func (r *MovieRepository) AddGenresToMovie(ctx context.Context, movieID int, genres []models.Genre) (*models.Movie, error) {
	var movie models.Movie

	if err := r.db.First(&movie, movieID).Error; err != nil {
		return nil, err
	}

	err := r.db.Model(&movie).Association("Genres").Replace(genres)
	if err != nil {
		return nil, err
	}

	return &movie, nil
}
