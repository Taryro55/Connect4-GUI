package main

import (
	"fmt"
	"math/rand"
	"time"
)

// rl "github.com/gen2brain/raylib-go/raylib"

func (c *C4) aiMove() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	col := r.Intn(COLLUMNS)
	
	if (0 <= col && col <= 6) { // and check if col is ful
		fmt.Println(r.Intn(col))
	}

	// movesMade += 1
	// collHeight[collCurrent] += 1
	// c.board[ROWS-collHeight[collCurrent]][collCurrent] = c.turn
	// c.detect4()
}
