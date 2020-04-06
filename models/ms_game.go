package models

import "time"

type MineSweeperGameService interface {
	CreateGame(game *Game) error
	Start(name string) (*Game, error)
	Click(name, clickType string, i, j int) (*Game, error) //click type [click, flag, mark]
}

type MineSweeperGameStorage interface {
	Create(game *Game) error
	Update(game *Game) error
	GetByName(name string) (*Game, error)
}

type MineSweeperUserService interface {
	CreateUser(u *User) error
	GetUserByName(username string) (*User, error)
}

type MineSweeperUserStorage interface {
	Create(u *User) error
	GetByName(username string) (*User, error)
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

type ClickAction struct {
	Row       int    `json:"row"`
	Col       int    `json:"col"`
	ClickType string `json:"click_type"`
}
