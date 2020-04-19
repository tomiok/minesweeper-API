package minesweepersvc

import (
	"errors"
)

type GameEngineStorage struct {
	db DB
}

func NewGameEngineStorage(db DB) *GameEngineStorage {
	return &GameEngineStorage{db: db}
}

func (s *GameEngineStorage) GetUser(username string) (*User, error) {
	user, err := s.db.Get(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *GameEngineStorage) CreateUser(user *User) error {
	if s.db.Exists(user.Username) {
		return errors.New("already_exists")
	}

	if err := s.db.Save(user.Username, user); err != nil {
		return errors.New("persistence error")
	}
	return nil
}

func (s *GameEngineStorage) CreateGame(game *Game) error {
	if s.db.Exists(game.Name) && game.S == "in_progress" {
		return errors.New("already_exists")
	}

	if err := s.db.Save(game.Name, game); err != nil {
		return errors.New("persistence_error")
	}
	return nil
}

func (s *GameEngineStorage) Update(game *Game) error {
	if _, err := s.db.Get(game.Name); err != nil {
		return errors.New("game do not exist")
	}
	return s.db.Save(game.Name, game)
}

func (s *GameEngineStorage) GetGame(name string) (*Game, error) {
	game, err := s.db.GetGame(name)
	if err != nil || game == nil {
		return nil, errors.New("game not found")
	}
	return game, nil
}

func (s *GameEngineStorage) FlushAll() error {
	return s.db.FlushAll()
}
