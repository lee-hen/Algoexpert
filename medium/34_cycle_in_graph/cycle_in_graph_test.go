package cycle_in_graph

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1},
		{2, 3, 4, 5, 6, 7},
		{},
		{2, 7},
		{5},
		{},
		{4},
		{},
	}
	expected := false
	actual := CycleInGraph(input)
	require.Equal(t, expected, actual)
}
