package main

import (
	"github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH = int32(800)
	HEIGHT = int32(450)
	WINDOW_TITLE = "Connect 4"
	ROWS = 6
	COLLUMNS = 7
	GRID_SIZE = 400
)


var (
	twhite			rl.Texture2D
	tblack			rl.Texture2D
	mousePos		rl.Vector2
)

var (
	gameOver           bool = false
	gameDraw           bool = false
	gameOngoing		   bool = false
	mouseButtonPressed bool = false
	/*
	playerMove         bool = false // If bool false, its whites turn, if true, its blacks turn
	positions          byte = 0
	*/
)


func main() {
	// init
	rl.InitWindow(WIDTH , HEIGHT, WINDOW_TITLE)
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)

	// textures
	//twhite = rl.LoadTexture()
	//tblack = rl.LoadTexture()

	for !rl.WindowShouldClose() {

		// Fullscreen When Alt+Enter
		if (rl.IsKeyPressed(257) && (rl.IsKeyDown(342) || rl.IsKeyDown(346))) {
			display := rl.GetCurrentMonitor()

			if rl.IsWindowFullscreen() {
				rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
			} else {
				rl.SetWindowSize(int(WIDTH), int(HEIGHT))
			}
			rl.ToggleFullscreen()
		}

		
		if (!gameOngoing) { // menu
			rl.BeginDrawing()
			rl.ClearBackground(rl.Black)
			rl.DrawRectangleLines(WIDTH*2/100, HEIGHT*2/100, WIDTH-WIDTH*4/100, HEIGHT-HEIGHT*4/100, rl.LightGray)

			rl.DrawText("Welcome to Connect 4!\nPress enter or click!", WIDTH*32/100, HEIGHT/12, 28, rl.LightGray)
		}
		
		// process events
		if (rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton)) {
			mouseButtonPressed = true
		}

		if (!gameOver || !gameDraw) {
			// game logic
			mousePos = rl.GetMousePosition()
			if mouseButtonPressed {
				rl.ClearBackground(rl.White)
				rl.DrawText("Welcome to Connect 4!\nPress enter or click!", WIDTH*32/100, HEIGHT/12, 28, rl.Black)
			}
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}