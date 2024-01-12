package seed

import (
	"log"

	"github.com/rsheldon3ayers/dumb-module/backend/internal/models"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	// Sample data
	developers := []models.Developer{
		{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"},
		{FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com"},
		{FirstName: "Alice", LastName: "Johnson", Email: "alice.johnson@example.com"},
		{FirstName: "Bob", LastName: "Williams", Email: "bob.williams@example.com"},
		{FirstName: "Eva", LastName: "Miller", Email: "eva.miller@example.com"},
		{FirstName: "David", LastName: "Brown", Email: "david.brown@example.com"},
		{FirstName: "Catherine", LastName: "Jones", Email: "catherine.jones@example.com"},
		{FirstName: "Michael", LastName: "Wilson", Email: "michael.wilson@example.com"},
		{FirstName: "Olivia", LastName: "Anderson", Email: "olivia.anderson@example.com"},
		{FirstName: "Daniel", LastName: "Martinez", Email: "daniel.martinez@example.com"},
	}

	// Insert data into the database
	for _, developer := range developers {
		result := db.Create(&developer)
		if result.Error != nil {
			log.Fatalf("Error seeding data: %v", result.Error)
		}
	}

	log.Println("Seed data inserted successfully")
}
