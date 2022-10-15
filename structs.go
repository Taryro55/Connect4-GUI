package main

// * Custom datatypes

type LogDir struct { // Structs for logging
	LogDirectory string
}
type C4 struct { // Board struct
	board [][]int32
}

type vertice struct {
	xPos int32
	yPos int32
	xMag int32
	yMag int32
}