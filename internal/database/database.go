package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"movie-management/internal/models"
)

func ConnectDB() *gorm.DB {
	dsn := "host=db user=postgres password=postgres dbname=movie_management port=5432 ssl=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB error: ", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Movie{},
		&models.Genre{},
		&models.WatchHistory{},
	)

	if err != nil {
		log.Fatal("Greska pri migraciji: ", err)
	}

	return db
}
