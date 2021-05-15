package four_number_sum

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func doTest(t *testing.T, expected, output [][]int) {
	t.Helper()
	require.Len(t, output, len(expected))
	for _, quad := range expected {
		sort.Ints(quad)
		ourquad := fmt.Sprintf("%v", quad)
		found := false
		for _, theirquad := range output {
			sort.Ints(theirquad)
			if fmt.Sprintf("%v", theirquad) == ourquad {
				found = true
				break
			}
		}
		require.True(t, found)
	}
}

func TestCase1(t *testing.T) {
	expected := [][]int{{7, 6, 4, -1}, {7, 6, 1, 2}}
	output := FourNumberSum([]int{7, 6, 4, -1, 1, 2}, 16)
	doTest(t, expected, output)
}
