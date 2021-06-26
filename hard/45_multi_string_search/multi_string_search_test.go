package multi_string_search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []bool{true, false, true, true, false, true, false}
	output := MultiStringSearch("this is a big string", []string{"this", "yo", "is", "a", "bigger", "string", "kappa"})
	require.Equal(t, expected, output)
}
