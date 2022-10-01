package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	twhite      rl.Texture2D
	tblack      rl.Texture2D
	tboard      rl.Texture2D
	tbackground rl.Texture2D
	// mousePos    rl.Vector2
	music       rl.Music
)

var (
	debug 			   bool = true
	executing          bool = true
	gameOver           bool = false
	gameDraw           bool = false
	gameOngoing        bool = false
	mouseButtonPressed bool = false
	// playerMove         bool = false // If bool false, its whites turn, if true, its blacks turn
	//positions          byte = 0
	musicPaused 	   bool
)
