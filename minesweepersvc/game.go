// All the interfaces implementations start with MS (for minesweepersvc)
package minesweepersvc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tomiok/minesweeper-API/internal/logs"
)

const (
	defaultRows  = 6
	defaultCols  = 6
	defaultMines = 12
	maxMines     = 25
	maxRows      = 30
	maxCols      = 30
)

type MSGameService struct {
	gameStorage MineSweeperGameStorage
	userService MineSweeperUserService
}

type MSUserService struct {
	MineSweeperUserStorage
}

func (u *MSUserService) CreateUser(user *User) error {
	return u.MineSweeperUserStorage.Create(user)
}

func (u *MSUserService) GetUserByName(name string) (*User, error) {
	return u.MineSweeperUserStorage.GetByName(name)
}

func NewGameService(db *DB) MineSweeperGameService {
	return &MSGameService{
		gameStorage: NewGameEngineStorage(db),
		userService: &MSUserService{NewUserStorage(db)},
	}
}

func NewUserService(db *DB) MineSweeperUserService {
	return &MSUserService{
		NewUserStorage(db),
	}
}

func (s *MSGameService) CreateGame(game *Game) error {
	username := game.Username
	if username == "" {
		return errors.New("username empty is not allowed")
	}

	_, err := s.userService.GetUserByName(username)

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
	game.Status = "new"

	err = s.gameStorage.Create(game)
	return err
}

func (s *MSGameService) Start(name string) (*Game, error) {
	game, err := s.gameStorage.GetByName(name)
	if err != nil {
		return nil, err
	}

	buildBoard(game)

	game.Status = "in_progress"
	err = s.gameStorage.Update(game)
	logs.Sugar().Infof("%#v\n", game.Grid)
	return game, err
}

func (s *MSGameService) Click(name, clickType string, i, j int) (*Game, error) {
	game, err := s.gameStorage.GetByName(name)
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

	if err := s.gameStorage.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}

func isNormalClick(clickType string) bool {
	return clickType == "click"
}

func getUUIDName() string {
	return uuid.New().String()
}
