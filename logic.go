package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)


func isEven(l int) bool {
	if l < 2 {
		return l%2 == 0
	}
	return isEven(l - 2)
}


func (c *C4) logic() {
	c.menuLogic()
}


// * main menu logic
func (c *C4) menuLogic() {
	if !gameOngoing {
		mainmenu()
		if (rl.IsKeyPressed(rl.KeyEnter) || mouseButtonPressed) {
			gameOngoing = true
		}
	} else if (gameOngoing && !oponentSelected) {
		c.oponentLogic()
	} else if (gameOngoing && oponentSelected) {
		c.boardLogic()
	}
}


// * select oponent menu logic
func (c *C4) oponentLogic() {
	if gameOngoing {
		oponentSelect()
		// Do a timer, when is a even sec blink, then hide.
		if shouldBlink {
			xPos := (width/2)-(2*gridSize)
			if isOponentAI {
				xPos += (4*gridSize)
			}
			blink(xPos)
		}

		if (rl.IsKeyPressed(rl.KeyEnter) || mouseButtonPressed) {
			oponentSelected = true
		}
	}
}


// * board logic
func (c *C4) boardLogic() {
	if firstLoop {
		// board
		c.board = make([][]int32, ROWS)
		for row := range c.board {
			c.board[row] = make([]int32, COLLUMNS)
		}
		// Visualise board
		for row := range c.board {
			fmt.Println(c.board[row])
		}
		firstLoop = false
	}

	// render calls
	if !isOponentAI {
		pvp()
	} else if isOponentAI {
		pvai()
	}
	board()


	// render every cell
	if gameOngoing && (!gameOver || !gameDraw) && oponentSelected && boardRendered {
		for row := range c.board {
			for col := range c.board[row] {
				if c.board[row][col] == EMPTY {
					grid(col, row)
				}
			}
		}
	}
}
