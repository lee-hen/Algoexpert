package next_greater_element

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{2, 5, -3, -4, 6, 7, 2}
	expected := []int{5, 6, 6, 6, 7, -1, 5}
	actual := NextGreaterElement(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{6, 4, 5, 7, 2, 1, 3}
	expected := []int{7, 5, 7, -1, 3, 3, 6}
	actual := NextGreaterElement(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{5, 6, 1, 3, 1, -2, -1, 3, 4, 5}
	expected := []int{6, -1, 3, 4, 3, -1, 3, 4, 5, 6}
	actual := NextGreaterElement(input)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{1, 2, 3, 4, 1, 2, 3, 4, -8, -7, 6, 2, 17, 2, -8, 9, 0, 2}
	expected := []int{2, 3, 4, 6, 2, 3, 4, 6, -7, 6, 17, 17, -1, 9, 9, 17, 2, 3}
	actual := NextGreaterElement(input)
	require.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := []int{-8, -1, -1, -2, -4, -5, -6, 0, -9, -91, -2, 8}
	expected := []int{-1, 0, 0, 0, 0, 0, 0, 8, -2, -2, 8, -1}
	actual := NextGreaterElement(input)
	require.Equal(t, expected, actual)
}


