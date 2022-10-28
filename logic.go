package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (c *C4) logic() {
	// General logic func calls
	if screenMenu {
		c.mainMenuLogic()
	} else if !screenMenu && screenConf && !screenOponent {
		c.configMenuLogic()
	} else if !screenMenu && !screenConf && screenOponent {
		c.oponentMenuLogic()
	} else if !screenMenu && !screenConf && !screenOponent && screenBoard {
		c.boardMenuLogic()
	}

	// Timer based on amount of loops
	if mainLoops == 6 {
		mainLoops = 0
		runningMilisecsTimesTen += 1
		runningMilisecs = float32(runningMilisecsTimesTen) / 10
	}

	// Activates debug if needed.
	if debugMenu {
		c.boardDebug()
	}

}

// * Menus Logic

// main menu logic
func (c *C4) mainMenuLogic() {
	if screenMenu {
		mainmenu()
		if continuePressed {
			screenMenu = false
			screenOponent = true
		} else if rl.IsKeyPressed(CONF_KEY) {
			screenMenu = false
			screenConf = true
		}
	}
}

func (c *C4) configMenuLogic() {
	if screenConf {
		confmenu()

		if rl.IsKeyPressed(rl.KeyDown) && configHover != 2 {
			configHover += 1
		} else if rl.IsKeyPressed(rl.KeyUp) && configHover != 0 {
			configHover -= 1
		}

		if rl.IsKeyPressed(CONF_KEY) {
			c.backToMenu()
		}
	}
}

// oponent menu logic
func (c *C4) oponentMenuLogic() {
	if screenOponent { // if not on the screen menu
		oponentSelect()
		// Do a timer, when is a even sec blink, then hide.
		if shouldBlink {
			xPos := (width / 2) - (2 * gridSize)
			if oponentHover {
				xPos += (4 * gridSize)
			}
			oponentBlink(xPos)
		}

		if continuePressed {
			screenOponent = false
			screenBoard = true
			isOponentAI = oponentHover
		}
	}
}

// create and render board logic
func (c *C4) boardMenuLogic() {
	if !boardMade {
		// board
		c.board = make([][]int32, ROWS)
		for row := range c.board {
			c.board[row] = make([]int32, COLLUMNS)
		}

		// Players
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

		boardMade = true
	}

	// render board and every cell, including floating one
	if screenBoard {
		board()

		// render diff stuff acording to oponent
		if !isOponentAI {
			c.pvp()
		} else if isOponentAI {
			pvai()
		}

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

// * General logic funcs

// Checks if integer is even, returns true
func isEven(l int) bool {
	if l < 2 {
		return l%2 == 0
	}
	return isEven(l - 2)
}

// Checks if a []int contains a int, returns true
func contains(i []int, intg int) bool {
	for _, v := range i {
		if v == intg {
			return true
		}
	}

	return false
}

func insert(a []int32, index int, value int32) []int32 {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

// * Gameplay general funcs

// Gets the current collumn
func currentCol(i []int, intg int) int {
	for k, v := range i {
		if int(rl.GetMouseX())-v < intg {
			return k
		}
	}
	return -1
}

// Checks if the collumn is full
func (c *C4) collFull() bool {
	collCurrent = currentCol(collsPos, int(HEIGHT)*33/256)
	if collHeight[collCurrent] < COLLUMNS-1 {
		return false
	} else if collHeight[collCurrent] >= COLLUMNS-1 {
		return true
	}
	return true
}

// Clears the board from every cell and state
func (c *C4) resetBoard() {
	for row := range c.board {
		for col := range c.board[row] {
			c.board[row][col] = EMPTY
		}
	}
	movesMade = 0
	collHeight = []int{0, 0, 0, 0, 0, 0, 0}
}

// Return to the main Menu
func (c *C4) backToMenu() {
	screenMenu = true
	screenConf = false
	screenOponent = false
	screenBoard = false
	gameWinner = 0
}

// Handles debug menu
func (c *C4) debugToggle() {
	if debugMenu {
		debugMenu = false
	} else if !debugMenu {
		debugMenu = true
	}
}

// * Gameplay Logic

// handles turn switching
func (c *C4) switchTurn() {
	if c.turn == c.P1.ID {
		c.turn = c.P2.ID
	} else if c.turn == c.P2.ID {
		c.turn = c.P1.ID
	}
}

func (c *C4) makeMove() {
	if cursorOverBoard && !c.collFull() {
		movesMade += 1
		collHeight[collCurrent] += 1
		c.board[ROWS-collHeight[collCurrent]][collCurrent] = c.turn
		c.detect4()
	}
}

// Detects if 4 have been connected on the X and Y axis
func (c *C4) detectXY4(l []int32, s int32) bool {
	y = 0
	for _, x := range l {
		if x == s {
			y += 1
			if y == 4 {
				return true
			}
		} else if x != s {
			y = 0
		}
	}
	return false
}

func (c *C4) createDiag2DArrays(s int32) (diagTopRight, diagTopLeft [][]int32) {
	tempSlice1 := []int32{}
	tempSlice2 := []int32{}


	diag2D_TopToRight := [][]int32{}
	diag2D_TopToLeft := [][]int32{}
	for p := 0; p <= 12; p++ {
		for iRow, _ := range c.board {
			for iCol, cellState := range c.board[iRow] {
				if cellState != 0 {
					if oldP != p {
						diag2D_TopToRight, diag2D_TopToLeft = append(diag2D_TopToRight, tempSlice1), append(diag2D_TopToLeft, tempSlice2)
						tempSlice1, tempSlice2 = nil, nil
					}
					if iRow+iCol == p {
						coord = c.board[iRow][iCol]
						tempSlice1 = append(tempSlice1, coord)
					}
					if iRow-iCol+6 == p {
						coord = c.board[iRow][iCol]
						tempSlice2 = append(tempSlice2, coord)
					}
	
					oldP = p
				}
			}
		}
	}
	for w := range diag2D_TopToLeft {
		fmt.Println(diag2D_TopToLeft[w])
	}
	return diag2D_TopToRight, diag2D_TopToLeft
}

// Gives a 2D array of Y based board
func (c *C4) multiDimSliceY() [][]int32 {
	for y := 0; y < 6; y++ {
		tempSlice := []int32{}
		for x := 0; x < 6; x++ {
			tempSlice = append(tempSlice, c.board[x][y])
		}
		twoDimY = append(twoDimY, tempSlice)
	}
	return twoDimY
}

func (c *C4) detect4() {

	// Detects on Y
	twoDimY := c.multiDimSliceY()
	for col := range twoDimY {
		if c.detectXY4(twoDimY[col], c.P1.ID) {
			gameWinner = c.P1.ID
		} else if c.detectXY4(twoDimY[col], c.P2.ID) {
			gameWinner = c.P2.ID

		}
	}
	// Detects on X
	for row := range c.board {
		if c.detectXY4(c.board[row], c.P1.ID) {
			gameWinner = c.P1.ID
		} else if c.detectXY4(c.board[row], c.P2.ID) {
			gameWinner = c.P2.ID
		}
	}
	// Detects Diag
	diagTtR, diagTtL := c.createDiag2DArrays(c.P1.ID)
	for row := range diagTtR {
		if c.detectXY4(diagTtR[row], c.P1.ID) {
			gameWinner = c.P1.ID
		} else if c.detectXY4(diagTtR[row], c.P2.ID) {
			gameWinner = c.P2.ID
		}
	}

	for row := range diagTtL {
		if c.detectXY4(diagTtL[row], c.P1.ID) {
			gameWinner = c.P1.ID
		} else if c.detectXY4(diagTtL[row], c.P2.ID) {
			gameWinner = c.P2.ID
		}
	}

	c.switchTurn()
}
