package longest_palindromic_substring


import (
"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	expected := "xyzzyx"
	output := LongestPalindromicSubstring("abaxyzzyxf")
	require.Equal(t, expected, output)
}
