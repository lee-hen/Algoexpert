package underscorify_substring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := "_test_this is a _testtest_ to see if _testestest_ it works"
	output := UnderscorifySubstring("testthis is a testtest to see if testestest it works", "test")
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := "_test_this is a _testtest_ t__t t_es_t to see if _testest_t_es__test_ it works"
	output := UnderscorifySubstring("testthis is a testtest t__t t_es_t to see if testestt_es_test it works", "test")
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := "_testest_t_test_ tt"
	output := UnderscorifySubstring("testestttest tt", "test")
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := "_a_b_a_b_a_b_a_b_a_b_a_b_a_b_a_b_a_b_a_b_a_b_a_b_a_b_aa_b_a_b_aaa_bb_a_b_a_b_aa_"
	output := UnderscorifySubstring("abababababababababababababaababaaabbababaa", "a")
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := "b_aaa_b"
	output := UnderscorifySubstring("baaab", "a")
	require.Equal(t, expected, output)
}

func TestCase6(t *testing.T) {
	expected := "_abcabcabcabcabcabcabcabcabcabcabcabcabcabc_"
	output := UnderscorifySubstring("abcabcabcabcabcabcabcabcabcabcabcabcabcabc", "abc")
	require.Equal(t, expected, output)
}

func TestCase7(t *testing.T) {
	expected := "_abcabcabcabcabc_ a"
	output := UnderscorifySubstring("abcabcabcabcabc a", "abc")
	require.Equal(t, expected, output)
}
