package main

import rl "github.com/gen2brain/raylib-go/raylib"

// * Custom datatypes

type LogDir struct { // Structs for logging
	LogDirectory string
}

type Gamer struct {
	ID			int32
	CellColour	rl.Color
}
type C4 struct { // Board struct
	P1	  Gamer
	P2	  Gamer
	AI	  Gamer
	turn  int32

	gamers []Gamer
	board Board
}

type Vertice struct {
	xPos int32
	yPos int32
	xMag int32
	yMag int32
}

type Board struct {
	grid [][]int32
	heights []int32
	colls int32
	rows int32
	
	moves int32
}