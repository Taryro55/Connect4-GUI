package main

import (
	"math/rand"
	"time"
)

// rl "github.com/gen2brain/raylib-go/raylib"

func (c *C4) aiMove() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	col := r.Intn(COLLUMNS)
	row := 5 - collHeight[col]

	if row != -1 { // check if col is ful
		movesMade += 1
		collHeight[col] += 1
		c.board[row][col] = c.P2.ID
		c.detect4()
	}

	// movesMade += 1
	// collHeight[collCurrent] += 1
	// c.board[ROWS-collHeight[collCurrent]][collCurrent] = c.turn
	// c.detect4()
}
