package smallest_substring_containing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	bigString := "abcd$ef$axb$c$"
	smallString := "$$abf"
	expected := "f$axb$"
	require.Equal(t,
		SmallestSubstringContaining(bigString, smallString), expected,
	)
}
