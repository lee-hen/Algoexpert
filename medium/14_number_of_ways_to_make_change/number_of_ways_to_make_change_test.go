package number_of_ways_to_make_change

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	output := NumberOfWaysToMakeChange(6, []int{1, 5})
	require.Equal(t, 2, output)
}
