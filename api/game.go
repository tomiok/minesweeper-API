package api

import (
	"encoding/json"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"github.com/tomiok/minesweeper-API/models"
	"go.uber.org/zap"
	"net/http"
)

func (s *Services) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	body := r.Body
	defer body.Close()

	logs.Log().Info("create a new game")

	if err := json.NewDecoder(body).Decode(&game); err != nil {
		logs.Log().Error("cannot parse the request", zap.Error(err))
		ErrInvalidJSON.Send(w)
		return
	}

	if err := s.gameService.CreateGame(&game); err != nil {
		logs.Log().Error("cannot create the game", zap.Error(err))
		ErrBadRequest.Send(w)
		return
	}

	Success(game, http.StatusCreated)
}

func (s *Services) createUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	body := r.Body
	defer body.Close()

	logs.Log().Info("creating a new user")

	if err := json.NewDecoder(body).Decode(&user); err != nil {
		logs.Log().Error("cannot parse the request", zap.Error(err))
		ErrInvalidJSON.Send(w)
		return
	}

	if err := s.userService.CreateUser(&user); err != nil {
		logs.Log().Error("cannot create user", zap.Error(err))
		ErrBadRequest.Send(w)
		return
	}

	Success(&user, http.StatusCreated)
}

func (s *Services) startGame(w http.ResponseWriter, r *http.Request) {
	username := ""
	game := ""
	s.userService.GetUserByName(username)
	s.gameService.Start(game)
}

func (s *Services) clickHandler(w http.ResponseWriter, r *http.Request) {

}
