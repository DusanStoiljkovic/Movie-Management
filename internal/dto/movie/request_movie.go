package dto

type RequestMovie struct {
	Title    string  `json:"title"`
	Year     int     `json:"year"`
	Rating   float64 `json:"rating"`
	GenreIDs []int   `json:"genre_ids"`
}
