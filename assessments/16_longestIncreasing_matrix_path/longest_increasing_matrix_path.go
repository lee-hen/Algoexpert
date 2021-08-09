package longestIncreasing_matrix_path

import (
	"fmt"
	"math"
)

// LongestIncreasingMatrixPath
// O(n) time | O(n) space - where n is the total number of elements in the matrix
func LongestIncreasingMatrixPath(matrix [][]int) int {
	longestPathLengths := [][]*int{}
	for i := 0; i < len(matrix); i++ {
		row := []*int{}
		for j := 0; j < len(matrix[0]); j++ {
			row = append(row, nil)
		}
		longestPathLengths = append(longestPathLengths, row)
	}

	longestPathLength := 0
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			longestPathLength = max(
				getLongestPathLengthAt(matrix, row, col, math.MinInt32, longestPathLengths),
				longestPathLength,
			)
		}
	}
	return longestPathLength
}

func getLongestPathLengthAt(matrix [][]int, row, col, lastPathValue int, longestPathLengths [][]*int) int {
	isOutOfBounds := row < 0 || col < 0 || row >= len(matrix) || col >= len(matrix[0])
	if isOutOfBounds {
		return 0
	}

	currentValue := matrix[row][col]
	if currentValue <= lastPathValue {
		return 0
	}

	if longestPathLengths[row][col] == nil {
		newMax := 1 + max(
			getLongestPathLengthAt(matrix, row-1, col, currentValue, longestPathLengths),
			getLongestPathLengthAt(matrix, row+1, col, currentValue, longestPathLengths),
			getLongestPathLengthAt(matrix, row, col-1, currentValue, longestPathLengths),
			getLongestPathLengthAt(matrix, row, col+1, currentValue, longestPathLengths),
		)
		longestPathLengths[row][col] = &newMax
	}
	return *longestPathLengths[row][col]
}

// [1,  2,  4,  3,  2],
// [7,  6,  5,  5,  1],
// [8,  9,  7, 15, 14],
// [5, 10, 11, 12, 13],

// see test case 4, if we choose one branch of the answer(from above below left right thw answer is 16) this question will be very hard.

func longestIncreasingMatrixPath(matrix [][]int) int {
	cache := make(map[string]int)

	var longestIncreasingPath int
	for i := range matrix {
		for j := range matrix[i] {
			key := createHashtableKey(i, j)
			cache[key] = traverseNode(i, j, cache, matrix)
			longestIncreasingPath = max(longestIncreasingPath, cache[key])
		}
	}

	return longestIncreasingPath
}

func traverseNode(i, j int, cache map[string]int, matrix [][]int) int {
	key := createHashtableKey(i, j)

	if _, ok := cache[key]; ok {
		return cache[key]
	}

	var longestIncreasingPath = 1

	if !beyondBorder(i-1, j, matrix) && matrix[i][j] < matrix[i-1][j] { // above
		longestIncreasingPath = max(longestIncreasingPath, traverseNode(i-1, j, cache, matrix)+1)
	}

	if !beyondBorder(i+1, j, matrix) && matrix[i][j] < matrix[i+1][j] { // below
		longestIncreasingPath = max(longestIncreasingPath, traverseNode(i+1, j, cache, matrix)+1)
	}

	if !beyondBorder(i, j+1, matrix) && matrix[i][j] < matrix[i][j+1] { // right
		longestIncreasingPath = max(longestIncreasingPath, traverseNode(i, j+1, cache, matrix)+1)
	}

	if !beyondBorder(i, j-1, matrix) && matrix[i][j] < matrix[i][j-1] { // left
		longestIncreasingPath = max(longestIncreasingPath, traverseNode(i, j-1, cache, matrix)+1)
	}

	cache[key] = longestIncreasingPath
	return longestIncreasingPath
}

func beyondBorder(i, j int, matrix [][]int) bool {
	if i < 0 || i > len(matrix)-1 || j < 0 || j > len(matrix[i])-1 {
		return true
	}
	return false
}

func max(a int, others ...int) int {
	m := a
	for _, other := range others {
		if other > m {
			m = other
		}
	}
	return m
}

func createHashtableKey(i, j int) string {
	return fmt.Sprintf("%d:%d", i, j)
}
