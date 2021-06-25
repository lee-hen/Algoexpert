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
