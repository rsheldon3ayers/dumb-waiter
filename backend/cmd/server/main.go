package main

import (
	"net/http"

	"github.com/rsheldon3ayers/dumb-module/backend/internal/api/routes"
)

func main() {
	r := routes.Routes()
	http.ListenAndServe(":8080", r)
}
