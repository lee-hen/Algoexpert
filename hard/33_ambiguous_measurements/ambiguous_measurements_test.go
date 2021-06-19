package ambiguous_measurements

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	cups := [][]int{
		{200, 210},
		{450, 465},
		{800, 850},
	}
	low := 2100
	high := 2300
	expected := true
	actual := AmbiguousMeasurements(cups, low, high)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	cups := [][]int{
		{1, 3},
		{2, 4},
		{5, 6},
	}
	low := 100
	high := 101
	expected := false
	actual := AmbiguousMeasurements(cups, low, high)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	cups := [][]int{
		{50, 65},
		{100, 120},
		{20, 40},
		{10, 15},
		{400, 500},
		{300, 350},
		{10, 25},
	}

	low := 3000
	high := 3300
	expected := false
	actual := AmbiguousMeasurements(cups, low, high)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	cups := [][]int{
		{50, 65},
		{100, 120},
		{20, 40},
		{10, 15},
		{400, 500},
		{300, 350},
		{10, 25},
	}

	low := 3000
	high := 3300
	expected := false
	actual := ambiguousMeasurements(cups, low, high)
	require.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	cups := [][]int{
		{50, 60},
		{100, 120},
		{20, 40},
		{400, 500},
	}

	low := 1000
	high := 1050
	expected := false
	actual := AmbiguousMeasurements(cups, low, high)
	require.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	cups := [][]int{
		{1, 3},
		{2, 4},
		{5, 6},
	}
	low := 100
	high := 120
	expected := true
	actual := AmbiguousMeasurements(cups, low, high)
	require.Equal(t, expected, actual)
}