package models

import "time"

type Movie struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Year      int
	Rating    float64
	CreatedAt time.Time

	Genres []Genre `gorm:"many2many:movie_genres;constraint:OnDelete:CASCADE;"`
}
