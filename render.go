package main

import (
	"fmt"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// * render is a funcion that starts every single menu and texture
func (c *C4) render() {
	logger.Debug().Println("render() called.")

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)
	rl.DrawRectangleLines(offsetX, offsetY, width-(2*offsetX), height-(2*offsetY), rl.LightGray)
}

// * draws the main menu
func mainMenuRender() {
	rl.DrawText("Connect 4!", width/2-height*162/768, height/12, height*5/64, rl.LightGray)
	rl.DrawText("Start!", width/2-height*35/384, height-height/4+height/32, height/16, rl.LightGray)
	rl.DrawText("Config", width/2-height*5/64, height-height/4+2*height*5/96, height*5/96, rl.LightGray)
	rl.DrawText("Quit", width/2-height*15/256, height-height/4+(7/2)*height*5/96, height*5/96, rl.LightGray)

	rl.DrawTexture(txrLogo, (width-278)/2, height/14*3, rl.White)
}

func confMenuRender() {
	rl.DrawText("Config", width/2-height*110/768, height/12, height*5/64, rl.LightGray)
	rl.DrawText("Music", width/2-height*35/384, height*25/64, height*5/96, rl.LightGray)
	rl.DrawText("Sounds", width/2-height*35/384, height*175/384, height*5/96, rl.LightGray)
	rl.DrawText("Fullscreen", width/2-height*35/384, height*25/48, height*5/96, rl.LightGray)
	rl.DrawText("Return", width/2-height*75/768, height-height/4+(7/2)*height*5/96, height*5/96, rl.LightGray)
}

func blinkRender(xPos, yPos, xMag, yMag int32) {
	xMag, yMag = xMag-xPos, yMag-yPos
	rl.DrawRectangleLines(
		xPos, yPos, xMag, yMag, rl.White)
}

// * Oponent selection screen
func oponentMenuRender() {
	rl.DrawText("Do you wish to play with someone or against a bot?", width*10/100, height/12, height*5/96, rl.LightGray)

	rl.DrawText("PvP", (width/2)-(2*gridSize), height/12*(40/35*8), height*7/128, rl.LightGray)
	rl.DrawText("AI", (width/2)+(2*gridSize), height/12*(40/35*8), height*7/128, rl.LightGray)
}
func oponentBlink(xPos int32) {
	rl.DrawRectangleLines(
		xPos-(height/100),
		blinkOponentVer.yPos,
		blinkOponentVer.xMag,
		blinkOponentVer.yMag, rl.LightGray)
}

func (c *C4) pvpRender(startText, p1WinText, p2WinText string) {
	if movesMade != 42 {
		switch gameWinner {
		case 0:
			rl.DrawText(startText, width*10/100, height/12, height*7/192, rl.LightGray)
		case c.P1.ID:
			rl.DrawText(p1WinText, width*10/100, height/12, height*7/192, rl.LightGray)
		case c.P2.ID:
			rl.DrawText(p2WinText, width*10/100, height/12, height*7/192, rl.LightGray)
		}
	}
	if movesMade == 42 || gameWinner != 0 {
		rl.DrawText("Press enter to reset board", width*10/100, height/6, height*7/192, rl.LightGray)
	}
}

func boardRender() {
	// rectangle for the board
	rl.DrawRectangle(
		boardVer.xPos,
		boardVer.yPos,
		boardVer.xMag,
		boardVer.yMag, rl.DarkBlue)
}
func gridRender(col, row int, color rl.Color) {
	rl.DrawCircle(
		int32(col*int(gridSize)+int(boardVer.xPos)+int(1.3*gridRadius)),
		int32(row*int(gridSize)+int(boardVer.yPos)+int(1.3*gridRadius)),
		gridRadius, color)
}

func (c *C4) floatingCellRender() {
	// checks if the mouse its over the board.
	if cursorOverBoard {
		rl.DrawCircle(rl.GetMouseX(), boardVer.yPos-gridSize/2, gridRadius, c.gamers[c.turn-1].CellColour)
	}
}

func (c *C4) boardDebugRender() {
	debPosY = 0

	fmt.Println(rl.GetMouseX(), rl.GetMouseY())

	rl.DrawLine(width/2, 0, width/2, height, rl.Beige)
	rl.DrawLine(0, height/2, width, height/2, rl.Beige)
	if screenBoard {
		rl.DrawText("Current Collumn: "+strconv.FormatInt(int64(collCurrent), 10), offsetX, offsetY, height*5/192, rl.Beige)
		rl.DrawText("Moves Made: "+strconv.FormatInt(int64(movesMade), 10), offsetX, 2*offsetY, height*5/192, rl.Beige)

		for row := range c.board.grid {
			a := fmt.Sprint(c.board.grid[row])
			rl.DrawText(a, width-2*offsetX-height*25/150, debPosY+offsetY+height*5/384, height*5/192, rl.Beige)
			debPosY += height / 128 * 5
		}
	}
}
