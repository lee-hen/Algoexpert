package radix_sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := []int{8762, 654, 3008, 345, 87, 65, 234, 12, 2}
	expected := []int{2, 12, 65, 87, 234, 345, 654, 3008, 8762}
	actual := radixSort(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{2, 4, 8, 5, 7, 5, 4, 2, 2}
	expected := []int{2, 2, 2, 4, 4, 5, 5, 7, 8}
	actual := radixSort(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	actual := radixSort(input)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{1000, 100, 10, 1, 10, 100, 1000, 10001, 10201, 1001, 0, 1, 1}
	expected := []int{0, 1, 1, 1, 10, 10, 100, 100, 1000, 1000, 1001, 10001, 10201}
	actual := radixSort(input)
	require.Equal(t, expected, actual)
}
