package two_edge_connected_graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 2, 5},
		{0, 2},
		{0, 1, 3},
		{2, 4, 5},
		{3, 5},
		{0, 3, 4},
	}
	expected := true
	actual := TwoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}