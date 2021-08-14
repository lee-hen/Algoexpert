package maximum_sum_submatrix
// important question

import (
	"math"
)

// MaximumSumSubmatrix
// O(w * h) time | O(w * h) space - where w is
// the width of the matrix and h is the height
func MaximumSumSubmatrix(matrix [][]int, size int) int {
	sums := createSumMatrix(matrix)
	maxSubMatrixSum := math.MinInt32

	for row := size - 1; row < len(matrix); row++ {
		for col := size - 1; col < len(matrix[row]); col++ {
			total := sums[row][col]

			touchesTopBorder := row-size < 0
			if !touchesTopBorder {
				total -= sums[row-size][col]
			}

			touchesLeftBorder := col-size < 0
			if !touchesLeftBorder {
				total -= sums[row][col-size]
			}

			touchesTopOrLeftBorder := touchesTopBorder || touchesLeftBorder
			if !touchesTopOrLeftBorder {
				total += sums[row-size][col-size]
			}

			maxSubMatrixSum = max(maxSubMatrixSum, total)
		}
	}

	return maxSubMatrixSum
}

func createSumMatrix(matrix [][]int) [][]int {
	sums := make([][]int, len(matrix))
	for i := range sums {
		sums[i] = make([]int, len(matrix[0]))
	}

	sums[0][0] = matrix[0][0]
	for idx := 1; idx < len(matrix[0]); idx++ {
		sums[0][idx] = sums[0][idx-1] + matrix[0][idx]
	}

	for idx := 1; idx < len(matrix); idx++ {
		sums[idx][0] = sums[idx-1][0] + matrix[idx][0]
	}

	for row := 1; row < len(matrix); row++ {
		for col := 1; col < len(matrix[row]); col++ {
			sums[row][col] = sums[row-1][col] + sums[row][col-1] - sums[row-1][col-1] + matrix[row][col]
		}
	}
	return sums
}

// 5, 3, -1, 5

// [5, 3, -1, 5]  [8,   2,   4]
// [-7, 3, 7, 4]  [-4,  10, 11]
// [12, 8, 0, 0]  [20,  8,   0]
// [1, -8, -8, 2] [-7, -16, -6]

// [-, -, -, -]
// [-, 3, 7, -]
// [-, 8, 0, -]
// [-, -, -, -]

//18

// height and width are unequal
// but at least size * size

// my solution
func maximumSumSubmatrix(matrix [][]int, size int) int {
	sumsMatrix := make([][]int, len(matrix), len(matrix))

	for i := range matrix {
		sumsMatrix[i] = make([]int, 0)
		for j := 0; j < len(matrix[i])-size+1; j++ {
			var currSum int
			for k := j; k < j+size; k++ {
				currSum += matrix[i][k]
			}
			sumsMatrix[i] = append(sumsMatrix[i], currSum)
		}
	}

	maxVal := math.MinInt32
	for i := 0; i < len(sumsMatrix)-size+1; i++ {
		for j := 0; j < len(sumsMatrix[i]); j++ {
			var currSum int
			for k := i; k < i+size; k++ {
				if j > len(sumsMatrix[k])-1 {
					break
				} else {
					currSum += sumsMatrix[k][j]
				}
			}
			maxVal = max(currSum, maxVal)
		}
	}
	return maxVal
}

func max(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}
