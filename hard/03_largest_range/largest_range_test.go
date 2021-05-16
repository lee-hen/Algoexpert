package largest_range

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{0, 7}
	output := LargestRange([]int{1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12, 6})
	require.Equal(t, expected, output)
}
