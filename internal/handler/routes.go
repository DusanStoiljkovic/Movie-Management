package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(MovieHandler *MovieHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/movies", func(r chi.Router) {
		r.Post("/", MovieHandler.CreateMovie)
		r.Get("/", MovieHandler.GetMovies)
		r.Get("/{id}", MovieHandler.GetMovieByID)
		r.Put("/{id}", MovieHandler.UpdateMovie)
		r.Delete("/{id}", MovieHandler.DeleteMovie)
	})

	return r
}
