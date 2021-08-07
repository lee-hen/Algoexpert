package spin_rings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	array := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	expected := [][]int{
		{5, 1, 2, 3},
		{9, 10, 6, 4},
		{13, 11, 7, 8},
		{14, 15, 16, 12},
	}
	SpinRings(array)
	require.Equal(t, expected, array)
}

func TestCase2(t *testing.T) {
	array := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}
	expected := [][]int{
		{6, 1, 2, 3, 4},
		{11, 12, 7, 8, 5},
		{16, 17, 13, 9, 10},
		{21, 18, 19, 14, 15},
		{22, 23, 24, 25, 20},
	}
	SpinRings(array)
	require.Equal(t, expected, array)
}
