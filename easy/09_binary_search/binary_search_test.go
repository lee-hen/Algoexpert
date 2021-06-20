package binary_search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := 3
	output := BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 33)
	require.Equal(t, expected, output)
}
