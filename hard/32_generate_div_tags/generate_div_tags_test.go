package generate_div_tags

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := 3
	expected := []string{
		"<div><div><div></div></div></div>",
		"<div><div></div><div></div></div>",
		"<div><div></div></div><div></div>",
		"<div></div><div><div></div></div>",
		"<div></div><div></div><div></div>",
	}
	actual := GenerateDivTags(input)
	require.Equal(t, expected, actual)
}
