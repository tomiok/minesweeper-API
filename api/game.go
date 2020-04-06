package api

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"github.com/tomiok/minesweeper-API/models"
	"go.uber.org/zap"
	"net/http"
	"time"
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

	Success(game, http.StatusCreated).Send(w)
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

	user.CreatedAt = time.Now()
	if err := s.userService.CreateUser(&user); err != nil {
		logs.Log().Error("cannot create user", zap.Error(err))
		ErrAlreadyExists.Send(w)
		return
	}

	Success(&user, http.StatusCreated).Send(w)
}

func (s *Services) startGame(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	gameID := chi.URLParam(r, "gameID")

	if _, err := s.userService.GetUserByName(username); err != nil {
		logs.Log().Error("user is not present", zap.Error(err))
		ErrBadRequest.Send(w)
		return
	}

	game, err := s.gameService.Start(gameID)
	if err != nil {
		logs.Log().Error("cannot start game	", zap.Error(err))
		ErrBadRequest.Send(w)
		return
	}

	Success(game, http.StatusOK).Send(w)
}

func (s *Services) clickHandler(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	gameID := chi.URLParam(r, "gameID")

	body := r.Body
	defer body.Close()

	if _, err := s.userService.GetUserByName(username); err != nil {
		logs.Log().Error("user is not present", zap.Error(err))
		ErrBadRequest.Send(w)
		return
	}

	var clickAction models.ClickAction
	if err := json.NewDecoder(body).Decode(&clickAction); err != nil {
		ErrInvalidJSON.Send(w)
		return
	}

	// always keep the order ROW-COL
	game, err := s.gameService.Click(gameID, clickAction.ClickType, clickAction.Row, clickAction.Col)

	if err != nil {
		logs.Log().Error("cannot click the current cell", zap.Error(err))
		ErrWrongClick.Send(w)
		return
	}

	type Res struct {
		Game    *models.Game        `json:"game"`
		Clicked *models.ClickAction `json:"clicked"`
	}

	res := Res{
		Game:    game,
		Clicked: &clickAction,
	}

	Success(&res, http.StatusOK).Send(w)
}