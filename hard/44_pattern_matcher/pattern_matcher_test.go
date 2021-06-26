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
