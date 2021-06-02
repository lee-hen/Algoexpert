package longest_common_subsequence

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []byte("XYZW")
	output := longestCommonSubsequence("ZXVVYZW", "XKYKZPW")
	require.Equal(t, expected, output)
}


func TestCase2(t *testing.T) {
	expected := []byte("CDEGHJKLTUV")
	output := longestCommonSubsequence("CCCDDEGDHAGKGLWAJWKJAWGKGWJAKLGGWAFWLFFWAGJWKAGTUV", "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	require.Equal(t, expected, output)
}
