package dijkstras_algorithm

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	start := 0
	edges := [][][]int{
		{{1, 7}},
		{{2, 6}, {3, 20}, {4, 3}},
		{{3, 14}},
		{{4, 2}},
		{},
		{},
	}
	expected := []int{0, 7, 13, 27, 10, -1}
	actual := DijkstrasAlgorithm(start, edges)
	require.Equal(t, expected, actual)
}
