package sorted_squared_array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{-9, -7, -3, -1, 1, 9}
	expected := []int{1, 1, 9, 49, 81, 81}
	actual := SortedSquaredArray(input)
	require.Equal(t, expected, actual)
}
