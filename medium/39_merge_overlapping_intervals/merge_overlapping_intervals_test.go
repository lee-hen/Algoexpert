package merge_overlapping_intervals

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := [][]int{
		{1, 22},
		{-20, 30},
	}
	expected := [][]int{
		{-20, 30},
	}
	actual := MergeOverlappingIntervals(input)
	require.Equal(t, expected, actual)
}

//[9, 12],
//[12, 54],
//[43, 49],
//[45, 90],
//[91, 93]