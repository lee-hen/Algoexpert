package min_number_of_jumps

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 4
	output := MinNumberOfJumps([]int{3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3})
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := 4
	output := MinNumberOfJumps([]int{2, 1, 2, 2, 1, 1, 1})
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := 0
	output := MinNumberOfJumps([]int{1})
	require.Equal(t, expected, output)
}
