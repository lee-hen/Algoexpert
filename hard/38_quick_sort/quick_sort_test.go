package quick_sort

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	output := QuickSort([]int{8, 5, 2, 9, 5, 6, 3})
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []int{-10, -7, -7, -6, -6, -5, -5, -4, -4, -4, -2, -1, 1, 3, 5, 5, 6, 8, 8, 10}
	output := QuickSort([]int{-4, 5, 10, 8, -10, -6, -4, -2, -5, 3, 5, -4, -5, -1, 1, 6, -7, -6, -7, 8})
	require.Equal(t, expected, output)
}
