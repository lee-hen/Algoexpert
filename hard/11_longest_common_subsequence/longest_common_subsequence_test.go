package longest_common_subsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []byte("XYZW")
	output := LongestCommonSubsequence("ZXVVYZW", "XKYKZPW")
	require.Equal(t, expected, output)
}
