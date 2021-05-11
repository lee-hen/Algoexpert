package reverse_words_in_string

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "AlgoExpert is the best!"
	expected := "best! the is AlgoExpert"
	actual := ReverseWordsInString(input)
	require.Equal(t, expected, actual)
}
