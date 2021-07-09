package line_through_points

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 1},
		{2, 2},
		{3, 3},
		{0, 4},
		{-2, 6},
		{4, 0},
		{2, 1},
	}
	expected := 4
	actual := LineThroughPoints(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := [][]int{
		{3, 3},
		{0, 4},
		{-2, 6},
		{4, 0},
		{2, 1},
		{3, 4},
		{5, 6},
		{0, 0},
	}
	expected := 3
	actual := LineThroughPoints(input)
	require.Equal(t, expected, actual)
}


func TestCase3(t *testing.T) {
	input := [][]int{
		{1, 1},
		{1, 2},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 1},
		{2, 2},
		{2, 3},
		{2, 4},
		{2, 5},
		{3, 1},
		{3, 2},
		{3, 4},
		{3, 5},
		{4, 1},
		{4, 2},
		{4, 3},
		{4, 4},
		{4, 5},
		{5, 1},
		{5, 2},
		{5, 3},
		{5, 4},
		{5, 5},
		{6, 6},
		{2, 6},
	}
	expected := 6
	actual := LineThroughPoints(input)
	require.Equal(t, expected, actual)
}

