package solve_sudoku
// important question

//    0  1  2   3  4  5   6  7  8
// 0 [7, 8, 0,| 4, 0, 0,| 1, 2, 0],
// 1 [6, 0, 0,| 0, 7, 5,| 0, 0, 9],
// 2 [0, 0, 0,| 6, 0, 1,| 0, 7, 8],
//   [-, -, -,| -, -, -,| -, -, -],
// 3 [0, 0, 7,| 0, 4, 0,| 2, 6, 0],
// 4 [0, 0, 1,| 0, 5, 0,| 9, 3, 0],
// 5 [9, 0, 4,| 0, 6, 0,| 0, 0, 5],
//   [-, -, -,| -, -, -,| -, -, -],
// 6 [0, 7, 0,| 3, 0, 0,| 0, 1, 2],
// 7 [1, 2, 0,| 0, 0, 7,| 4, 0, 0],
// 8 [0, 4, 9,| 2, 0, 6,| 0, 0, 7],


// [7, 8, 5, 4, 3, 9, 1, 2, 6],
// [6, 1, 2, 8, 7, 5, 3, 4, 9],
// [4, 9, 3, 6, 2, 1, 5, 7, 8],
// [8, 5, 7, 9, 4, 3, 2, 6, 1],
// [2, 6, 1, 7, 5, 8, 9, 3, 4],
// [9, 3, 4, 1, 6, 2, 7, 8, 5],
// [5, 7, 8, 3, 9, 4, 6, 1, 2],
// [1, 2, 6, 5, 8, 7, 4, 9, 3],
// [3, 4, 9, 2, 1, 6, 8, 5, 7]

func SolveSudoku(board [][]int) [][]int {
	solvePartialSudoku(0, 0, board)

	return board
}

func solvePartialSudoku(i, j int, board [][]int) bool {
	if j == 9 {
		i, j = i+1, 0
		if i == 9 {
			return true
		}
	}

	if board[i][j] == 0 {
		for digit := 1; digit <= 9; digit++ {
			if isValidAtPosition(digit, i, j, board) {  // back tracking
				board[i][j] = digit

				// ex. i: row, j: col, if i is 0
				// when j is 4, and it's(board[0][4]) digit is temporarily valid, start from 4, j plus 1,
				// then j is 5, but the digit of board[0][5] is invalid, reset the digit of board[0][5] to zero(57行目),
				// and return false(58行目), then the current position of j is 4
				// then retry the digit from previous digit at the position of board[0][4](43行目)
				if solvePartialSudoku(i, j+1, board) { // back tracking
					return true
				}
			}
		}

		board[i][j] = 0
		return false  // back tracking
	}

	return solvePartialSudoku(i, j+1, board)
}

func isValidAtPosition(value int, i int, j int, board [][]int) bool {
	if rowContains(i, value, board) || columnContains(j, value, board) {
		return false
	}

	subgridRowStart := (i / 3) * 3
	subgridColStart := (j / 3) * 3

	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			rowToCheck := subgridRowStart + row
			colToCheck := subgridColStart + col
			existingValue := board[rowToCheck][colToCheck]

			if existingValue == value {
				return false
			}
		}
	}

	return true
}

func rowContains(i, value int, board [][]int) bool {
	for _, element := range board[i] {
		if value == element {
			return true
		}
	}

	return false
}

func columnContains(j, value int, board [][]int) bool {
	for _, row := range board {
		if row[j] == value {
			return true
		}
	}

	return false
}

func solveSudoku(board [][]int) [][]int {
	emptySquare := make([][]bool, len(board), len(board))

	for i := range board {
		emptySquare[i] = make([]bool, len(board[i]), len(board[i]))
		for j := range board[i] {
			if board[i][j] == 0 {
				emptySquare[i][j] = true
			}
		}
	}

	var i, j int
	for i < 9 && j < 9 {
		if board[i][j] == 0 {
			if !tryDigit(i,j, board) {
				i, j = retryDigit(i, j, board, emptySquare)
			} else {
				i, j = nextPosition(i, j)
			}
		} else {
			i, j = nextPosition(i, j)
		}
	}

	return board
}

func tryDigit(i, j int, board [][]int) bool {
	for d := board[i][j]+1; d <= 9; d++ {
		if validDigitBoard(i, j, d, board) {
			board[i][j] = d
			return true
		}
	}

	board[i][j] = 0
	return false
}

func retryDigit(i, j int, board [][]int, emptySquare [][]bool) (int, int){
	i, j = prevPosition(i, j)
	if emptySquare[i][j] {
		if tryDigit(i, j, board) {
			return i, j
		}
	}

	return retryDigit(i, j, board, emptySquare)
}

func prevPosition(i, j int) (int, int) {
	if j > 0 {
		j--
	} else  {
		j = 8
		i--
	}
	return i, j
}

func nextPosition(i, j int) (int, int) {
	if j < 8 {
		j++
	} else {
		j = 0
		i++
	}
	return i, j
}

func validDigitBoard(i, j, digit int, board [][]int) bool {
	var startRow, startCol int

	if i < 3  {
		startRow = 0

		if j < 3 {
			startCol = 0
		}

		if j > 2 && j < 6 {
			startCol = 3
		}

		if j > 5 && j < 9 {
			startCol = 6
		}
	}

	if i > 2 && i < 6 {
		startRow = 3

		if j < 3 {
			startCol = 0
		}

		if j > 2 && j < 6 {
			startCol = 3
		}

		if j > 5 && j < 9 {
			startCol = 6
		}
	}


	if i > 5 && i < 9 {
		startRow = 6

		if j < 3 {
			startCol = 0
		}

		if j > 2 && j < 6 {
			startCol = 3

		}

		if j > 5 && j < 9 {
			startCol = 6
		}
	}

	return validGrid(i, j, digit, startRow, startCol, startRow+2, startCol+2, board) &&
		validRow(i, j, digit, board) && validCol(i, j, digit, board)
}

func validGrid(i, j, digit, startRow, startCol, endRow, endCol int, board [][]int) bool {
	for row := startRow; row <= endRow; row++ {
		if row == i {
			continue
		}

		for col := startCol; col <= endCol; col++ {
			if col == j {
				continue
			}

			if board[row][col] == digit {
				return false
			}
		}
	}

	return true
}

func validRow(i, j, digit int, board [][]int) bool {
	m, n := 0, 8

	for m < j {
		if board[i][m] == digit {
			return false
		}
		m++
	}

	for n > j {
		if board[i][n] == digit {
			return false
		}
		n--
	}
	return true
}


func validCol(i, j, digit int, board [][]int) bool {
	m, n := 0, 8

	for m < i {
		if board[m][j] == digit {
			return false
		}
		m++
	}

	for n > i {
		if board[n][j] == digit {
			return false
		}
		n--
	}
	return true
}
