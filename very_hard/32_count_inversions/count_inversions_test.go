package count_inversions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := []int{2, 3, 3, 1, 9, 5, 6}
	expected := 5
	actual := CountInversions(input)
	require.Equal(t, expected, actual)
}
