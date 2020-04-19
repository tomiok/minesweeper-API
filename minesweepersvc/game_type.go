package minesweepersvc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"time"
)

const (
	defaultRows  = 6
	defaultCols  = 6
	defaultMines = 12
	maxMines     = 25
	maxRows      = 30
	maxCols      = 30
	Click        = "click"
	InProgress   = "in_progress"
	Lost         = "lost"
	New          = "new"
	Won          = "won"
)

type status struct{}

var gameStatus status

type statuses func() string
type CellGrid []Cell

type Cell struct {
	Mine        bool   `json:"mine"`
	Clicked     bool   `json:"clicked"`
	Flagged     bool   `json:"flagged"` // add a red flag in the cell
	Marked      bool   `json:"marked"`  // add a question mark
	Coordinates string `json:"coordinates"`
}

type Game struct {
	Name         string     `json:"name"`
	Rows         int        `json:"rows"`
	Cols         int        `json:"cols"`
	Mines        int        `json:"mines"`
	Status       statuses   `json:"-"` //new, in_progress, won, lost
	S            string     `json:"status"`
	Grid         []CellGrid `json:"grid,omitempty"`
	ClickCounter int        `json:"-"`
	Username     string     `json:"username"`
}

type User struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"-"`
}

type ClickAction struct {
	Row       int    `json:"row"`
	Col       int    `json:"col"`
	ClickType string `json:"click_type"`
}

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
	Update(game *Game) error
	GetGame(name string) (*Game, error)
	CreateUser(u *User) error
	GetUser(username string) (*User, error)
	FlushAll() error
}

type MSGameService struct {
	gameStorage MineSweeperGameStorage
}

func (s status) new() string {
	return New
}

func (s status) won() string {
	return Won
}

func (s status) inProgress() string {
	return InProgress
}

func (s status) lost() string {
	return Lost
}

func NewGameService(db DB) MineSweeperGameService {
	return &MSGameService{
		gameStorage: NewGameEngineStorage(db),
	}
}

func (s *MSGameService) GetUser(username string) (*User, error) {
	return s.gameStorage.GetUser(username)
}

func (s *MSGameService) CreateUser(user *User) error {
	return s.gameStorage.CreateUser(user)

}

func (s *MSGameService) CreateGame(game *Game) error {
	username := game.Username
	if username == "" {
		return errors.New("username empty is not allowed")
	}

	_, err := s.gameStorage.GetUser(username)

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

	err = s.gameStorage.CreateGame(game)
	return err
}

func (s *MSGameService) Start(name string) (*Game, error) {
	game, err := s.gameStorage.GetGame(name)
	if err != nil {
		return nil, err
	}

	buildBoard(game)

	game.Status = gameStatus.inProgress
	game.S = game.Status()
	err = s.gameStorage.Update(game)
	logs.Sugar().Infof("%#v\n", game.Grid)
	return game, err
}

func (s *MSGameService) Click(name, clickType string, i, j int) (*Game, error) {
	game, err := s.gameStorage.GetGame(name)
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

func (s *MSGameService) FlushAll() error {
	return s.gameStorage.FlushAll()
}

func isNormalClick(clickType string) bool {
	return clickType == Click
}

func getUUIDName() string {
	return uuid.New().String()
}

func CheckLost(status string) bool {
	return status == "lost"
}
