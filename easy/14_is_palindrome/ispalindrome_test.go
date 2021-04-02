package ispalindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := true
	output := IsPalindrome("abcdcba")

	require.Equal(t, expected, output)
}
