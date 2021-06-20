package quick_select

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	expected := 5
	output := QuickSelect([]int{8, 5, 2, 9, 7, 6, 3}, 3)
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := 7
	output := QuickSelect([]int{8, 5, 2, 9, 7, 6, 3}, 5)
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := 7
	output := QuickSelect([]int{8, 3, 2, 5, 1, 7, 4, 6}, 7)
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := 8
	output := quickSelect([]int{8, 3, 2, 5, 1, 7, 4, 6}, 8)
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := 1
	output := quickSelect([]int{1}, 1)
	require.Equal(t, expected, output)
}

// 0 1 2 3 4 5 6
// 2 3 5 6 7 8 9

