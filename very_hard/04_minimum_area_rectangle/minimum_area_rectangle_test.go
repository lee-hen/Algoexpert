package minimum_area_rectangle

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 5},
		{5, 1},
		{4, 2},
		{2, 4},
		{2, 2},
		{1, 2},
		{4, 5},
		{2, 5},
		{-1, -2},
	}
	expected := 3
	actual := MinimumAreaRectangle(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := [][]int{
		{-4, 4},
		{4, 4},
		{4, -2},
		{-4, -2},
		{0, -2},
		{4, 2},
		{0, 2},
		{0, 4},
		{2, 3},
		{0, 3},
		{2, 4},
	}
	expected := 2
	actual := MinimumAreaRectangle(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := [][]int{
		{-4, 4},
		{4, 4},
		{4, -2},
		{-4, -2},
		{0, -2},
		{4, 2},
		{0, 2},
		{0, 4},
		{2, 3},
		{2, 5},
		{2, 6},
		{0, 3},
		{0, 5},
		{0, 6},
	}
	expected := 2
	actual := MinimumAreaRectangle(input)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	expected := 0
	actual := MinimumAreaRectangle([][]int{})
	require.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := [][]int{
		{0, 0},
		{4, 4},
		{8, 8},
		{0, 8},
		{0, 4},
		{6, 0},
		{6, 4},
	}
	expected := 24
	actual := MinimumAreaRectangle(input)
	require.Equal(t, expected, actual)
}
