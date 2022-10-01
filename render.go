package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// * render is a funcion that starts every single menu and texture
// TODO Check if the border works on every single screen
func render() {
	logger.Debug().Println("render() called.")

	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangleLines(WIDTH*2/100, HEIGHT*2/100, WIDTH-WIDTH*4/100, HEIGHT-HEIGHT*4/100, rl.LightGray)
	renderMenu()

	drawTxr()
	rl.EndDrawing()
}

func renderMenu() {
	logger.Debug().Println("renderMenu() called.")

	if !gameOngoing {
		rl.DrawText("Welcome to Connect 4!\nPress enter or click!", WIDTH*32/100, HEIGHT/12, 28, rl.LightGray)
		if mouseButtonPressed {
			gameOngoing = true
		}
	}
}

// func renderBoard() {
// TODO Render a board thats more than a texture
// }

func drawTxr() {
	logger.Debug().Println("drawTxr() called.")
	// rl.DrawTexture(tbackground, 100, 0, rl.White)
	// rl.DrawTexture(tboard, 0, 0, rl.Blue)
	// rl.DrawTexture(twhite, 0, 100, rl.White)
	rl.DrawTexture(tblack, 0, 200, rl.White)
}
