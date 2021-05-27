package maximum_sum_submatrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	matrix := [][]int{
		{5, 3, -1, 5},
		{-7, 3, 7, 4},
		{12, 8, 0, 0},
		{1, -8, -8, 2},
	}
	size := 2
	expected := 18
	actual := MaximumSumSubmatrix(matrix, size)
	require.Equal(t, expected, actual)
}
