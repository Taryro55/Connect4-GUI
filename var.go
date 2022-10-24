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

	blinkOponentVer = Vertice{
		0,
		HEIGHT / 12 * (40 / 35 * 8),
		(HEIGHT / 8),
		(HEIGHT / 16)}
	blinkConfVer = Vertice{
		
	}
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
	gameWinner 	int32

	// screen events
	screenMenu    bool = true
	screenConf	  bool
	screenOponent bool
	screenBoard   bool

	// input events
	continuePressed bool

	// Multi Dimensional Arrays
	twoDimY	   [][]int32

	// hover events
	mainMenuHover	   int
	configHover		   int
	oponentHover	   bool


	// general events
	mainLoops		   int   // How many Main loops have happend
	runningMilisecsTimesTen	   int // how many secs have elapsed
	runningMilisecs	   float32
	runningSecs 	   int

	shouldBlink        bool
	isOponentAI        bool
	boardMade		   bool // checks if the 2D array was made.
	cursorOverBoard    bool
	debPosY            int32
	collsPos           []int
	collCurrent        int
	collHeight         []int
	movesMade          int

	y =0
)

var (
	// textures, vectors, any rl types.
	txrWhite      rl.Texture2D
	txrBlack      rl.Texture2D
	txrBoard      rl.Texture2D
	txrBackground rl.Texture2D
	txrLogo       rl.Texture2D
	music rl.Music
)
