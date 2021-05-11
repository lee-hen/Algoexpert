package sunset_views

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "EAST"
	expected := []int{1, 3, 6, 7}
	actual := SunsetViews(buildings, direction)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	buildings := []int{3, 5, 4, 4, 3, 1, 3, 2}
	direction := "WEST"
	expected := []int{0, 1}
	actual := SunsetViews(buildings, direction)
	require.Equal(t, expected, actual)
}
