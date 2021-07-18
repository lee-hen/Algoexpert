package palindrome_partitioning_min_cuts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	output := PalindromePartitioningMinCuts("noonabbad")
	require.Equal(t, 2, output)
}

func TestCase2(t *testing.T) {
	output := PalindromePartitioningMinCuts2("noonooon")
	require.Equal(t, 2, output)
}

func TestCase3(t *testing.T) {
	output := PalindromePartitioningMinCuts("abbb")
	require.Equal(t, 1, output)
}

