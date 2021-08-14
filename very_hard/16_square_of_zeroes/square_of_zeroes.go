package square_of_zeroes
// important question

import "fmt"

type InfoEntry struct {
	NumZeroesRight int
	NumZeroesBelow int
}

// SquareOfZeroes
// O(n^3) time | O(n^3) space - where n is the height and width of the matrix
func SquareOfZeroes(matrix [][]int) bool {
	infoMatrix := preComputeNumOfZeroes(matrix)
	lastIdx := len(matrix) - 1
	return hasSquareOfZeroes(infoMatrix, 0, 0, lastIdx, lastIdx, map[string]bool{})
}

// r1 is the top row, c1 is the left column
// r2 is the bottom row, c2 is the right column
func hasSquareOfZeroes(infoMatrix [][]InfoEntry, r1, c1, r2, c2 int, cache map[string]bool) bool {
	if r1 >= r2 || c1 >= c2 {
		return false
	}

	key := fmt.Sprintf("%d-%d-%d-%d", r1, c1, r2, c2)
	if out, found := cache[key]; found {
		return out
	}

	cache[key] = isSquareOfZeroes(infoMatrix, r1, c1, r2, c2) ||
		hasSquareOfZeroes(infoMatrix, r1+1, c1+1, r2-1, c2-1, cache) ||
		hasSquareOfZeroes(infoMatrix, r1, c1+1, r2-1, c2, cache) ||
		hasSquareOfZeroes(infoMatrix, r1+1, c1, r2, c2-1, cache) ||
		hasSquareOfZeroes(infoMatrix, r1+1, c1+1, r2, c2, cache) ||
		hasSquareOfZeroes(infoMatrix, r1, c1, r2-1, c2-1, cache)

	return cache[key]
}

func preComputeNumOfZeroes(matrix [][]int) [][]InfoEntry {
	infoMatrix := make([][]InfoEntry, len(matrix))
	for i, row := range matrix {
		infoMatrix[i] = make([]InfoEntry, len(row))
	}

	n := len(matrix)
	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			numZeroes := 0
			if matrix[row][col] == 0 {
				numZeroes = 1
			}
			infoMatrix[row][col] = InfoEntry{
				NumZeroesBelow: numZeroes,
				NumZeroesRight: numZeroes,
			}
		}
	}

	lastIdx := len(matrix) - 1
	for row := n - 1; row >= 0; row-- {
		for col := n - 1; col >= 0; col-- {
			if matrix[row][col] == 1 {
				continue
			}

			if row < lastIdx {
				infoMatrix[row][col].NumZeroesBelow += infoMatrix[row+1][col].NumZeroesBelow
			}

			if col < lastIdx {
				infoMatrix[row][col].NumZeroesRight += infoMatrix[row][col+1].NumZeroesRight
			}
		}
	}
	return infoMatrix
}

// r1 is the top row, c1 is the left column
// r2 is the bottom row, c2 is the right column
func isSquareOfZeroes(infoMatrix [][]InfoEntry, r1, c1, r2, c2 int) bool {
	squareLength := c2 - c1 + 1
	hasTopBorder := infoMatrix[r1][c1].NumZeroesRight >= squareLength
	hasLeftBorder := infoMatrix[r1][c1].NumZeroesBelow >= squareLength
	hasBottomBorder := infoMatrix[r2][c1].NumZeroesRight >= squareLength
	hasRightBorder := infoMatrix[r1][c2].NumZeroesBelow >= squareLength
	return hasTopBorder && hasLeftBorder && hasBottomBorder && hasRightBorder
}

// SquareOfZeroes1
// O(n^3) time | O(n^2) space - where n is the height and width of the matrix
func SquareOfZeroes1(matrix [][]int) bool {
	infoMatrix := preComputeNumOfZeroes(matrix)
	n := len(matrix)
	for topRow := 0; topRow < n; topRow++ {
		for leftCol := 0; leftCol < n; leftCol++ {
			squareLength := 2
			for squareLength <= n-leftCol && squareLength <= n-topRow {
				bottomRow := topRow + squareLength - 1
				rightCol := leftCol + squareLength - 1
				if isSquareOfZeroes(infoMatrix, topRow, leftCol, bottomRow, rightCol) {
					return true
				}
				squareLength += 1
			}
		}
	}
	return false
}

//[1, 1, 1, 0, 1 ,0],
//[0, 0, 0, 0, 0, 1],
//[0, 1, 1, 1, 0, 0],
//[0, 0, 0, 1, 0, 0],
//[0, 1, 1, 1, 0, 1],
//[0, 0, 0, 0, 0, 1],

// squareOfZeroes
// naive approach  O(2n^2+2n)*n time | O(2n^2+2n)*n space
func squareOfZeroes(matrix [][]int) bool {
	if len(matrix) < 2 {
		return false
	}
	if len(matrix[0]) < 2 {
		return false
	}

	//visit := make([][]bool, len(matrix), len(matrix))
	//
	//for i := range matrix {
	//	visit[i] = make([]bool, len(matrix[i]), len(matrix[i]))
	//}

	for i, row := range matrix {
		for j, digit := range row {
			//if visit[i][j] == true {
			//	continue
			//}

			if digit == 0 {
				if i < len(matrix)-1 && j < len(matrix[i])-1 && matrix[i][j+1] == 0 && matrix[i+1][j] == 0 {
					if dfs(i,j+1, i, j,-1,-1, matrix) {
						return true
					}
				}
			}
		}
	}

	return false
}

func dfs(i, j, rowStart, colStart, rowEnd, colEnd int, matrix [][]int) bool {
	if i == rowStart && j == colStart {
		return true
	}

	var nextRow, nextCol int
	if i == rowStart {
		if j < len(matrix[i])-1 && matrix[i][j+1] == 0 && matrix[i+1][j] == 1 {
			// go right
			nextRow = i
			nextCol = j+1
		} else {
			// go down
			nextRow = i+1
			nextCol = j
			colEnd = j
		}

		if matrix[nextRow][nextCol] == 1 {
			return false
		}

		if dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix) {
			return true
		}

		if j < len(matrix[i])-1 && matrix[i][j+1] == 0 {
			// go right
			nextRow = i
			nextCol = j+1
		}

		return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix)
	}

	if j == colEnd {
		if i < len(matrix)-1 && matrix[i+1][j] == 0 && matrix[i][j-1] == 1 {
			// go down
			nextRow = i+1
			nextCol = j
		} else {
			// go left
			nextRow = i
			nextCol = j-1
			rowEnd = i
		}

		if matrix[nextRow][nextCol] == 1 {
			return false
		}

		if dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix) {
			return true
		}

		if i < len(matrix)-1 && matrix[i+1][j] == 0 {
			// go down
			nextRow = i+1
			nextCol = j
		}

		return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix)
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

		if matrix[nextRow][nextCol] == 1 {
			return false
		}

		return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix)
	}

	if j == colStart {
		if i > 0 && matrix[i-1][j] == 0 {
			// go up
			nextRow = i-1
			nextCol = j
		}

		if matrix[nextRow][nextCol] == 1 {
			return false
		}

		return dfs(nextRow, nextCol, rowStart, colStart, rowEnd, colEnd, matrix)
	}

	return false
}
