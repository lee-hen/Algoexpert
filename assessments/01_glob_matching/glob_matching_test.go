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

func TestCase2(t *testing.T) {
	fileName := "abcdagvvvvag"
	pattern := "*?ag"
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}

func TestCase4(t *testing.T) {
	fileName := "abcdagvvgaavg"
	pattern := "*?aa?g"
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}

func TestCase5(t *testing.T) {
	fileName := "abcdefg"
	pattern := "abcddefg"
	output := GlobMatching(fileName, pattern)
	require.False(t, output)
}

func TestCase6(t *testing.T) {
	fileName := ""
	pattern := ""
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}

func TestCase7(t *testing.T) {
	fileName := ""
	pattern := "?"
	output := GlobMatching(fileName, pattern)
	require.False(t, output)
}

func TestCase8(t *testing.T) {
	fileName := "abcdefg"
	pattern := "*c*cdefg"
	output := GlobMatching(fileName, pattern)
	require.False(t, output)
}

func TestCase3(t *testing.T) {
	fileName := "abcdagvvvvavg"
	pattern := "*?ag"
	output := GlobMatching(fileName, pattern)
	require.False(t, output)
}

func TestCase9(t *testing.T) {
	fileName := "abcdefg"
	pattern := "abcd*efg"
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}

func TestCase10(t *testing.T) {
	fileName := "abcdefg"
	pattern := "?"
	output := globMatching(fileName, pattern)
	require.False(t, output)
}

func TestCase11(t *testing.T) {
	fileName := "abcdefg"
	pattern := "*****a*?f*********g"
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}

func TestCase12(t *testing.T) {
	fileName := ""
	pattern := "***"
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}
func TestCase13(t *testing.T) {
	fileName := "abcdefg"
	pattern := ""
	output := GlobMatching(fileName, pattern)
	require.False(t, output)
}

func TestCase14(t *testing.T) {
	fileName := ""
	pattern := ""
	output := GlobMatching(fileName, pattern)
	require.True(t, output)
}
