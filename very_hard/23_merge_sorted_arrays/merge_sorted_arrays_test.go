package merge_sorted_arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	arrays := [][]int{
		{1, 5, 9, 21},
		{-1, 0},
		{-124, 81, 121},
		{3, 6, 12, 20, 150},
	}
	expected := []int{-124, -1, 0, 1, 3, 5, 6, 9, 12, 20, 21, 81, 121, 150}
	actual := MergeSortedArrays(arrays)
	require.Equal(t, expected, actual)
}
