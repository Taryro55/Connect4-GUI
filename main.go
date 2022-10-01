package main

import (
	// "fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var logger = logging()

func update() {
	logger.Debug().Println("update() called.")

	executing = !rl.WindowShouldClose()

	// Music
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else if !musicPaused {
		rl.ResumeMusicStream(music)
	}
}

// * init is a func that... inits the basic stuff for raylib to work
func init() {
	logger.Debug().Println("init() executed.")	

	rl.InitWindow(WIDTH, HEIGHT, WINDOW_TITLE)
	rl.SetTargetFPS(60)
	rl.SetMouseScale(1.0, 1.0)

	// textures
	// TODO make a func thats on render thats called and loops over every single item on textures to load, same for unloading
	twhite = rl.LoadTexture("./textures/red.jpg")
	tblack = rl.LoadTexture("./textures/oreo.png")
	tboard = rl.LoadTexture("./textures/board.png")
	tbackground = rl.LoadTexture("./textures/background.jpg")
	

	// Loads music
	rl.InitAudioDevice()
	music = rl.LoadMusicStream("./music/menu.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)
}


// * quit unloads everything that was loaded on init and quits
func quit() {
	logger.Debug().Println("quit() called.")

	rl.UnloadTexture(twhite)
	rl.UnloadTexture(tblack)
	rl.UnloadTexture(tboard)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.UnloadTexture(tbackground)

	rl.CloseWindow()
}



// * main calls every single other funcion
func main() {
	logger.Debug().Println("main() called.")
	for executing { // ? Using executing makes the windows insta close. WHY		
		logger.Debug().Println("in executing loop.")

		input()
		update()
		render()
		gameLogic()
	}
	quit()
}
