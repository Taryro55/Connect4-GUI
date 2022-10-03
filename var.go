package main

import (
	"math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	txrWhite      rl.Texture2D
	txrBlack      rl.Texture2D
	txrBoard      rl.Texture2D
	txrBackground rl.Texture2D
	txrLogo		  rl.Texture2D
	// mousePos    rl.Vector2
	music       rl.Music
)

var (
	width = (HEIGHT/9) * 16
	gridSize = int32(math.Round((math.Hypot(float64(width), float64(HEIGHT))) * 0.1))
)

var (
	debug 			   bool = true
	executing          bool = true
	gameOver           bool = false
	gameDraw           bool = false
	gameOngoing        bool = false
	mouseButtonPressed bool = false
	// playerMove         bool = false // If bool false, its whites turn, if true, its blacks turn
	// positions          byte = 0
	musicPaused 	   bool
)
