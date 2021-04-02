package findthreelargestnumbers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{-2, -1, 7}
	output := FindThreeLargestNumbers([]int{-1, -2, -3, -7, -17, -27, -18, -541, -8, -7, 7})
	require.Equal(t, expected, output)
}
