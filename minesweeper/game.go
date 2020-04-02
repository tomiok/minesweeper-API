package minesweeper

import (
	"github.com/google/uuid"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"github.com/tomiok/minesweeper-API/models"
)

const (
	defaultRows  = 6
	defaultCols  = 6
	defaultMines = 12
	maxMines     = 25
	maxRows      = 30
	maxCols      = 30
)

type GameService struct {
	store models.MineSweeperStorage
}

func (s *GameService) Create(game *models.Game) error {

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

	err := s.store.Create(game)
	return err
}

func (s *GameService) Start(name string) (*models.Game, error) {
	game, err := s.store.GetByName(name)
	if err != nil {
		return nil, err
	}

	buildBoard(game)

	game.Status = "started"
	err = s.store.Update(game)
	logs.Sugar().Infof("%#v\n", game.Grid)
	return game, err
}

func (s *GameService) Click(name string, i, j int) (*models.Game, error) {
	game, err := s.store.GetByName(name)
	if err != nil {
		return nil, err
	}

	if err := clickCell(game, i, j); err != nil {
		return nil, err
	}

	if err := s.store.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}

func getUUIDName() string {
	return uuid.New().String()
}
