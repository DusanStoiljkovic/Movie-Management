package handler

import (
	"encoding/json"
	"movie-management/internal/models"
	"movie-management/internal/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
}

func (handler *GenreHandler) GetAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := handler.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(&genres)
}

func (handler *GenreHandler) AddGenre(w http.ResponseWriter, r *http.Request) {
	var genre models.Genre

	if err := json.NewDecoder(r.Body).Decode(&genre); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := handler.service.AddGenre(r.Context(), &genre)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(genre)
}

func (handler *GenreHandler) DeleteGenre(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	genreID, _ := strconv.Atoi(idParam)

	err := handler.service.DeleteGenreByID(r.Context(), []int{genreID})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	genres, err := handler.service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(&genres)
}
