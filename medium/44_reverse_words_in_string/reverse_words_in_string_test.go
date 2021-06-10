package reverse_words_in_string

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := "Algoexpert is the best!"
	expected := "best! the is Algoexpert"
	actual := ReverseWordsInString(input)
	require.Equal(t, expected, actual)
}
