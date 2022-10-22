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
		c.boardDebug()
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

func (c *C4) collFull() bool {
	collCurrent = currentCol(collsPos, int(HEIGHT)*33/256)
	if collHeight[collCurrent] < COLLUMNS-1 {
		return false
	} else if collHeight[collCurrent] >= COLLUMNS-1 {
		return true
	}
	return true
}

func (c *C4) makeMove() {
	if boardVer.xPos+boardXtra < rl.GetMouseX() && rl.GetMouseX() < boardVer.xPos+boardVer.xMag-boardXtra {
		if !c.collFull() {
			movesMade += 1
			collHeight[collCurrent] += 1
			c.board[ROWS-collHeight[collCurrent]][collCurrent] = c.turn
			c.detect4()
		}
	}
}

func (c *C4) detectXY4(sli []int32, state int32) bool {
	y = 0
	for _, x := range sli {
		if x == state {
			y += 1
			if y == 4 {
				return true
			}
		} else if x != state {
			y = 0
		}
	}
	return false
}

func (c *C4) detect4() {

	// This creates a 2d array based on colls to row, different from c.board that is from row to colls
	for y := 0; y < 6; y++ {
		tempSlice := []int32{}
		for x := 0; x < 6; x++ {
			tempSlice = append(tempSlice, c.board[x][y])
		}
		Colls2dSlice = append(Colls2dSlice, tempSlice)
	}
	for col := range Colls2dSlice {
		if c.detectXY4(Colls2dSlice[col], c.P1.ID) {
			winner = c.P1.ID
		} else if c.detectXY4(Colls2dSlice[col], c.P2.ID) {
			winner = c.P2.ID
		}
		gameOver = true


	}

	c.switchTurn()
}

func (c *C4) switchTurn() {
	if c.turn == c.P1.ID {
		c.turn = c.P2.ID
	} else if c.turn == c.P2.ID {
		c.turn = c.P1.ID
	}
}

func (c *C4) endGame() {
	for row := range c.board {
		for col := range c.board[row] {
			c.board[row][col] = EMPTY
		}
	}
	movesMade = 0
	collHeight = []int{0, 0, 0, 0, 0, 0, 0}
}

func (c *C4) backToMenu() {
	gameOngoing = false
	oponentSelected = false
	boardRendered = false
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
		c.pvp()
	} else if isOponentAI {
		pvai()
	}
	board()

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
