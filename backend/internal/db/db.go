package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/rsheldon3ayers/dumb-module/backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("DATABASE_URL")

	// Connect to database
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	// Migrate the schema
	db.AutoMigrate(&models.Developer{})

	return db
}
