package index_equals_value

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{-5, -3, 0, 3, 4, 5, 9}
	expected := 3
	actual := IndexEqualsValue(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{-5, -3, 2, 3, 4, 5, 9}
	expected := 2
	actual := IndexEqualsValue(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := 0
	actual := indexEqualsValue(input)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	expected := -1
	actual := IndexEqualsValue([]int{})
	require.Equal(t, expected, actual)
}
