package main

import (
	"math"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	height = HEIGHT

	width = (height / 9) * 16

	gridSize   = int32(math.Round((math.Hypot(float64(width), float64(height)))*0.1) * 0.6)
	gridRadius = float32(gridSize * 5 / 13)

	offsetX = width * 2 / 100
	offsetY = height * 3 / 100

	blinkOponentVer = Vertice{
		0,
		height / 12 * (40 / 35 * 8),
		(height / 8),
		(height / 16)}
	boardVer = Vertice{ // vertice and magnitude on X & Y
		(width / 2) - (gridSize*7)/2,
		height - gridSize*6 - offsetY,
		gridSize * 7,
		gridSize * 6}
	boardXtra = height * 5 / 96
)

var (
	// options
	musicPaused bool
	sfxPaused 	bool

	// main events
	debug     bool = true
	executing bool = true
	debugMenu bool

	// game events
	gameWinner int32

	// screen events
	screenMenu    bool = true
	screenConf    bool
	screenOponent bool
	screenBoard   bool

	// input events
	continuePressed bool

	// Multi Dimensional Arrays
	twoDimY [][]int32

	// hover events
	mainMenuHover    int
	configMenuHover  int
	oponentMenuHover bool

	// general events
	mainLoops               int // How many Main loops have happend
	runningMilisecsTimesTen int // how many secs have elapsed
	runningMilisecs         float32
	// runningSecs             int

	shouldBlink     bool
	isOponentAI     bool
	boardMade       bool // checks if the 2D array was made.
	cursorOverBoard bool
	debPosY         int32
	collsPos        []int32
	collCurrent     int32
	movesMade       int
	coord           int32
)

var (
	// textures, vectors, any rl types.
	txrLogo rl.Texture2D
	music   rl.Music
	sfxPop  rl.Sound
)
