package handler

import (
	"encoding/json"
	dto "movie-management/internal/dto/movie"
	"movie-management/internal/mapper"
	"movie-management/internal/models"
	"movie-management/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type MovieHandler struct {
	service *service.MovieService
}

func NewMovieHandler(service *service.MovieService) *MovieHandler {
	return &MovieHandler{service: service}
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var req dto.RequestMovie

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateMovie(r.Context(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(req)
}

func (h *MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	movie, err := h.service.GetMovieByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(&movie)
}

func (h *MovieHandler) GetMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	sort := query.Get("sort")
	genre := query.Get("genre")

	minYear, _ := strconv.Atoi(query.Get("minYear"))
	maxYear, _ := strconv.Atoi(query.Get("maxYear"))

	minRating, _ := strconv.ParseFloat(query.Get("minRating"), 64)

	movies, err := h.service.GetMovies(r.Context(), limit, offset, sort, genre, minYear, maxYear, minRating)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	response := []dto.ResponseMovie{}

	for _, m := range movies {
		response = append(response, *mapper.MapToMovieResponse(&m))
	}
	json.NewEncoder(w).Encode(response)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)

	var movie models.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	movie.ID = id

	err := h.service.UpdateMovie(r.Context(), &movie)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, _ := strconv.Atoi(idParam)

	err := h.service.DeleteMovie(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MovieHandler) AddGenresToMovie(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	movieID, _ := strconv.Atoi(idParam)

	var genreIDs []int

	err := json.NewDecoder(r.Body).Decode(&genreIDs)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	movie, err := h.service.AssignGenresToMovie(r.Context(), movieID, genreIDs)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(movie)
}

func (h *MovieHandler) DeleteSpecificMoviesGenre(w http.ResponseWriter, r *http.Request) {
	idParam1 := chi.URLParam(r, "movieID")
	idParam2 := chi.URLParam(r, "genreID")
	movieID, _ := strconv.Atoi(idParam1)
	genreID, _ := strconv.Atoi(idParam2)

	err := h.service.DeleteSpecificMoviesGenre(r.Context(), movieID, genreID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	movie, err := h.service.GetMovieByID(r.Context(), movieID)

	json.NewEncoder(w).Encode(movie)

}
