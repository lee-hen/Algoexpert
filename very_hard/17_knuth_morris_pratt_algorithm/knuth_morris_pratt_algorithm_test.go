package knuth_morris_pratt_algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := true
	output := KnuthMorrisPrattAlgorithm("aefoaefcdaefcdaed", "aefcdaed")
	require.Equal(t, expected, output)
}
