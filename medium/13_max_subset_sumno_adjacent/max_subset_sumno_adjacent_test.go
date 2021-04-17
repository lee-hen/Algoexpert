package max_subset_sumno_adjacent

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	res := MaxSubsetSumNoAdjacent([]int{75, 105, 120, 75, 90, 135})
	require.Equal(t, res, 330)
}
