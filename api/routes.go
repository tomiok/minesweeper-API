package api

import (
	"github.com/go-chi/chi"
)

func routes(services *Services) *chi.Mux {
	r := chi.NewMux()

	r.Get("/heartbeat", services.healthCheck)
	r.Post("/users", services.createUserHandler)
	r.Post("/games", services.createGameHandler)
	r.Post("/games/{gameID}/users/{username}", services.startGame)
	r.Post("/games/{gameID}/users/{username}/click", services.clickHandler)

	return r
}
