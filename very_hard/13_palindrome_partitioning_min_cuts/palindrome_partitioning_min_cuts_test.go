package palindrome_partitioning_min_cuts

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	output := PalindromePartitioningMinCuts("noonabbad")
	require.Equal(t, 2, output)
}
