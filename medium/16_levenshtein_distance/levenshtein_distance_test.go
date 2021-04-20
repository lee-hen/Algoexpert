package levenshtein_distance


import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	require.Equal(t, LevenshteinDistance("abc", "yabd"), 2)
}
