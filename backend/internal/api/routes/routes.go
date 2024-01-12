// Routes for the API
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rsheldon3ayers/dumb-module/backend/internal/api/handlers"
	"github.com/rsheldon3ayers/dumb-module/backend/internal/db"
)

func Routes() http.Handler {
	DB := db.Init()
	h := handlers.New(DB)

	r := mux.NewRouter()

	r.HandleFunc("/", h.Home).Methods(http.MethodGet)

	r.HandleFunc("/developers", h.GetAllDevelopers).Methods("GET")
	return r
}
