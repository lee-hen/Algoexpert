package zigzag_traverse

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestNormalCase(t *testing.T) {
	input := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	output := ZigzagTraverse1(input)
	require.Equal(t, expected, output)
}

// 5 * 6
func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 3, 4, 10, 11},
		{2, 5, 9, 12, 20},
		{6, 8, 13, 19, 21},
		{7, 14, 18, 22, 27},
		{15, 17, 23, 26, 28},
		{16, 24, 25, 29, 30},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	output := ZigzagTraverse(input)
	require.Equal(t, expected, output)
}

// 6 * 5
func TestCase2(t *testing.T) {
	input := [][]int{
		{1, 3, 4, 10, 11, 20},
		{2, 5, 9, 12, 19, 21},
		{6, 8, 13, 18, 22, 27},
		{7, 14, 17, 23, 26, 28},
		{15, 16, 24, 25, 29, 30},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	output := ZigzagTraverse(input)
	require.Equal(t, expected, output)
}

// 5 * 1
func TestCase3(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 4, 5},
	}
	expected := []int{1, 2, 3, 4, 5}
	output := ZigzagTraverse(input)
	require.Equal(t, expected, output)
}

// 1 * 5
func TestCase4(t *testing.T) {
	input := [][]int{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	expected := []int{1, 2, 3, 4, 5}
	output := ZigzagTraverse(input)
	require.Equal(t, expected, output)
}
