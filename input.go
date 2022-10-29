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

	// Back to menu
	if rl.IsKeyPressed(rl.KeyBackspace) {
		c.resetBoard()
		c.backToMenu()
	}

	// Fullscreen When Alt+Enter
	if rl.IsKeyPressed(257) && (rl.IsKeyDown(342) || rl.IsKeyDown(346)) {
		fullScreen()
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

func mainMenuInput() {
	if rl.IsKeyPressed(rl.KeyDown) && mainMenuHover <= 1 {
		mainMenuHover++
	} else if rl.IsKeyPressed(rl.KeyUp) && mainMenuHover >= 0 {
		mainMenuHover--
	}
}

func configMenuInput() {
	if rl.IsKeyPressed(rl.KeyDown) && configMenuHover <= 2 {
		configMenuHover++
	} else if rl.IsKeyPressed(rl.KeyUp) && configMenuHover >= 0 {
		configMenuHover--
	}
}

func oponentMenuInput() {
	if rl.IsKeyPressed(rl.KeyRight) && !oponentMenuHover {
		oponentMenuHover = true
	} else if rl.IsKeyPressed(rl.KeyLeft) && oponentMenuHover {
		oponentMenuHover = false
	}
}