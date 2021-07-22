package knuth_morris_pratt_algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := true
	output := knuthMorrisPrattAlgorithm("aefaefaefaedaefaedaefaefa", "aefaedaefaefa")
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := true
	output := KnuthMorrisPrattAlgorithm("aefoaefcdaefcdaed", "aefcdaed")
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := true
	output := KnuthMorrisPrattAlgorithm("aabc", "abc")
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := true
	output := KnuthMorrisPrattAlgorithm("aefcdfaecdaefaefcdaefeaefcdcdeae", "aefcdaefeaefcd")
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := false
	output := KnuthMorrisPrattAlgorithm("decadaafcdf", "daf")
	require.Equal(t, expected, output)
}

func TestCase6(t *testing.T) {
	expected := false
	output := KnuthMorrisPrattAlgorithm("bccbefbcdabbbcabfdcfe", "abc")
	require.Equal(t, expected, output)
}

