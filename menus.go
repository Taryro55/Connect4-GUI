package main

import (
	"os"

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
		c.boardDebugRender()
	}

}

// * Menus Logic

// main menu logic
func (c *C4) mainMenuLogic() {
	if screenMenu {
		mainMenuRender()
		mainMenuInput()
		if shouldBlink {
			switch mainMenuHover {
			case 0:
				blinkRender(605, 595, 755, 645)
			case 1:
				blinkRender(615, 655, 750, 698)
			case 2:
				blinkRender(630, 696, 715, 735)
			}

		}
		if continuePressed {
			screenMenu = false
			switch mainMenuHover {
			case 0:
				screenOponent = true
			case 1:
				screenConf = true
			case 2:
				os.Exit(0)
			}
		}
	}
}

func (c *C4) configMenuLogic() {
	if screenConf {
		confMenuRender()
		configMenuInput()
		if musicPaused {
			rl.DrawCircle(580, 320, 20, rl.Red)
		}
		if soundPaused {
			rl.DrawCircle(580, 370, 20, rl.Red)
		}
		if shouldBlink {
			switch configMenuHover {
			case 0:
				blinkRender(605, 300, 724, 340)
			case 1:
				blinkRender(605, 350, 760, 390)
			case 2:
				blinkRender(605, 400, 830, 440)
			case 3:
				blinkRender(600, 690, 750, 740)
			}
		}
		if continuePressed {
			switch configMenuHover {
			case 0:
				musicPaused = !musicPaused
			case 1:
				soundPaused = !soundPaused
			case 2:
				fullScreen()
			case 3:
				screenConf = false
				screenMenu = true
			}
		}
	}
}

// oponent menu logic
func (c *C4) oponentMenuLogic() {
	if screenOponent { // if not on the screen menu
		oponentMenuRender()
		oponentMenuInput()
		// Do a timer, when is a even sec blink, then hide.
		if shouldBlink {
			xPos := (width / 2) - (2 * gridSize)
			if oponentMenuHover {
				xPos += (4 * gridSize)
			}
			oponentBlink(xPos) // TODO deprecated, replace with blinkRender()
		}

		if continuePressed {
			screenOponent = false
			screenBoard = true
			isOponentAI = oponentMenuHover
		}
	}
}

// create and render board logic
func (c *C4) boardMenuLogic() {
	if !boardMade {
		// board
		c.board = boardMake(ROWS, COLLUMNS)

		// Players
		c.P1 = Gamer{ID: 1, CellColour: rl.Red}
		c.P2 = Gamer{ID: 2, CellColour: rl.Yellow}
		c.gamers = []Gamer{c.P1, c.P2}
		c.turn = c.P1.ID

		boardMade = true
	}

	// render board and every cell, including floating one
	if screenBoard {
		boardRender()

		// render diff stuff acording to oponent
		if !isOponentAI {
			c.pvpRender("Game on! Humans compete!", "P1 Won!", "P2 Won!")
			c.floatingCellRender()
		} else if isOponentAI {
			c.pvpRender("Game on! Machine v/s Human!", "You won!", "Bot wins")
			if c.turn != c.P2.ID {
				c.floatingCellRender()
			}
		}

		for row := range c.board.grid {
			for col := range c.board.grid[row] {
				cellState := c.board.grid[row][col]
				if cellState == EMPTY {
					gridRender(col, row, rl.Black)
				} else {
					gridRender(col, row, c.gamers[cellState-1].CellColour)
				}
			}
		}
	}
}





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

// Clears the board from every cell and state
func (c *C4) resetBoard() {
	boardReset(c.board)
	c.turn = c.P1.ID
	gameWinner = 0
	movesMade = 0
	c.board.heights = []int32{0, 0, 0, 0, 0, 0, 0}
}

// Return to the main Menu
func (c *C4) backToMenu() {
	screenMenu = true
	screenConf = false
	screenOponent = false
	screenBoard = false
	c.resetBoard()
}

// Handles debug menu
func (c *C4) debugToggle() {
	if debugMenu {
		debugMenu = false
	} else if !debugMenu {
		debugMenu = true
	}
}