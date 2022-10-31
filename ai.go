package main

import (
	// "math"
	"math/rand"
	"time"
)

// rl "github.com/gen2brain/raylib-go/raylib"

func (c *C4) aiMove() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	col := r.Intn(COLLUMNS)
	row := 5 - c.board.heights[col]

	if row != -1 { // check if col is ful
		movesMade += 1
		c.board.heights[col] += 1
		c.board.grid[row][col] = c.P2.ID
		c.switchTurn()

	}
}

// func (c *C4) bestMove(passedBoard [][]int32) int32 {

// 	colScoreBest := math.Inf(-1)

// 	for col := range collHeight {
// 		row := 5 - collHeight[col]
// 		if row != -1 {
// 			// creates a copy of the original board, independant of it

// 			board := make([][]int32, len(passedBoard))
// 			for i := range board {
// 				board[i] = make([]int32, len(passedBoard[i]))
// 				copy(board[i], passedBoard[i])
// 			}

// 			colSelected := collHeight[col]
// 			board[row][colSelected] = c.P2.ID

// 			// score the board
// 			score := 0
// 			switch c.detect4(c.board.grid) {
// 			case 0:
// 				if movesMade == 42 {
// 					score = 0
// 				}
// 			case c.P1.ID:
// 				score += -1_000_000_000
// 			case c.P2.ID:
// 				score += 1_000_000_000
// 			}

// 			// if the score gotten is > bestScore { bestScore = score, colSelected = }
// 		}
// 	}
// }
