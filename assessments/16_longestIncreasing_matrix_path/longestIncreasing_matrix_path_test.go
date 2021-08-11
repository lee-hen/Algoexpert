package longestIncreasing_matrix_path

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	matrix := [][]int{
		{1, 2, 4, 3, 2},
		{7, 6, 5, 5, 1},
		{8, 9, 7, 15, 14},
		{5, 10, 11, 12, 13},
	}
	expected := 15
	actual := LongestIncreasingMatrixPath(matrix)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	matrix := [][]int{
		{0, 1, 2, 3, 4},
		{1, 2, 4, 3, 2},
		{7, 6, 5, 5, 1},
		{8, 9, 7, 15, 14},
		{5, 10, 11, 12, 13},
	}
	expected := 15
	actual := LongestIncreasingMatrixPath(matrix)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	matrix := [][]int{
		{1, 2, -1, -2, 2},
		{7, 6, 5, 5, 1},
		{8, 9, 7, 15, 14},
		{5, 10, 11, 12, 11},
	}
	expected := 11
	actual := LongestIncreasingMatrixPath(matrix)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	matrix := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{6, 4, 3, 0, 0, 0, 0, 0, 0, 0},
		{9, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{10, 11, 0, 0, 27, 34, 0, 0, 0, 0},
		{0, 14, 16, 22, 25, 37, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 49, 0, 0, 0, 0},
		{0, 0, 0, 53, 52, 51, 0, 0, 0, 0},
	}
	expected := 19
	actual := LongestIncreasingMatrixPath(matrix)
	require.Equal(t, expected, actual)
}
