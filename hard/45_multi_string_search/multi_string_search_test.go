package multi_string_search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []bool{true, false, true, true, false, true, false}
	output := MultiStringSearch("this is a big string", []string{"this", "yo", "is", "a", "bigger", "string", "kappa"})
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []bool{true, true, false, true, true, false}
	output := MultiStringSearch("abcdefghijklmnopqrstuvwxyz", []string{"abc", "mnopqr", "wyz", "no", "e", "tuuv"})
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := []bool{true, true, true, true}
	output := MultiStringSearch("abcdefghijklmnopqrstuvwxyz", []string{"abcdefghijklmnopqrstuvwxyz", "abc", "j", "mnopqr"})
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := []bool{true}
	output := MultiStringSearch("bbbabb", []string{"bbabb"})
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := []bool{false, true, true, true, false, false, false}
	output := MultiStringSearch("hj!)!%Hj1jh8f1985n!)51", []string{"%Hj7", "8f198", "!)5", "!)!", "!!", "jh81", "j181hf"})
	require.Equal(t, expected, output)
}
