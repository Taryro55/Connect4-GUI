package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var logger = logging()

// * reciver func that is constantly looping listening for events
func (c *C4) update() {
	logger.Debug().Println("update() called.")

	executing = !rl.WindowShouldClose()

	// Music managment
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else if !musicPaused {
		rl.ResumeMusicStream(music)
	}

	// Sfx managment
	if sfxPaused {
		rl.PauseSound(sfxPop)
	} else if !sfxPaused {
		rl.ResumeSound(sfxPop)
	}

	// Selection of oponent
	if isEven(int(runningMilisecs * 2)) {
		shouldBlink = true
	} else if !isEven(int(runningMilisecs * 2)) {
		shouldBlink = false
	}

	// Slice containing pixels of the 7 coll
	for i, colls := 0, int(height*355/768); i < 7; i, colls = i+1, colls+int(height*95/768) {
		if !contains(collsPos, int32(colls)) {
			collsPos = append(collsPos, int32(colls))
		}
	}

	// Checks if the mouse if above the rendered board
	if boardVer.xPos+boardXtra < rl.GetMouseX() && rl.GetMouseX() < boardVer.xPos+boardVer.xMag-boardXtra {
		cursorOverBoard = true
	} else {
		cursorOverBoard = false
	}
}

// * starts basic stuff
func init() {
	logger.Debug().Println("init() executed.")

	rl.InitWindow(width, height, WINDOW_TITLE)
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)

	// Loads stuff
	rl.InitAudioDevice()

	txrLogo = rl.LoadTexture(TXR_PATH + "/logo.png")
	music = rl.LoadMusicStream(MUSIC_PATH + "/menu.mp3")
	sfxPop = rl.LoadSound(SFX_PATH + "/pop.mp3")

	musicPaused = false
	rl.PlayMusicStream(music)
}

// * unloads everything that was loaded on init and quits
func quit() {
	logger.Debug().Println("quit() called.")

	rl.UnloadTexture(txrLogo)
	rl.UnloadSound(sfxPop)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()

	rl.CloseWindow()
}

// * main calls every single other funcion
func main() {
	connect4 := C4{}
	logger.Debug().Println("main() called.")
	for executing {
		logger.Debug().Println("in executing loop.")
		connect4.input()
		connect4.update()
		connect4.render()
		connect4.logic()
		mainLoops += 1

		// I just wanna say that dark chocolate is underrated - Lem 10/25/2022
		// Overdose on dark chocolate - Lem 10/25/2022
		// A hot woman wouldn't stop talking to me - Lem 1/11/2022
	}
	quit()
}
