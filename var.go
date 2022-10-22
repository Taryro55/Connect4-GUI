package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	width = (HEIGHT / 9) * 16

	gridSize   = int32(math.Round((math.Hypot(float64(width), float64(HEIGHT)))*0.1) * 0.6)
	gridRadius = float32(gridSize * 5 / 13)

	offsetX = width * 2 / 100
	offsetY = HEIGHT * 3 / 100

	blinkVer = Vertice{
		0,
		HEIGHT / 12 * (40 / 35 * 8),
		(HEIGHT / 8),
		(HEIGHT / 16)}
	boardVer = Vertice{ // vertice and magnitude on X & Y
		(width / 2) - (gridSize*7)/2,
		HEIGHT - gridSize*6 - offsetY,
		gridSize * 7,
		gridSize * 6}
	boardXtra = HEIGHT * 5 / 96
)

var (
	// options
	musicPaused bool

	// main events
	debug     bool = true
	executing bool = true
	debugMenu bool

	// game events
	gameOver    bool
	gameDraw    bool
	gameOngoing bool

	// general events
	mouseButtonPressed bool
	shouldBlink        bool
	isOponentAI        bool
	oponentSelected    bool
	boardRendered      bool
	firstLoop          bool = true
	runningTimeInt     int
	runningTime        float64
	debPosY            int32
	collsPos           []int
	collCurrent        int
	collHeight         []int
	movesMade          int

	Colls2dSlice	   [][]int32
	y =0
	winner 			   int32
)

var (
	// textures, vectors, any rl types.
	txrWhite      rl.Texture2D
	txrBlack      rl.Texture2D
	txrBoard      rl.Texture2D
	txrBackground rl.Texture2D
	txrLogo       rl.Texture2D
	// mousePos    rl.Vector2
	music rl.Music
)
