package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(MovieHandler *MovieHandler, GenreHandler *GenreHandler, UserHandler *UserHandler) http.Handler {
	r := chi.NewRouter()

	r.Route("/movies", func(r chi.Router) {
		r.Post("/", MovieHandler.CreateMovie)
		r.Get("/", MovieHandler.GetMovies)
		r.Get("/{id}", MovieHandler.GetMovieByID)
		r.Put("/{id}", MovieHandler.UpdateMovie)
		r.Delete("/{id}", MovieHandler.DeleteMovie)
	})

	r.Route("/genres", func(r chi.Router) {
		r.Post("/", GenreHandler.AddGenre)
	})

	r.Route("/", func(r chi.Router) {
		r.Post("/register", UserHandler.Register)
		r.Post("/login", UserHandler.Login)
		r.Get("/users", UserHandler.GetAllUsers)
	})

	return r
}
