package dto

import "time"

type UserResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	Role            string    `json:"role"`
	FavouriteGenres []string  `json:"fav_genres"`
	CreatedAt       time.Time `json:"created_at"`
}
