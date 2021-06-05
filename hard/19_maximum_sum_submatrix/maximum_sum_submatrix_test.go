package maximum_sum_submatrix

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	matrix := [][]int{
		{5, 3, -1, 5},
		{-7, 3, 7, 4},
		{12, 8, 0, 0},
		{1, -8, -8, 2},
	}
	size := 3
	expected := 30
	actual := MaximumSumSubmatrix(matrix, size)
	require.Equal(t, expected, actual)
}

func MaximumSumSubArray(array []int, size int) int {
	sums := make([]int, len(array))
	for idx := 1; idx < len(array); idx++ {
		sums[idx] = sums[idx-1] + array[idx]
	}

	maxSum := math.MinInt32
	for idx := size - 1; idx < len(array); idx++ {
		total := sums[idx]

		if idx >= size {
			total -= sums[idx-size]
		}

		maxSum = max(maxSum, total)
	}

	return maxSum
}

func TestCase2(t *testing.T) {
	array := []int{-7, 3, 7, 4}
	size := 2
	expected := 11
	actual := MaximumSumSubArray(array, size)
	require.Equal(t, expected, actual)
}
