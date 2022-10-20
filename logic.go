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

func contains(i []int, intg int) bool {
	for _, v := range i {
		if v == intg {
			return true
		}
	}

	return false
}

func currentCol(i []int, intg int) int {
	for k, v := range i {
		if int(rl.GetMouseX())-v < intg {
			return k
		}
	}
	return -1
}

func (c *C4) logic() {
	c.menuLogic()

	if debugMenu {
		centerHelper()
		c.boardState()
		c.printDebug()
	}
}

// * main menu logic
func (c *C4) menuLogic() {
	if !gameOngoing {
		mainmenu()
		if rl.IsKeyPressed(rl.KeyEnter) || mouseButtonPressed {
			gameOngoing = true
		}
	} else if gameOngoing && !oponentSelected {
		c.oponentLogic()
	} else if gameOngoing && oponentSelected {
		c.boardLogic()
	}
}

// * select oponent menu logic
func (c *C4) oponentLogic() {
	if gameOngoing {
		oponentSelect()
		// Do a timer, when is a even sec blink, then hide.
		if shouldBlink {
			xPos := (width / 2) - (2 * gridSize)
			if isOponentAI {
				xPos += (4 * gridSize)
			}
			blink(xPos)
		}

		if rl.IsKeyPressed(rl.KeyEnter) || mouseButtonPressed {
			oponentSelected = true
		}
	}
}

func (c *C4) switchTurn() {
	if c.turn == c.P1.ID {
		c.turn = c.P2.ID
	} else if c.turn == c.P2.ID {
		c.turn = c.P1.ID
	}
}

func (c *C4) getSelectCol() {
	if boardVer.xPos+boardXtra < rl.GetMouseX() && rl.GetMouseX() < boardVer.xPos+boardVer.xMag-boardXtra {
		collCurrent = currentCol(collsPos, int(HEIGHT)*33/256)
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

		c.P1 = Gamer{ID: 1, CellColour: rl.Red}
		if !isOponentAI {
			c.P2 = Gamer{ID: 2, CellColour: rl.Yellow}
			c.gamers = []Gamer{c.P1, c.P2}
			c.turn = c.P1.ID
		} else if isOponentAI {
			c.AI = Gamer{ID: 3, CellColour: rl.Gold}
			c.gamers = []Gamer{c.P1, c.AI}
			c.turn = c.AI.ID
		}

		firstLoop = false
	}

	// render stuff
	if !isOponentAI {
		pvp()
	} else if isOponentAI {
		pvai()
	}
	board()

	// test render a cell
	c.board[ROWS-1][2] = c.P2.ID

	// ! The board starts at 400ish (boardXtra), increments on each cell by 96 (boardVer.xMag/7)

	// render every cell
	if gameOngoing && (!gameOver || !gameDraw) && oponentSelected && boardRendered {
		for row := range c.board {
			for col := range c.board[row] {
				cellState := c.board[row][col]
				if cellState == EMPTY {
					grid(col, row, rl.Black)
				} else {
					grid(col, row, c.gamers[cellState-1].CellColour)
				}
			}
		}

		c.floatingCell()
	}
}

func (c *C4) debugToggle() {
	if debugMenu {
		debugMenu = false
	} else if !debugMenu {
		debugMenu = true
	}
}