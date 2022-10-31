package main

// * Checker Funcs //

func isColValid(c int32, board Board) bool { // Checks if a collumn is ful
	if -1 < c && c < board.colls {
		if board.rows-1-board.heights[c] != -1 {
			return true
		}
	}
	return false
}

// Passes a list, returns if there is a win and the player id
func isWinDetected(l []int32) int32 {
	y := 0
	x0 := int32(0)
	for _, x := range l {
		if (x == 1 || x == 2) && (x == x0 || x0 == 0) {
			x0 = x
			y += 1
			if y == 4 {
				return x
			}
		} else if x != x0 {
			y = 0
			x0 = 0
		}
	}
	return EMPTY
}

// * End of Checker Funcs //

// * Getters(?) Funcs //
func get_Y_Board(board Board) [][]int32 {
	g := [][]int32{}
	for y := 0; y < 7; y++ {
		z := []int32{}
		for x := 0; x < 6; x++ {
			z = append(z, board.grid[x][y])
		}
		g = append(g, z)
	}
	return g
}

// creates two diagonal 2d arrays of the - and + slopes
func getDiagBoard(b Board) ([][]int32, [][]int32) {
	p0 := 0
	t, g := []int32{}, []int32{}
	r, l := [][]int32{}, [][]int32{}
	for p := 0; p <= 12; p++ {
		for ir := range b.grid {
			for ic := range b.grid[ir] {
				if p0 != p {
					r, l = append(r, t), append(l, g)
					t, g = nil, nil
				}
				if ir+ic == p {
					coord = b.grid[ir][ic]
					t = append(t, coord)
				}
				if ir-ic+6 == p {
					coord = b.grid[ir][ic]
					g = append(g, coord)
				}
				p0 = p
			}
		}
	}
	return r, l
}

func getRow(c int32, b Board) int32 {
	return b.heights[c]
}

func getHoveringCol(i []int32, intg int32, X int32) int32 {
	for k, v := range i {
		if X-v < intg {
			return int32(k)
		}
	}
	return -1
}

func getValidColls(board Board) []int32 {
	v := (make([]int32, board.colls))
	for i, _ := range board.grid {
		if isColValid(int32(i), board) {
			v = append(v, int32(i))
		}
	}
	return v
}

func getWinner(b Board) {
	switch getConnectedFours(b) {
	case 0:
		gameWinner = 0
	case 1:
		gameWinner = 1
	case 2:
		gameWinner = 2
	}
}

func getConnectedFours(board Board) int32 {

	// Detects on Y
	twoDimY = get_Y_Board(board)
	for col := range twoDimY {
		if isWinDetected(twoDimY[col]) != EMPTY {
			return isWinDetected(twoDimY[col])
		}
	}

	// Detects on X
	for row := range board.grid {
		if isWinDetected(board.grid[row]) != EMPTY {
			return isWinDetected(board.grid[row])
		}
	}

	// Detects Diag
	diagTtR, diagTtL := getDiagBoard(board)
	for row := range diagTtR {
		if isWinDetected(diagTtR[row]) != EMPTY {
			return isWinDetected(diagTtR[row])
		}
	}

	for row := range diagTtL {
		if isWinDetected(diagTtL[row]) != EMPTY {
			return isWinDetected(diagTtL[row])
		}
	}
	return 0
}

// * End of Getters(?) Funcs //

// * Board Funcs //

// Creates a board
func boardMake(r, c int32) Board {
	board := Board{
		grid:    make([][]int32, r),
		heights: make([]int32, r+1),
		colls:   c,
		rows:    r,
	}
	for row := range board.grid {
		board.grid[row] = make([]int32, c)
	}
	return board
}

// Resets all values of a board
func boardReset(board Board) bool {
	for row := range board.grid {
		for col := range board.grid[row] {
			board.grid[row][col] = EMPTY
		}
	}
	return true
}

// Changes the state of one of the points on the grid
func boardDropPiece(b Board, col, state int32) bool {
	b.grid[getRow(col, b)][col] = state
	return true
}

// * End of Board Funcs //
