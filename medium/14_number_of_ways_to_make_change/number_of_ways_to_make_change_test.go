package number_of_ways_to_make_change

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	output := NumberOfWaysToMakeChange(6, []int{1, 2, 3})
	require.Equal(t, 7, output)

	output = NumberOfWaysToMakeChange(6, []int{1, 2, 3, 4})
	require.Equal(t, 9, output)

	output = NumberOfWaysToMakeChange(10, []int{1, 5, 10, 25})
	require.Equal(t, 4, output)

	output = NumberOfWaysToMakeChange(25, []int{1, 5, 10, 25})
	require.Equal(t, 13, output)
}
