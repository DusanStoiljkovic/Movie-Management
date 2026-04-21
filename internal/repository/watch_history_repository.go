package repository

import (
	"context"
	"movie-management/internal/models"

	"gorm.io/gorm"
)

type WatchHistoryRepository struct {
	db *gorm.DB
}

func NewWatchHistoryRepository(db *gorm.DB) *WatchHistoryRepository {
	return &WatchHistoryRepository{db: db}
}

func (r *WatchHistoryRepository) Add(ctx context.Context, wh *models.WatchHistory) (*models.WatchHistory, error) {
	err := r.db.WithContext(ctx).Create(wh).Error
	if err != nil {
		return nil, err
	}

	return wh, nil
}
