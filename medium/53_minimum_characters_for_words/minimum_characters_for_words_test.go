package minimum_characters_for_words

import (
	"github.com/stretchr/testify/require"

	"sort"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []string{"this", "that", "did", "deed", "them!", "a"}
	expected := []string{"t", "t", "h", "i", "s", "a", "d", "d", "e", "e", "m", "!"}
	actual := MinimumCharactersForWords(input)
	sort.Strings(actual)
	sort.Strings(expected)
	require.Equal(t, expected, actual)
}
