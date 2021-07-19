package square_of_zeroes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 0, 1},
		{0, 0, 0, 1, 0, 1},
		{0, 1, 1, 1, 0, 1},
		{0, 0, 0, 0, 0, 1},
	}
	require.True(t, SquareOfZeroes(input))
}

func TestCase2(t *testing.T) {
	input := [][]int{
		{1, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 0, 1},
		{1, 1, 1, 1},
	}
	require.True(t, SquareOfZeroes(input))
}


func TestCase3(t *testing.T) {
	input := [][]int{
		{0, 0, 0},
		{1, 0, 0},
		{0, 0, 0},
	}
	require.True(t, SquareOfZeroes(input))
}
