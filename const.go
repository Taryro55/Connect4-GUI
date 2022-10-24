package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	// Paths
	LOGS_PATH	 = "logs"
	TXR_PATH	 = "textures"
	MUSIC_PATH   = "music"
	SFX_PATH	 = "sfx"

	// Game config
	HEIGHT       = int32(768)
	ROWS         = 6
	COLLUMNS     = 7
	WINDOW_TITLE = "Connect 4"

	// States
	EMPTY = 0

	// Keys
	CONF_KEY = rl.KeyTab
)