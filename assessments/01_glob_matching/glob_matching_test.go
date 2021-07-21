package glob_matching

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	fileName := "abcdefg"
	pattern := "a*e?g"
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}
