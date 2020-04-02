package api

import (
	"github.com/go-chi/chi"
)

func routes(services *Services) *chi.Mux {
	r := chi.NewMux()

	r.Get("/heartbeat", services.healthCheck)

	return r
}
