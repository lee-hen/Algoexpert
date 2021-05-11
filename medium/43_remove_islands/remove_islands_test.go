package remove_islands

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 1, 0, 1, 1, 1},
		{0, 0, 1, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 1, 1, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}

	//input := [][]int{
	//	{1, 1, 1, 1, 1, 1},
	//	{1, 1, 1, 1, 1, 1},
	//	{1, 1, 1, 1, 1, 1},
	//	{1, 1, 1, 1, 1, 1},
	//	{1, 1, 1, 1, 1, 1},
	//	{1, 1, 1, 1, 1, 1},
	//}
	expected := [][]int{
		{1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 1, 1},
		{0, 0, 0, 0, 1, 0},
		{1, 1, 0, 0, 1, 0},
		{1, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 1},
	}
	actual := RemoveIslands(input)
	require.Equal(t, expected, actual)
}

