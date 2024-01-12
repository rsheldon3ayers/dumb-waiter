package db

import (
	"log"

	"github.com/rsheldon3ayers/dumb-module/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	DB_URL := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"

	// Connect to database
	db, err := gorm.Open(postgres.Open(DB_URL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	// Migrate the schema
	db.AutoMigrate(&models.Developer{})

	return db
}
