package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// * General logic funcs

// Checks if integer is even, returns true
func isEven(l int) bool {
	if l < 2 {
		return l%2 == 0
	}
	return isEven(l - 2)
}

// Checks if a []int contains a int, returns true
func contains(i []int32, intg int32) bool {
	for _, v := range i {
		if v == intg {
			return true
		}
	}

	return false
}

// * Gameplay Logic

// handles turn switching
func (c *C4) switchTurn() {
	getWinner(c.board)
	if gameWinner == 0 {
		if c.turn == c.P1.ID {
			c.turn = c.P2.ID
		} else if c.turn == c.P2.ID {
			c.turn = c.P1.ID
		}
	}
}

func (c *C4) makeMove() {
	collCurrent = getHoveringCol(collsPos, height*33/256, rl.GetMouseX())
	if cursorOverBoard && isColValid(collCurrent, c.board) {
		movesMade += 1
		c.board.heights[collCurrent] += 1
		c.board.grid[ROWS-c.board.heights[collCurrent]][collCurrent] = c.turn
		c.switchTurn()
	}
}
