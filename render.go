package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// * render is a funcion that starts every single menu and texture
func (c *C4) render() {
	logger.Debug().Println("render() called.")

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)
	rl.DrawRectangleLines(offsetX, offsetY, width-(2*offsetX), HEIGHT-(2*offsetY), rl.LightGray)
}

// * draws the main menu
func mainmenu() {
	rl.DrawText("Welcome to Connect 4!", width*32/100, HEIGHT/12, 28, rl.LightGray)
	rl.DrawText("Press enter or click to start!", width*25/100, HEIGHT-HEIGHT/4, 28, rl.LightGray)
	rl.DrawTexture(txrLogo, (width-278)/2, HEIGHT/14*3, rl.White)
}

// * Oponent selection screen
func oponentSelect() {
	// TODO Change Music when here
	rl.DrawText("Do you wish to play with someone or against a bot?", width*10/100, HEIGHT/12, 28, rl.LightGray)

	rl.DrawText("PvP", (width/2)-(2*gridSize), HEIGHT/12*(40/35*8), 28, rl.LightGray)
	rl.DrawText("AI", (width/2)+(2*gridSize), HEIGHT/12*(40/35*8), 28, rl.LightGray)
}
func blink(xPos int32) {
	rl.DrawRectangleLines(
		xPos-(HEIGHT/100),
		blinkVer.yPos,
		blinkVer.xMag,
		blinkVer.yMag, rl.LightGray)
}

func pvp() {
	rl.DrawText("Game on, let the better human win!", width*10/100, HEIGHT/12, 28, rl.LightGray)
}
func pvai() {
	rl.DrawText("Game on, I wish you the best...", width*10/100, HEIGHT/12, 28, rl.LightGray)
}

func board() {
	// rectangle for the board
	rl.DrawRectangle(
		boardVer.xPos,
		boardVer.yPos,
		boardVer.xMag,
		boardVer.yMag, rl.DarkBlue)
	if !boardRendered {
		boardRendered = true
	}
}
func grid(col, row int, color rl.Color) {
	rl.DrawCircle(
		int32(col*int(gridSize)+int(boardVer.xPos)+int(1.3*gridRadius)),
		int32(row*int(gridSize)+int(boardVer.yPos)+int(1.3*gridRadius)),
		gridRadius, color)
}

func (c *C4) floatingCell() {
	// checks if the mouse its over the board.
	if (boardVer.xPos+boardXtra < rl.GetMouseX() && rl.GetMouseX() < boardVer.xPos+boardVer.xMag-boardXtra) {
		rl.DrawCircle(rl.GetMouseX(), boardVer.yPos-gridSize/2, gridRadius, c.gamers[c.turn-1].CellColour)
	}
}

func centerHelper() {
	rl.DrawLine(width/2, 0, width/2, HEIGHT, rl.Beige)
	rl.DrawLine(0, HEIGHT/2, width, HEIGHT/2, rl.Beige)
}

func (c *C4) boardState() {
	debPosY = 0
	if gameOngoing && (!gameOver || !gameDraw) && oponentSelected && boardRendered {
		for row := range c.board {
			a := fmt.Sprint(c.board[row])
			rl.DrawText(a, width-2*offsetX-HEIGHT*25/150, debPosY+offsetY+HEIGHT*5/384, 20, rl.Beige)
			debPosY += HEIGHT/128*5
		}
	}
}

// General Printing checks
func (c *C4) printDebug() {

}