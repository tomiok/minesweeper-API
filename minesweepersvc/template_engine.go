package minesweepersvc

import (
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
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

	// Set cell values
	for i, row := range game.Grid {
		for j, cell := range row {
			if cell.Mine {
				setAdjacentValues(game, i, j)
			}
		}
	}
}

func setAdjacentValues(game *Game, i, j int) {
	for z := i - 1; z < i+2; z++ {
		if z < 0 || z > game.Rows-1 {
			continue
		}
		for w := j - 1; w < j+2; w++ {
			if w < 0 || w > game.Cols-1 {
				continue
			}
			if z == i && w == j {
				continue
			}
			game.Grid[z][w].Value++
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
		return nil
	}
	game.ClickCounter += 1
	if checkWon(game) {
		game.Status = gameStatus.won
		game.S = game.Status()
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

func checkWon(game *Game) bool {
	started := game.Status() == gameStatus.inProgress()
	return game.ClickCounter == ((game.Rows*game.Cols)-game.Mines) && started
}

func clicked(game *Game, i, j int) bool {
	return game.Grid[i][j].Clicked
}
