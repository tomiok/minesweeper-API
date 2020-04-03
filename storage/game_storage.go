package storage

import "github.com/tomiok/minesweeper-API/models"

type DB struct {
	games map[string]*models.Game
	users map[string]*models.User
}

func New() *DB {
	return &DB{
		games: make(map[string]*models.Game),
		users: make(map[string]*models.User),
	}
}
