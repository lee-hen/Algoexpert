package repeated_matrix_values

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 3, 7, 4, 5},
		{2, 5, 9, 3, 3},
		{1, 8, 5, 3, 5},
		{5, 0, 3, 5, 6},
		{3, 8, 3, 5, 6},
		{1, 0, 3, 0, 5},
	}
	expected := []int{3, 5}
	actual := RepeatedMatrixValues(input)
	sort.Ints(actual)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := [][]int{}
	expected := []int{}
	actual := RepeatedMatrixValues(input)
	sort.Ints(actual)
	require.Equal(t, expected, actual)
}

