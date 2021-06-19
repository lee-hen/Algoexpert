package shifted_binary_search

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	expected := 8
	output := ShiftedBinarySearch([]int{45, 61, 71, 72, 73, 0, 1, 21, 33, 37}, 33)
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := 7
	output := ShiftedBinarySearch([]int{61, 71, 72, 73, 0, 1, 21, 33, 37, 45}, 33)
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := 2
	output := ShiftedBinarySearch([]int{5, 23, 111, 1}, 111)
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := 2
	output := ShiftedBinarySearch([]int{71, 72, 73, 0, 1, 21, 33, 37, 45, 61}, 73)
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := 0
	output := ShiftedBinarySearch([]int{45, 61, 71, 72, 73, 0, 1, 21, 33, 37}, 45)
	require.Equal(t, expected, output)
}
