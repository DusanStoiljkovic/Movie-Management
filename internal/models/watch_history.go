package models

import "time"

type WatchHistory struct {
	ID        int `gorm:"primaryKey"`
	UserID    int `gorm:"uniqueIndex:idx_user_movie"`
	MovieID   int `gorm:"uniqueIndex:idx_user_movie"`
	WatchedAt time.Time
}
