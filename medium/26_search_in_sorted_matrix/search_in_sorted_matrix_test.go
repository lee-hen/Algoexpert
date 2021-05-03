package search_in_sorted_matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	var matrix = [][]int{
		{1, 4, 7, 12, 15, 1000},
		{2, 5, 19, 31, 32, 1001},
		{3, 8, 24, 33, 35, 1002},
		{40, 41, 42, 44, 45, 1003},
		{99, 100, 103, 106, 128, 1004},
	}
	expected := []int{3, 3}
	output := SearchInSortedMatrix(matrix, 44)
	require.Equal(t, expected, output)
}
