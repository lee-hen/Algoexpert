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

func TestCase2(t *testing.T) {
	input := [][]int{
		{1, 2, 5},
		{0, 2},
		{0, 1},
		{4, 5},
		{3, 5},
		{0, 3, 4},
	}
	expected := false
	actual := TwoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := [][]int{
		{1, 2},
		{0, 2, 3},
		{1, 3, 0},
		{1, 2},
	}
	expected := true
	actual := TwoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := [][]int{
		{1},
		{0, 2, 3},
		{1, 3},
		{1, 2},
	}
	expected := false
	actual := TwoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := [][]int{
		{1, 2, 3, 5},
		{0, 2},
		{0, 1},
		{0, 4, 5},
		{3, 5},
		{3, 4, 0},
	}
	expected := true
	actual := TwoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}

func TestCase6(t *testing.T) {
	input := [][]int{
		{1, 2, 3},
		{0, 2},
		{0, 1},
		{0, 4, 5},
		{3, 5},
		{3, 4},
	}
	expected := false
	actual := TwoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}

func TestCase7(t *testing.T) {
	input := [][]int{
		{1, 2},
		{0, 2, 3},
		{1, 0, 4},
		{1, 4},
		{3, 2},
	}
	expected := true
	actual := TwoEdgeConnectedGraph(input)
	require.Equal(t, expected, actual)
}
