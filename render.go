package main

import (
	// "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// * render is a funcion that starts every single menu and texture
func (c *C4) render() {
	logger.Debug().Println("render() called.")

	rl.BeginDrawing()
	defer rl.EndDrawing()

	rl.ClearBackground(rl.Black)
	rl.DrawRectangleLines(width*2/100, HEIGHT*3/100, width-width*4/100, HEIGHT-HEIGHT*6/100, rl.LightGray)

	menu()
	txr()
}

// * draws the main menu
func menu() {
	logger.Debug().Println("menu() called.")
	
	if !gameOngoing {
		// Perfectly centered rectangle
		// rl.DrawRectangle((width - (gridSize*COLLUMNS))/2, gridSize, gridSize*COLLUMNS, gridSize*ROWS, rl.Gold)
		rl.DrawText("Welcome to Connect 4!", width*32/100, HEIGHT/12, 28, rl.LightGray)
		rl.DrawText("Press enter or click to start!", width*25/100, HEIGHT-HEIGHT/4, 28, rl.LightGray)
		rl.DrawTexture(txrLogo, (width-278)/2, HEIGHT/14*3, rl.White)
		if mouseButtonPressed {
			gameOngoing = true
		}
	} else if gameOngoing {
		board()
	}
}

// * draws the board
func board() {
	// TODO Change Music when here
	if gameOngoing {
			
		rl.DrawTexture(txrBlack, 0, 200, rl.White)
		rl.DrawText("BOARD!", width*32/100, HEIGHT/12, 28, rl.LightGray)
		if mouseButtonPressed {
			gameOngoing = false
		}
	}

	// TODO Board Logic
}

// * draws textures
func txr() {
	logger.Debug().Println("drawTxr() called.")

	// rl.DrawTexture(txrWhite, 0, 100, rl.White)
	// rl.DrawTexture(txrBackground, 100, 0, rl.White)
	// rl.DrawTexture(txrBoard, 0, 0, rl.Blue)
}
