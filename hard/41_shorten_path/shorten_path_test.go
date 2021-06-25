package shorten_path

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := "/foo/bar/baz"
	output := ShortenPath("/foo/../test/../test/../foo//bar/./baz")
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := "/foo/bar/baz"
	output := ShortenPath("/../../foo/bar/baz")
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := "/bar/baz"
	output := ShortenPath("/../../foo/../../bar/baz")
	require.Equal(t, expected, output)
}

func TestCase4(t *testing.T) {
	expected := "/foo/kappa"
	output := ShortenPath("/foo/./././bar/./baz///////////test/../../../kappa")
	require.Equal(t, expected, output)
}

func TestCase5(t *testing.T) {
	expected := "foo"
	output := ShortenPath("foo/bar/..")
	require.Equal(t, expected, output)
}

func TestCase6(t *testing.T) {
	expected := "/"
	output := ShortenPath("/")
	require.Equal(t, expected, output)
}

func TestCase7(t *testing.T) {
	expected := "foo/bar"
	output := ShortenPath("./foo/bar")
	require.Equal(t, expected, output)
}

func TestCase8(t *testing.T) {
	expected := ".."
	output := ShortenPath("./..")
	require.Equal(t, expected, output)
}

func TestCase9(t *testing.T) {
	expected := ".."
	output := ShortenPath("foo/../..")
	require.Equal(t, expected, output)
}

func TestCase10(t *testing.T) {
	expected := "/"
	output := ShortenPath("/../../../this////one/./../../is/../../going/../../to/be/./././../../../just/a/forward/slash/../../../../../..")
	require.Equal(t, expected, output)
}

func TestCase11(t *testing.T) {
	expected := ".."
	output := ShortenPath("../")
	require.Equal(t, expected, output)
}

func TestCase12(t *testing.T) {
	expected := ""
	output := ShortenPath("./")
	require.Equal(t, expected, output)
}

func TestCase13(t *testing.T) {
	expected := "../../../bar/baz"
	output := shortenPath("../../foo/../../bar/baz")
	require.Equal(t, expected, output)
}

// ../../foo/../../bar/baz
// ../../../../