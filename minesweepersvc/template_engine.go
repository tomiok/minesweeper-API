package minesweepersvc

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func buildBoard(game *Game) {
	numCells := game.Cols * game.Rows
	cells := make(CellGrid, numCells)

	// Randomly set mines
	i := 0
	for i < game.Mines {
		idx := rand.Intn(numCells)
		if !cells[idx].Mine {
			cells[idx].Mine = true
			i++
		}
	}

	game.Grid = make([]CellGrid, game.Rows)
	for row := range game.Grid {
		game.Grid[row] = cells[(game.Cols * row):(game.Cols * (row + 1))]
	}

	for row, col := range game.Grid {
		for c := range col {
			game.Grid[row][c].Coordinates = fmt.Sprintf("row: %d, col:%d", row, c)
		}
	}
}

func clickCell(game *Game, i, j int) error {
	if clicked(game, i, j) {
		return errors.New("cell already clicked")
	}
	game.Grid[i][j].Clicked = true
	if game.Grid[i][j].Mine {
		game.Status = gameStatus.lost
		game.S = game.Status()
		// just lost the game, is not an error
		return nil
	}
	game.ClickCounter += 1

	if checkWon(game) {
		game.Status = gameStatus.won
		game.S = game.Status()
		game.calculateScoring()

		return nil
	}

	return nil
}

func flagOrQuestionMarkCell(game *Game, i, j int, clickType string) error {
	if clicked(game, i, j) {
		return errors.New("cell already clicked")
	}
	switch clickType {
	// assume that a flagged cell could change to marked or vice versa
	case "flag":
		if game.Grid[i][j].Flagged {
			return errors.New("cell already flagged")
		}
		game.Grid[i][j].Flagged = true
	case "mark":
		if game.Grid[i][j].Marked {
			return errors.New("cell already marked")
		}
		game.Grid[i][j].Marked = true
	default:
		return errors.New("unknown event")
	}

	//cannot check won, you need to click those not only mark or flag
	return nil
}

// checkWon if the user clicked all the cells without any mine
func checkWon(game *Game) bool {
	if game.Status == nil {
		return false
	}
	started := game.Status() == gameStatus.inProgress()
	return game.ClickCounter == ((game.Rows*game.Cols)-game.Mines) && started
}

func clicked(game *Game, i, j int) bool {
	return game.Grid[i][j].Clicked
}
