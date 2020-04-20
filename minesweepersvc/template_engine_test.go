package minesweepersvc

import (
	"testing"
)

var game = Game{
	Name:         "test_game",
	Rows:         10,
	Cols:         10,
	Mines:        20,
	Status:       nil,
	Grid:         nil,
	ClickCounter: 0,
	Username:     "tomas",
}

func Test_CheckWon_notPlayed(t *testing.T) {
	g := game
	won := checkWon(&g)
	if won {
		t.Fail()
		t.Error("game is not even started")
	}
}

func Test_CheckWon_notStarted(t *testing.T) {
	g := game
	g.ClickCounter = 80
	won := checkWon(&g)
	if won {
		t.Fail()
		t.Error("should be a victory, but the game is not started")
	}
}

func Test_CheckWon_started(t *testing.T) {
	g := game
	g.ClickCounter = 80
	g.Status = gameStatus.inProgress
	g.S = g.Status()
	won := checkWon(&g)
	if !won {
		t.Fail()
		t.Error("should be a victory")
	}
}

func Test_buildBoard(t *testing.T) {
	g := game
	buildBoard(&g)

	if g.Grid == nil {
		t.Fail()
	}
}

func Test_clickCell(t *testing.T) {
	g := game
	g.Mines = 0 // cannot add mines for this test
	buildBoard(&g)
	i, j := 1, 1

	if err := clickCell(&g, i, j); err != nil {
		t.Fail()
		t.Error(err.Error())
	}

	//click same cell
	if err := clickCell(&g, i, j); err == nil {
		t.Fail()
		t.Error("should fail here")
	}

}
