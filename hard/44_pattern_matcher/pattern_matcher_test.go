package pattern_matcher

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []string{"go", "powerranger"}
	output := PatternMatcher("xxyxxy", "gogopowerrangergogopowerranger")
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []string{"g", "ogg"}
	output := PatternMatcher("xyxx", "gogggg")
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := []string{"b", "a"}
	output := PatternMatcher("yxyx", "abab")
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := []string{"test", ""}
	output := PatternMatcher("xxxx", "testtesttesttest")
	require.Equal(t, expected, output)
}
