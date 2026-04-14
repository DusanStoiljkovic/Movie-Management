package repository

import (
	"context"
	"movie-management/internal/models"

	"gorm.io/gorm"
)

type GenreRepository struct {
	db *gorm.DB
}

func NewGenreRepository(db *gorm.DB) *GenreRepository {
	return &GenreRepository{db: db}
}

func (repo *GenreRepository) AddGenre(ctx context.Context, genre *models.Genre) error {
	return repo.db.WithContext(ctx).Create(genre).Error
}

func (repo *GenreRepository) GetGenresByIDs(ctx context.Context, genreIDs []int) ([]models.Genre, error) {
	var genres []models.Genre

	err := repo.db.WithContext(ctx).Where("id IN ?", genreIDs).Find(&genres).Error
	if err != nil {
		return nil, err
	}

	return genres, nil
}
