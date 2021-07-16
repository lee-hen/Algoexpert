package longest_balanced_substring

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := "(()))("
	expected := 4
	actual := LongestBalancedSubstring(input)
	require.Equal(t, expected, actual)
}
