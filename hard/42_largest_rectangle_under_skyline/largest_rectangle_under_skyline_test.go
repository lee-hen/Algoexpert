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
