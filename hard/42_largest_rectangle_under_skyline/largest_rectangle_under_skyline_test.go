package largest_rectangle_under_skyline

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{1, 3, 3, 2, 4, 1, 5, 3, 2}
	expected := 9
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{3, 3, 3, 4, 4, 4, 1, 3, 1, 2, 8, 9, 1}
	expected := 18
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{2, 2, 30, 20, 20, 2, 2, 1}
	expected := 60
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 11}
	expected := 12
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := []int{10, 21}
	expected := 21
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	input := []int{11, 21}
	expected := 22
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase7(t *testing.T) {
	input := []int{10, 1, 2, 3, 4, 5, 6, 7}
	expected := 16
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase8(t *testing.T) {
	input := []int{10, 1, 2, 3, 3, 5, 6, 7}
	expected := 15
	actual := LargestRectangleUnderSkyline(input)
	require.Equal(t, expected, actual)
}

func TestCase9(t *testing.T) {
	expected := 0
	actual := LargestRectangleUnderSkyline([]int{})
	require.Equal(t, expected, actual)
}
