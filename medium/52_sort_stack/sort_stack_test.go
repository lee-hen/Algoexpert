package sort_stack

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{-5, 2, -2, 4, 3, 1}
	expected := []int{-5, -2, 1, 2, 3, 4}
	actual := SortStack(input)
	require.Equal(t, expected, actual)
}
