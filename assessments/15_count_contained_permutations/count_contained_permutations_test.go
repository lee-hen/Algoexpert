package count_contained_permutations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	bigString := "cbabcacabca"
	smallString := "abc"
	expected := 6
	actual := CountContainedPermutations(bigString, smallString)
	require.Equal(t, expected, actual)
}
