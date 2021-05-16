package zigzag_traverse

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 3, 4, 10},
		{2, 5, 9, 11},
		{6, 8, 12, 15},
		{7, 13, 14, 16},
	}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	output := ZigzagTraverse(input)
	require.Equal(t, expected, output)
}
