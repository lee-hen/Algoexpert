package sort_k_sorted_array

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{3, 2, 1, 5, 4, 7, 6, 5}
	k := 3
	expected := []int{1, 2, 3, 4, 5, 5, 6, 7}
	actual := SortKSortedArray(input, k)
	require.Equal(t, expected, actual)
}
