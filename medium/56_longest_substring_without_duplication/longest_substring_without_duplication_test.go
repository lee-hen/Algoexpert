package longest_substring_without_duplication

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := "mentisac"
	output := LongestSubstringWithoutDuplication("clementisacap")
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := "bdmentisac"
	output := LongestSubstringWithoutDuplication("clebdmentisacap")
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := "abc"
	output := LongestSubstringWithoutDuplication("abc")
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := "abcdef"
	output := LongestSubstringWithoutDuplication("abcdeabcdefc")
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := "dabcef"
	output := LongestSubstringWithoutDuplication("abcdabcef")
	require.Equal(t, expected, output)
}

func TestCase6(t *testing.T) {
	expected := "cbde"
	output := LongestSubstringWithoutDuplication("abcbde")
	require.Equal(t, expected, output)
}


