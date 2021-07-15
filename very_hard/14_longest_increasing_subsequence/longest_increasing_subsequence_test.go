package longest_increasing_subsequence

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{-24, 2, 3, 5, 6, 35}
	input := []int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35}
	output := LongestIncreasingSubsequence(input)
	require.Equal(t, expected, output)
}
