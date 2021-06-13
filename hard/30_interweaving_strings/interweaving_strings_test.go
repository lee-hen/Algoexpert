package interweaving_strings

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func TestCase1(t *testing.T) {
	one := "algoexpert"
	two := "your-dream-job"
	three := "your-algodream-expertjob"
	require.Equal(t, true, InterweavingStrings(one, two, three))
}

func TestCase2(t *testing.T) {
	one := "abc"
	two := "123"
	three := "a1b2c3"
	require.Equal(t, true, InterweavingStrings(one, two, three))
}

func TestCase3(t *testing.T) {
	one := "a"
	two := "b"
	three := "ab"
	require.Equal(t, true, InterweavingStrings(one, two, three))
}

func TestCase4(t *testing.T) {
	one := "a"
	two := "b"
	three := "ac"
	require.Equal(t, false, InterweavingStrings(one, two, three))
}

func TestCase5(t *testing.T) {
	one := "aacaaaa"
	two := "aaabaaa"
	three := "aaaabacaaaaaaa"
	require.Equal(t, true, InterweavingStrings(one, two, three))
}

func TestCase6(t *testing.T) {
	one := "aacaaaa"
	two := "aaabaaa"
	three := "aaaacabaaaaaaa"
	require.Equal(t, true, InterweavingStrings(one, two, three))
}

func TestCase7(t *testing.T) {
	one := "algo"
	two := "frog"
	three := "fralgogo"
	require.Equal(t, true, interweavingStrings(one, two, three))
}

func TestCase8(t *testing.T) {
	one := "aaaaaaaaaaaa"
	two := "aaaaaaaaaac"
	three := "aaaaaaaaaaaaaaaaaaaaaaa"
	require.Equal(t, false, interweavingStrings(one, two, three))
}
