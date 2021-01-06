package minesweepersvc

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tomiok/minesweeper-API/internal/logs"
	"math"
	"time"
)

// some default game parameters if the user does not provide those.
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

var gameStatus status

type status struct{}
type statuses func() string
type CellGrid []Cell

//Cell express the status of every cell in the board.
type Cell struct {
	Mine        bool   `json:"mine"`
	Clicked     bool   `json:"clicked"`
	Flagged     bool   `json:"flagged"` // add a red flag in the cell
	Marked      bool   `json:"marked"`  // add a question mark int the cell
	Coordinates string `json:"coordinates"`
}

//Game contains all the structure of the game, user, results, parameters, etc.
type Game struct {
	Name         string        `json:"name"`
	Rows         int           `json:"rows"`
	Cols         int           `json:"cols"`
	Mines        int           `json:"mines"`
	Status       statuses      `json:"-"` //new, in_progress, won, lost
	S            string        `json:"status"`
	Grid         []CellGrid    `json:"grid,omitempty"`
	ClickCounter int           `json:"-"`
	Username     string        `json:"username"`
	CreatedAt    time.Time     `json:"created_at,omitempty"`
	StartedAt    time.Time     `json:"-"`
	TimeSpent    time.Duration `json:"time_spent"`
	Points       float32       `json:"points,omitempty"`
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
	GameStorage MineSweeperGameStorage
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
	err = s.GameStorage.Update(game)
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

	if err := s.GameStorage.Update(game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *MSGameService) FlushAll() error {
	return s.GameStorage.FlushAll()
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

//after 15 seconds the time matters.
func (game *Game) calculateScoring() float64 {
	seconds := game.TimeSpent.Seconds()
	totalGrid := game.Rows * game.Cols
	clicks := game.ClickCounter
	relativeValue := 100.0
	if seconds > 15 {
		u := seconds / 10
		positiveClicks := totalGrid - clicks
		a := float64(positiveClicks) / relativeValue
		return math.Dim(a, u)
	}

	positiveClicks := totalGrid - clicks
	return float64(positiveClicks) / relativeValue
}
