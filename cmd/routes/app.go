package routes

import (
	"github.com/Fasunle/golang-starters.git/cmd/handlers"
	"github.com/go-chi/chi/v5"
)

func appRoutes(r *chi.Mux) {
	r.Get("/", handlers.HomeHandler)
}
