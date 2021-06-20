package search_for_range

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{4, 9}
	output := SearchForRange([]int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}, 45)
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []int{0, 0}
	output := SearchForRange([]int{5, 7, 7, 8, 8, 10}, 5)
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := []int{5, 5}
	output := SearchForRange([]int{5, 7, 7, 8, 8, 10}, 10)
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := []int{1, 2}
	output := SearchForRange([]int{5, 7, 7, 8, 8, 10}, 7)
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := []int{4, 12}
	output := SearchForRange([]int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 45, 45, 45}, 45)
	require.Equal(t, expected, output)
}

func TestCase6(t *testing.T) {
	expected := []int{5, 5}
	output := SearchForRange([]int{5, 7, 7, 8, 8, 10}, 10)
	require.Equal(t, expected, output)
}
