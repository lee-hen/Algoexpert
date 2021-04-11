package is_monotonic

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	array := []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001}
	actual := IsMonotonic(array)
	require.True(t, actual)
}
