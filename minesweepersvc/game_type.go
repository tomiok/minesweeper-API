package minesweepersvc

import (
	"github.com/google/uuid"
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
