package models

import "time"

type WatchHistory struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint
	MovieID   uint
	WatchedAt time.Time
}
