package api

import (
	"github.com/tomiok/minesweeper-API/minesweeper"
	"github.com/tomiok/minesweeper-API/models"
	"github.com/tomiok/minesweeper-API/storage"
)

type Services struct {
	gameService models.MineSweeperService
}

func Start(port string) {
	db := storage.New()
	services := &Services{
		gameService: minesweeper.NewGameService(db),
	}

	r := routes(services)
	server := newServer(port, r)
	server.Start()
}
