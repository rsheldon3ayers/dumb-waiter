package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/gorm"

	"github.com/rsheldon3ayers/dumb-module/backend/internal/models"
)

type handler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handler {
	return handler{db}
}

// home handler
func (h handler) Home(w http.ResponseWriter, r *http.Request) {
	log.Println("Backend API is running!")
}
func (h handler) GetAllDevelopers(w http.ResponseWriter, r *http.Request) {
	developers := []models.Developer{}
	h.DB.Find(&developers)
	json.NewEncoder(w).Encode(developers)
}
