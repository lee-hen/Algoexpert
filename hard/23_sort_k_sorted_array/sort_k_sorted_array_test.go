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


func TestCase2(t *testing.T) {
	input := []int{5, 4, 3, 2, -100}
	k := 5
	expected := []int{-100, 2, 3, 4, 5}
	actual := SortKSortedArray(input, k)
	require.Equal(t, expected, actual)
}


func TestCase3(t *testing.T) {
	input := []int{-2, -3, 1, 2, 3, 1, 1, 2, 3, 8, 100, 130, 9, 12}
	k := 4
	expected := []int{-3, -2, 1, 1, 1, 2, 2, 3, 3, 8, 9, 12, 100, 130}
	actual := SortKSortedArray(input, k)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 1}
	k := 8
	expected := []int{1, 1, 2, 3, 4, 5, 6}
	actual := SortKSortedArray(input, k)
	require.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := []int{1, 0, 1, 1, 1, 1, 1, 1}
	k := 1
	expected := []int{0, 1, 1, 1, 1, 1, 1, 1}
	actual := SortKSortedArray(input, k)
	require.Equal(t, expected, actual)
}
