package api

import (
	"github.com/tomiok/minesweeper-API/minesweepersvc"
)

type Services struct {
	gameService minesweepersvc.MineSweeperGameService
	userService minesweepersvc.MineSweeperUserService
}

func Start(port string) {
	db := minesweepersvc.New()
	services := &Services{
		gameService: minesweepersvc.NewGameService(db),
		userService: minesweepersvc.NewUserService(db),
	}

	r := routes(services)
	server := newServer(port, r)
	server.Start()
}
