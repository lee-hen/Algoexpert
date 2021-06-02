package water_area

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 49
	output := WaterArea([]int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3, 0, 1})
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := 49
	output := WaterArea([]int{0, 8, 0, 0, 10, 0, 0, 10, 0, 0, 1, 1, 0, 3})
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := 39
	output := WaterArea([]int{0, 100, 0, 0, 10, 1, 1, 10, 1, 0, 1, 1, 0, 0})
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := 4
	output := WaterArea([]int{0, 1, 0, 1, 0, 2, 0, 3})

	// 0, 1, 0, 1, 0, 2, 0, 3
	// 0  0  1  0  1  0  2  0
  	// 0  0  1  0  1  0  3  0

	require.Equal(t, expected, output)
}
