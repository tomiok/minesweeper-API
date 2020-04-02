package main

import (
	"github.com/tomiok/minesweeper-API/api"
	"github.com/tomiok/minesweeper-API/internal/logs"
)

const defaultPort = "8080"

func main() {
	logs.InitDefault("dev") //hardcoded at dev environment for the PoC API
	api.Start(defaultPort)
}
