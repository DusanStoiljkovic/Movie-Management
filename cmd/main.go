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

	// Service
	movieService := service.NewMovieService(movieRepo)

	// Handler
	movieHandler := handler.NewMovieHandler(movieService)

	// Router
	r := handler.RegisterRoutes(movieHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", r)

}
