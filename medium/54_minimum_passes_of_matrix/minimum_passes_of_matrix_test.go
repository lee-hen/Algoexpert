package minimum_passes_of_matrix

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{0, -1, -3, 2, 0},
		{1, -2, -5, -1, -3},
		{3, 0, 0, -4, -1},
	}
	expected := 3
	actual := MinimumPassesOfMatrix(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := [][]int{
		{-2, -3, -4, -1, -9},
		{-4, -3, -4, -1, -2},
		{-6, -7, -2, -1, -1},
		{0, 0, 0, 0, -3},
		{1, -2, -3, -6, -1},
	}
	expected := 12
	actual := MinimumPassesOfMatrix(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := [][]int{
		{-1, 2, 3},
		{4, -5, -6},
	}
	expected := 1
	actual := MinimumPassesOfMatrix(input)
	require.Equal(t, expected, actual)
}
