package service

import (
	"context"
	"movie-management/internal/models"
	"movie-management/internal/repository"
	"time"
)

type WatchHistoryService struct {
	repo *repository.WatchHistoryRepository
}

func NewWatchHistoryService(repo *repository.WatchHistoryRepository) *WatchHistoryService {
	return &WatchHistoryService{repo: repo}
}

func (s *WatchHistoryService) Add(ctx context.Context, userID, movieID int) error {
	wh := &models.WatchHistory{
		UserID:    userID,
		MovieID:   movieID,
		WatchedAt: time.Now(),
	}

	return s.repo.Add(ctx, wh)
}
