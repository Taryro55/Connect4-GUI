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
	if !oponentSelected {
		if rl.IsKeyPressed(rl.KeyRight) && !isOponentAI {
			isOponentAI = true
		} else if rl.IsKeyPressed(rl.KeyLeft) && isOponentAI {
			isOponentAI = false
		}
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
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) || rl.IsMouseButtonPressed(rl.MouseRightButton) {
		mouseButtonPressed = true
	} else if !rl.IsMouseButtonPressed(rl.MouseLeftButton) || !rl.IsMouseButtonPressed(rl.MouseRightButton){
		mouseButtonPressed = false
	}

	// Switch turns
	if gameOngoing && (!gameOver || !gameDraw) && oponentSelected && boardRendered {
		if mouseButtonPressed {
			if c.turn == c.P1.ID {
				c.turn = c.P2.ID
			} else if c.turn == c.P2.ID {
				c.turn = c.P1.ID
			}
		}
	}

	// Debug
	if rl.IsKeyPressed(rl.KeyF3) {
		if debugMenu {
			debugMenu = false
		} else if !debugMenu {
			debugMenu = true
		}
	}

}
