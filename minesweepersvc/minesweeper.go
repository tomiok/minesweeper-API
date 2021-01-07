package minesweepersvc

import (
	"errors"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"time"
)

type MineSweeperGameService interface {
	CreateGame(game *Game) error
	Start(name string) (*Game, error)
	Click(name, clickType string, i, j int) (*Game, error) //click type [click, flag, mark]
	FlushAll() error
	CreateUser(user *User) error
	GetUser(username string) (*User, error)
}

type MineSweeperGameStorage interface {
	CreateGame(game *Game) error
	UpdateGame(game *Game) error
	GetGame(name string) (*Game, error)
	CreateUser(u *User) error
	GetUser(username string) (*User, error)
	FlushAll() error // for development purpose
}

type MSGameService struct {
	GameStorage MineSweeperGameStorage
}

func NewGameService(db DB) MineSweeperGameService {
	return &MSGameService{
		GameStorage: NewGameEngineStorage(db),
	}
}

func (s *MSGameService) GetUser(username string) (*User, error) {
	return s.GameStorage.GetUser(username)
}

func (s *MSGameService) CreateUser(user *User) error {
	return s.GameStorage.CreateUser(user)

}

func (s *MSGameService) CreateGame(game *Game) error {
	username := game.Username
	if username == "" {
		return errors.New("username empty is not allowed")
	}

	_, err := s.GameStorage.GetUser(username)

	if err != nil {
		return errors.New("user_not_found")
	}

	if game.Name == "" {
		game.Name = getUUIDName()
	}

	if game.Rows == 0 {
		game.Rows = defaultRows
	}

	if game.Cols == 0 {
		game.Cols = defaultCols
	}

	if game.Mines == 0 {
		game.Mines = defaultMines
	}

	if game.Mines > maxMines {
		game.Mines = maxMines
	}

	if game.Rows > maxRows {
		game.Rows = maxRows
	}
	if game.Cols > maxCols {
		game.Cols = maxCols
	}
	if game.Mines > (game.Cols * game.Rows) {
		game.Mines = game.Cols * game.Rows
	}
	game.Status = gameStatus.new
	game.S = game.Status()
	game.CreatedAt = time.Now()

	err = s.GameStorage.CreateGame(game)
	return err
}

func (s *MSGameService) Start(name string) (*Game, error) {
	game, err := s.GameStorage.GetGame(name)
	if err != nil {
		return nil, err
	}
	buildBoard(game)

	game.Status = gameStatus.inProgress
	game.S = game.Status()
	game.StartedAt = time.Now()
	err = s.GameStorage.UpdateGame(game)
	logs.Sugar().Infof("%#v\n", game.Grid)
	return game, err
}

func (s *MSGameService) Click(name, clickType string, i, j int) (*Game, error) {
	game, err := s.GameStorage.GetGame(name)
	if err != nil {
		return nil, err
	}

	if isNormalClick(clickType) {
		if err := clickCell(game, i, j); err != nil {
			return nil, err
		}
	} else {
		if err := flagOrQuestionMarkCell(game, i, j, clickType); err != nil {
			return nil, err
		}
	}

	//some millis are lost here :(
	game.TimeSpent = time.Now().Sub(game.StartedAt)
	if err := s.GameStorage.UpdateGame(game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *MSGameService) FlushAll() error {
	return s.GameStorage.FlushAll()
}
