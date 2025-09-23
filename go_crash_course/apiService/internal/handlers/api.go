package handlers

import (
	"apiService/internal/middleware"

	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

// function name starting with Uppercase meams public else private.
func Handler(r *chi.Mux) {
	r.Use(chimiddle.StripSlashes) // prevents 404 error by stripping / from last

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})

}


