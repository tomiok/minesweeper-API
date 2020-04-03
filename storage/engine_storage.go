package storage

import (
	"errors"
	"github.com/tomiok/minesweeper-API/models"
)

type GameEngineStorage struct {
	db *DB
}

type UserStorage struct {
	db *DB
}

func NewGameEngineStorage(db *DB) *GameEngineStorage {
	return &GameEngineStorage{db: db}
}

func NewUserStorage(db *DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) GetByName(username string) (*models.User, error) {
	if user, ok := s.db.users[username]; ok {
		u := *user
		return &u, nil
	}
	return nil, errors.New("game not found")
}


func (s *UserStorage) Create(user *models.User) error {
	if _, ok := s.db.users[user.Username]; ok {
		return errors.New("user already exist")
	}
	s.db.users[user.Username] = user
	return nil
}

func (s *GameEngineStorage) Create(game *models.Game) error {
	if _, ok := s.db.games[game.Name]; ok {
		return errors.New("game already exist")
	}
	s.db.games[game.Name] = game
	return nil
}

func (s *GameEngineStorage) Update(game *models.Game) error {
	g := *game
	if _, ok := s.db.games[game.Name]; !ok {
		return errors.New("game do not exist")
	}
	s.db.games[game.Name] = &g
	return nil
}

func (s *GameEngineStorage) GetByName(name string) (*models.Game, error) {
	if game, ok := s.db.games[name]; ok {
		g := *game
		return &g, nil
	}
	return nil, errors.New("game not found")
}
