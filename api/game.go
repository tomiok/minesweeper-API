package api

import (
	"encoding/json"
	"github.com/tomiok/minesweeper-API/models"
	"net/http"
)

func (s *Services) createGameHandler(w http.ResponseWriter, r *http.Request) {
	var game models.Game
	body := r.Body
	defer body.Close()
	if err := json.NewDecoder(body).Decode(&game); err != nil {

	}
}

func (s *Services) createUserHandler(w http.ResponseWriter, r *http.Request) {

}

func (s *Services) startGame(w http.ResponseWriter, r *http.Request) {

}

func (s *Services) clickHandler(w http.ResponseWriter, r *http.Request) {

}