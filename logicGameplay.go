package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// * Gameplay general funcs

func fullScreen() {
	display := rl.GetCurrentMonitor()

	if rl.IsWindowFullscreen() {
		height = int32(rl.GetMonitorHeight(display))
		rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
	} else {
		height = HEIGHT
		rl.SetWindowSize(int(width), int(height))
	}
	rl.ToggleFullscreen()
}

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
	collCurrent = currentCol(collsPos, int(height)*33/256)
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

// gets 2ds diagonal arrays of the board
func (c *C4) diag2DArrays() (diagTopRight, diagTopLeft [][]int32) {
	tempSlice1 := []int32{}
	tempSlice2 := []int32{}

	diag2D_TopToRight := [][]int32{}
	diag2D_TopToLeft := [][]int32{}
	for p := 0; p <= 12; p++ {
		for iRow := range c.board {
			for iCol := range c.board[iRow] {
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
	diagTtR, diagTtL := c.diag2DArrays()
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
