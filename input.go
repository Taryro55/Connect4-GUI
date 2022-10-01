package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func input() {
	logger.Debug().Println("input() called.")

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

}
