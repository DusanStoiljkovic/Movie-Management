package models

import "time"

type WatchHistory struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	MovieID   int
	WatchedAt time.Time
}
