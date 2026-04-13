package handler

import (
	"encoding/json"
	"movie-management/internal/models"
	"movie-management/internal/service"
	"net/http"
)

type GenreHandler struct {
	service *service.GenreService
}

func NewGenreHandler(service *service.GenreService) *GenreHandler {
	return &GenreHandler{service: service}
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
