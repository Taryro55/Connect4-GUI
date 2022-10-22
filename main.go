package main

import (
	// "fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var logger = logging()

// * reciver func that is constantly looping listening for events
func (c *C4) update() {
	logger.Debug().Println("update() called.")

	executing = !rl.WindowShouldClose()

	// Music
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else if !musicPaused {
		rl.ResumeMusicStream(music)
	}

	// Selection of oponent
	if isEven(runningTimeInt) {
		shouldBlink = true
	} else if !isEven(runningTimeInt) {
		shouldBlink = false
	}

	// Slice containing pixels of the 7 coll
	for i, colls := 0, int(HEIGHT*355/768); i < 7; i, colls = i+1, colls+int(HEIGHT*95/768) {
		if !contains(collsPos, colls) {
			collsPos = append(collsPos, colls)
		}
	}

	// Checks to end game
	if movesMade == 42 {
		if rl.IsKeyDown(rl.KeyEnter) {
			c.endGame()
		}
	}

	runningTime = rl.GetTime()
	runningTimeInt = int(runningTime)
}

// * starts basic stuff
func init() {
	logger.Debug().Println("init() executed.")

	rl.InitWindow(width, HEIGHT, WINDOW_TITLE)
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)

	collHeight = []int{0, 0, 0, 0, 0, 0, 0}

	// textures
	txrWhite = rl.LoadTexture(TXR_PATH + "/red.jpg")
	txrBlack = rl.LoadTexture(TXR_PATH + "/oreo.png")
	txrBoard = rl.LoadTexture(TXR_PATH + "/board.png")
	txrBackground = rl.LoadTexture(TXR_PATH + "/background.jpg")
	txrLogo = rl.LoadTexture(TXR_PATH + "/logo.png")

	// Loads music
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("./music/menu.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)
}

// * unloads everything that was loaded on init and quits
func quit() {
	logger.Debug().Println("quit() called.")

	rl.UnloadTexture(txrWhite)
	rl.UnloadTexture(txrBlack)
	rl.UnloadTexture(txrBoard)
	rl.UnloadTexture(txrBackground)
	rl.UnloadTexture(txrLogo)

	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()

	rl.CloseWindow()
}

// * main calls every single other funcion
func main() {
	connect4 := C4{}
	logger.Debug().Println("main() called.")
	for executing {
		// fmt.Println(mouseButtonPressed, gameOngoing)

		logger.Debug().Println("in executing loop.")
		connect4.input()
		connect4.update()
		connect4.render()
		connect4.logic()
	}
	quit()
}
