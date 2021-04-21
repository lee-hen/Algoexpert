package single_cycle_check

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{2, 3, 1, -4, -4, 2}
	output := HasSingleCycle(input)
	expected := true
	require.Equal(t, expected, output)
}
