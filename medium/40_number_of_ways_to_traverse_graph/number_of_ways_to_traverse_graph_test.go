package number_of_ways_to_traverse_graph

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	width := 3
	height := 4
	expected := 10
	actual := NumberOfWaysToTraverseGraph2(width, height)
	require.Equal(t, expected, actual)
}
