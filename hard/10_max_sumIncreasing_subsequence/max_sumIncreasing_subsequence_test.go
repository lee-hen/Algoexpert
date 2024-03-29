package max_sumIncreasing_subsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase2(t *testing.T) {
	outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{10, 15, 4, 5, 11, 14, 31, 25, 31, 23, 25, 31, 50})
	require.Equal(t, 164, outputSum)
	require.Equal(t, []int{10, 11, 14, 23, 25, 31, 50}, outputSequence)
}

func TestCase1(t *testing.T) {
	outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{10, 70, 20, 30, 50, 11, 30})
	require.Equal(t, 110, outputSum)
	require.Equal(t, []int{10, 20, 30, 50}, outputSequence)
}

func TestCase3(t *testing.T) {
	outputSum, outputSequence := MaxSumIncreasingSubsequence([]int{5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35})
	require.Equal(t, 69, outputSum)
	require.Equal(t, []int{5, 7, 10, 12, 35}, outputSequence)
}
