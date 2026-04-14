package mapper

import (
	dto "movie-management/internal/dto/movie"
	"movie-management/internal/models"
)

func MapToMovieResponse(movie *models.Movie) *dto.ResponseMovie {
	var genres []string

	for _, g := range movie.Genres {
		genres = append(genres, g.Name)
	}

	return &dto.ResponseMovie{
		ID:     int(movie.ID),
		Title:  movie.Title,
		Year:   int(movie.Year),
		Rating: float64(movie.Rating),
		Genres: genres,
	}
}
