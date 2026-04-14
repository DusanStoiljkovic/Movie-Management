package dto

type ResponseMovie struct {
	ID     int      `json:"id"`
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Rating float64  `json:"rating"`
	Genres []string `json:"genres"`
}
