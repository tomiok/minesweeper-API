package main

import (
	"github.com/tomiok/minesweeper-API/api"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"os"
)

const defaultPort = "8080"

func main() {
	serverPort := os.Getenv("PORT")
	testEnvVariable := os.Getenv("TEST")
	logs.InitDefault("dev") //hardcoded at dev environment for the PoC API
	if serverPort == "" {
		serverPort = defaultPort
	}
	logs.Log().Info(testEnvVariable)
	api.Start(serverPort)
}
