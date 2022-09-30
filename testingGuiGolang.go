package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	WIDTH        = int32(800)
	HEIGHT       = int32(450)
	WINDOW_TITLE = "Connect 4"
	ROWS         = 6
	COLLUMNS     = 7
	GRID_SIZE    = 400
)

var (
	twhite      rl.Texture2D
	tblack      rl.Texture2D
	tboard      rl.Texture2D
	tbackground rl.Texture2D
	mousePos    rl.Vector2
)

var (
	executing          bool = false
	gameOver           bool = false
	gameDraw           bool = false
	gameOngoing        bool = false
	mouseButtonPressed bool = false
	playerMove         bool = false // If bool false, its whites turn, if true, its blacks turn
	//positions          byte = 0
	musicPaused 	   bool
	music       rl.Music
)

func drawScene() {
	// rl.DrawTexture(tbackground, 100, 0, rl.White)
	// rl.DrawTexture(tboard, 0, 0, rl.Blue)
	// rl.DrawTexture(twhite, 0, 100, rl.White)
	rl.DrawTexture(tblack, 0, 200, rl.White)
}

func input() {
	//
	if rl.IsKeyPressed(rl.KeyM) {
		musicPaused = !musicPaused
	}

	// Fullscreen When Alt+Enter
	if rl.IsKeyPressed(257) && (rl.IsKeyDown(342) || rl.IsKeyDown(346)) {
		display := rl.GetCurrentMonitor()

		if rl.IsWindowFullscreen() {
			rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
		} else {
			rl.SetWindowSize(int(WIDTH), int(HEIGHT))
		}
		rl.ToggleFullscreen()
	}

	// process events
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) {
		mouseButtonPressed = true
		gameOngoing = true
	}

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("./music/menu.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)
}

func update() {
	executing = !rl.WindowShouldClose()

	// Music
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else if !musicPaused {
		rl.ResumeMusicStream(music)
	}
}

func render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)

	rl.DrawRectangleLines(WIDTH*2/100, HEIGHT*2/100, WIDTH-WIDTH*4/100, HEIGHT-HEIGHT*4/100, rl.LightGray)
	renderMenu()

	drawScene()
	rl.EndDrawing()
}

func renderMenu() {
	if !gameOngoing {
		rl.DrawText("Welcome to Connect 4!\nPress enter or click!", WIDTH*32/100, HEIGHT/12, 28, rl.LightGray)
		if mouseButtonPressed {
			gameOngoing = true
		}
	}
}

func renderBoard() {
	// Render a board thats more than a texture
}

func gameLogic() {
	if gameOngoing && (!gameOver || !gameDraw) {
		fmt.Println("HELLO")
		// logic
	}
}

func initStuff() {
	rl.InitWindow(WIDTH, HEIGHT, WINDOW_TITLE)
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)

	// textures
	twhite = rl.LoadTexture("./textures/red.jpg")
	tblack = rl.LoadTexture("./textures/oreo.png")
	tboard = rl.LoadTexture("./textures/board.png")
	tbackground = rl.LoadTexture("./textures/background.jpg")
}

func quit() {
	rl.UnloadTexture(twhite)
	rl.UnloadTexture(tblack)
	rl.UnloadTexture(tboard)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.UnloadTexture(tbackground)

	rl.CloseWindow()
}

func main() {
	initStuff()
	for !rl.WindowShouldClose() { // Using executing makes the windows insta close. WHY
		input()
		update()
		render()
		gameLogic()
	}
	quit()
}
