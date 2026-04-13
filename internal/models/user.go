package models

import "time"

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Role      string
	CreatedAt time.Time

	FavouriteGenres []Genre `gorm:"many2many:user_favourite_genres;"`
}
