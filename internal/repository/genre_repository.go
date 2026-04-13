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
