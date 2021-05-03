package cycle_in_graph

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 3},
		{2, 3, 4},
		{0},
		{},
		{2, 5},
		{},
	}
	expected := true
	actual := CycleInGraph(input)
	require.Equal(t, expected, actual)
}
