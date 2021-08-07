package build_failures

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	buildRuns := [][]bool{
		{true, true, true, false, false},
		{true, true, true, true, false},
		{true, true, true, true, true, true, false, false, false},
		{true, false, false, false, false, false},
		{true, true, true, true, true, true, true, true, true, true, true, true, false},
		{true, false},
		{true, true, true, true, false, false},
	}
	actual := BuildFailures(buildRuns)
	require.Equal(t, 3, actual)
}

func TestCase2(t *testing.T) {
	buildRuns := [][]bool{
		{true, false, true, true, true, true, true, true, true, true, true, true, false},
		{true, false},
		{true, false, false, false, false, false},
	}
	actual := BuildFailures(buildRuns)
	require.Equal(t, 3, actual)
}
