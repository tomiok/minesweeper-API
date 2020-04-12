package minesweepersvc

import (
	"errors"
)

type GameEngineStorage struct {
	db DB
}

type UserStorage struct {
	db DB
}

func NewGameEngineStorage(db DB) *GameEngineStorage {
	return &GameEngineStorage{db: db}
}

func NewUserStorage(db DB) *UserStorage {
	return &UserStorage{db: db}
}

func (s *UserStorage) GetByName(username string) (*User, error) {
	user, err := s.db.Get(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserStorage) Create(user *User) error {
	//TODO add exists here
	if err := s.db.Save(user.Username, user); err != nil {
		return errors.New("persistence error")
	}
	return nil
}

func (s *GameEngineStorage) Create(game *Game) error {
	//TODO add exists here
	if err := s.db.Save(game.Name, game); err != nil {
		return errors.New("persistence error")
	}
	return nil
}

func (s *GameEngineStorage) Update(game *Game) error {
	if _, err := s.db.Get(game.Name); err != nil {
		return errors.New("game do not exist")
	}
	return s.db.Save(game.Name, game)
}

func (s *GameEngineStorage) GetByName(name string) (*Game, error) {
	game, err := s.db.GetGame(name)
	if err != nil {
		return nil, errors.New("game not found")
	}
	return game, nil

}
