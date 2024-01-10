package main

import (
	"net/http"

	"github.com/rsheldon3ayers/dumb-module/backend/internal/db"
)

func main() {

	// Initialize database
	db.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend is up and running!"))
	})

	http.ListenAndServe(":8080", nil)
}
