package main

import (
	"log"
	"movie-management/internal/database"
	"movie-management/internal/handler"
	"movie-management/internal/repository"
	"movie-management/internal/service"
	"net/http"
)

func main() {
	//DB
	db := database.ConnectDB()

	// Repo
	movieRepo := repository.NewMovieRepository(db)
	genreRepo := repository.NewGenreRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Service
	movieService := service.NewMovieService(movieRepo, genreRepo)
	genreService := service.NewGenreService(genreRepo)
	userService := service.NewUserService(userRepo)

	// Handler
	movieHandler := handler.NewMovieHandler(movieService)
	genreHandler := handler.NewGenreHandler(genreService)
	userHandler := handler.NewUserHandler(userService)

	// Router
	r := handler.RegisterRoutes(movieHandler, genreHandler, userHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)

}
