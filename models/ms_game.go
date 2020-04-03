package models

import "time"

type MineSweeperService interface {
	CreateGame(game *Game) error
	Start(name string) (*Game, error)
	Click(name, clickType string, i, j int) (*Game, error)
}

type MineSweeperStorage interface {
	Create(game *Game) error
	Update(game *Game) error
	GetByName(name string) (*Game, error)
}

type UserStorage interface {
	Save(u *User) error
	GetByName(name string) (*User, error)
}

type Cell struct {
	Mine    bool `json:"mine"`
	Clicked bool `json:"clicked"`
	Value   int  `json:"value"`
	Flagged bool `json:"flagged"` // add a red flag in the cell
	Marked  bool `json:"marked"`  // add a question mark
}

type CellGrid []Cell

type Game struct {
	Name         string     `json:"name"`
	Rows         int        `json:"rows"`
	Cols         int        `json:"cols"`
	Mines        int        `json:"mines"`
	Status       string     `json:"status"` //new, in_progress, won, lost
	Grid         []CellGrid `json:"grid,omitempty"`
	ClickCounter int        `json:"-"`
	Username     string     `json:"username"`
}

type User struct {
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"-"`
}
