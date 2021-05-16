package min_rewards

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{8, 4, 2, 1, 3, 6, 7, 9, 5}
	output := MinRewards(input)
	expected := 25
	require.Equal(t, expected, output)
}
