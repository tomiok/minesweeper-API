package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(services *Services) *chi.Mux {
	r := chi.NewMux()

	r.Use(middleware.Recoverer)

	r.Get("/heartbeat", services.healthCheck)
	r.Post("/users", services.createUserHandler)
	r.Post("/games", services.createGameHandler)
	r.Post("/games/{gameID}/users/{username}", services.startGameHandler)
	r.Post("/games/{gameID}/users/{username}/click", services.clickHandler)
	r.Get("/configs/flush-cache", services.flushCacheHandler)

	return r
}
