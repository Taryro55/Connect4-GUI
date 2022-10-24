package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// * input manages every single input avaliable.
func (c *C4) input() {
	logger.Debug().Println("input() called.")

	// Music toggle
	if rl.IsKeyPressed(rl.KeyM) {
		musicPaused = !musicPaused
	}

	// Oponent select
	if screenOponent {
		if rl.IsKeyPressed(rl.KeyRight) && !oponentHover {
			oponentHover = true
		} else if rl.IsKeyPressed(rl.KeyLeft) && oponentHover {
			oponentHover = false
		}
	}

	// Back to menu
	if rl.IsKeyPressed(rl.KeyBackspace) {
		c.resetBoard()
		c.backToMenu()
	}

	// Fullscreen When Alt+Enter
	if rl.IsKeyPressed(257) && (rl.IsKeyDown(342) || rl.IsKeyDown(346)) {
		display := rl.GetCurrentMonitor()

		if rl.IsWindowFullscreen() {
			rl.SetWindowSize(rl.GetMonitorWidth(display), rl.GetMonitorHeight(display))
		} else {
			rl.SetWindowSize(int(width), int(HEIGHT))
		}
		rl.ToggleFullscreen()
	}

	// Process events
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) || rl.IsKeyPressed(rl.KeyEnter) {
		continuePressed = true
	} else if !rl.IsMouseButtonPressed(rl.MouseLeftButton) || !rl.IsMouseButtonPressed(rl.MouseRightButton) || !rl.IsKeyPressed(rl.KeyEnter) {
		continuePressed = false
	}

	// Switch turns & gameplay
	if screenBoard && continuePressed {
		c.makeMove()
	}

	// Debug
	if rl.IsKeyPressed(rl.KeyF3) {
		c.debugToggle()
	}

}
