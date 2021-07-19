package square_of_zeroes

//[1, 1, 1, 0, 1 ,0],
//[0, 0, 0, 0, 0, 1],
//[0, 1, 1, 1, 0, 0],
//[0, 0, 0, 1, 0, 0],
//[0, 1, 1, 1, 0, 1],
//[0, 0, 0, 0, 0, 1],

// SquareOfZeroes
// naive approach
func SquareOfZeroes(matrix [][]int) bool {
	if len(matrix) < 2 {
		return false
	}
	if len(matrix[0]) < 2 {
		return false
	}

	visit := make([][]bool, len(matrix), len(matrix))

	for i := range matrix {
		visit[i] = make([]bool, len(matrix[i]), len(matrix[i]))
	}

	for i, row := range matrix {
		for j, digit := range row {
			if visit[i][j] == true {
				continue
			}


			if digit == 0 {
				if i < len(matrix)-1 && j < len(matrix[i]) -1 && matrix[i][j+1] == 0 {
					if dfs(i,j+1, i, j,-1,-1, matrix, visit) {
						return true
					}
				}
			}

		}
	}

	return false
}

func dfs(i, j, rowStart, colStart, rowEnd, colEnd int, matrix [][]int, visit [][]bool) bool {
	if i == rowStart && j == colStart {
		return true
	}

	if i < len(matrix)-1 && j < len(matrix[i]) -1 {
		if !(matrix[i][j+1] == 0 && matrix[i+1][j] == 0) {
			visit[i][j] = true
		}
	}


	var nextRow, nextCol int
	if i == rowStart {
		if j < len(matrix[i])-1 && matrix[i][j+1] == 0 {
			// go right
			nextRow = i
			nextCol = j+1
		} else {
			// go down
			nextRow = i+1
			nextCol = j
			colEnd = j
		}

		if matrix[nextRow][nextCol] == 0 {
			return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix, visit)
		} else {
			return false
		}
	}

	if j == colEnd {
		if i < len(matrix)-1 && matrix[i+1][j] == 0 {
			// go down
			nextRow = i+1
			nextCol = j
		} else {
			// go left
			nextRow = i
			nextCol = j-1
			rowEnd = i
		}

		if matrix[nextRow][nextCol] == 0 {
			return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix, visit)
		} else {
			return false
		}
	}


	if i == rowEnd {
		if j > 0 && matrix[i][j-1] == 0 && j != colStart {
			// go left
			nextRow = i
			nextCol = j-1
		} else {
			// go up
			nextRow = i-1
			nextCol = j
		}
		if matrix[nextRow][nextCol] == 0 {
			return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix, visit)
		} else {
			return false
		}
	}

	if j == colStart {
		if i > 0 && matrix[i-1][j] == 0 {
			// go up
			nextRow = i-1
			nextCol = j
		}

		if matrix[nextRow][nextCol] == 0 {
			return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix, visit)
		} else {
			return false
		}
	}

	return false
}
