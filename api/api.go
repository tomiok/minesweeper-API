package api

import (
	"github.com/tomiok/minesweeper-API/minesweepersvc"
)

type Services struct {
	gameService minesweepersvc.MineSweeperGameService
}

func Start(port string) {
	db := minesweepersvc.NewDB()
	services := &Services{
		gameService: minesweepersvc.NewGameService(db),
	}

	r := routes(services)
	server := newServer(port, r)
	server.Start()
}
