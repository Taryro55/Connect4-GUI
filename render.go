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
	rl.DrawRectangleLines(offsetX, offsetY, width-(2*offsetX), HEIGHT-(2*offsetY), rl.LightGray)
}

// * draws the main menu
func mainmenu() {
	rl.DrawText("Connect 4!", width/2-HEIGHT*162/768, HEIGHT/12, HEIGHT*5/64, rl.LightGray)
	rl.DrawText("Start!", width/2-HEIGHT*35/384, HEIGHT-HEIGHT/4+HEIGHT/32, HEIGHT/16, rl.LightGray)
	rl.DrawText("Config", width/2-HEIGHT*5/64, HEIGHT-HEIGHT/4+2*HEIGHT*5/96, HEIGHT*5/96, rl.LightGray)
	rl.DrawText("Quit", width/2-HEIGHT*15/256, HEIGHT-HEIGHT/4+(7/2)*HEIGHT*5/96, HEIGHT*5/96, rl.LightGray)


	rl.DrawTexture(txrLogo, (width-278)/2, HEIGHT/14*3, rl.White)
}

func confmenu() {
	rl.DrawText("Config", width/2-HEIGHT*110/768, HEIGHT/12, HEIGHT*5/64, rl.LightGray)
	rl.DrawText("Music", width/2-HEIGHT*35/384, HEIGHT-HEIGHT/4+HEIGHT/32, HEIGHT*5/96, rl.LightGray)
	rl.DrawText("Sound Effects", width/2-HEIGHT*25/128, HEIGHT-HEIGHT/4+2*HEIGHT*5/96, HEIGHT*5/96, rl.LightGray)
	rl.DrawText("Fullscreen", width/2-HEIGHT*25/128, HEIGHT-HEIGHT/4+(7/2)*HEIGHT*5/96, HEIGHT*5/96, rl.LightGray)

}

func confBlink() {
	// rl.DrawRectangleLines(

	// )
}

// * Oponent selection screen
func oponentSelect() {
	// TODO Change Music when here
	rl.DrawText("Do you wish to play with someone or against a bot?", width*10/100, HEIGHT/12, 28, rl.LightGray)

	rl.DrawText("PvP", (width/2)-(2*gridSize), HEIGHT/12*(40/35*8), 28, rl.LightGray)
	rl.DrawText("AI", (width/2)+(2*gridSize), HEIGHT/12*(40/35*8), 28, rl.LightGray)
}
func oponentBlink(xPos int32) {
	rl.DrawRectangleLines(
		xPos-(HEIGHT/100),
		blinkOponentVer.yPos,
		blinkOponentVer.xMag,
		blinkOponentVer.yMag, rl.LightGray)
}

func (c *C4) pvp() {
	if movesMade != 42 && gameWinner == 0 {
		rl.DrawText("Game on, let the better human win!", width*10/100, HEIGHT/12, 28, rl.LightGray)
	} else if movesMade == 42 {
		rl.DrawText("Press enter to reset board", width*10/100, HEIGHT/12, 28, rl.LightGray)
	}
	switch gameWinner {
	case c.P1.ID:
		rl.DrawText("Red Won", width*10/100, HEIGHT/12, 28, rl.LightGray)
	case c.P2.ID:
		rl.DrawText("Yellow Won", width*10/100, HEIGHT/12, 28, rl.LightGray)

	}
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
}
func grid(col, row int, color rl.Color) {
	rl.DrawCircle(
		int32(col*int(gridSize)+int(boardVer.xPos)+int(1.3*gridRadius)),
		int32(row*int(gridSize)+int(boardVer.yPos)+int(1.3*gridRadius)),
		gridRadius, color)
}

func (c *C4) floatingCell() {
	// checks if the mouse its over the board.
	if cursorOverBoard {
		rl.DrawCircle(rl.GetMouseX(), boardVer.yPos-gridSize/2, gridRadius, c.gamers[c.turn-1].CellColour)
	}
}

func (c *C4) boardDebug() {
	debPosY = 0

	fmt.Println(rl.GetMouseX())

	rl.DrawLine(width/2, 0, width/2, HEIGHT, rl.Beige)
	rl.DrawLine(0, HEIGHT/2, width, HEIGHT/2, rl.Beige)
	if screenBoard {
		rl.DrawText("Current Collumn: "+strconv.FormatInt(int64(collCurrent), 10), offsetX, offsetY, 20, rl.Beige)
		rl.DrawText("Moves Made: "+strconv.FormatInt(int64(movesMade), 10), offsetX, 2*offsetY, 20, rl.Beige)

		for row := range c.board {
			a := fmt.Sprint(c.board[row])
			rl.DrawText(a, width-2*offsetX-HEIGHT*25/150, debPosY+offsetY+HEIGHT*5/384, 20, rl.Beige)
			debPosY += HEIGHT / 128 * 5
		}
	}
}