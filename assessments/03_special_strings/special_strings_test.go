package special_strings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	strings := []string{
		"foobarbaz",
		"foo",
		"bar",
		"foobarfoo",
		"baz",
		"foobaz",
		"foofoofoo",
		"foobazar",
	}
	actual := SpecialStrings(strings)
	expected := []string{"foobarbaz", "foobarfoo", "foobaz", "foofoofoo"}
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	strings := []string{
		"foobarbaz", "foo", "baz", "foobar",
	}
	actual := SpecialStrings(strings)
	expected := []string{"foobarbaz"}
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	strings := []string{
		"foobarbaz", "testfoobar", "test", "foo", "baz",
	}
	actual := SpecialStrings(strings)
	expected := []string{}
	require.Equal(t, expected, actual)
}
