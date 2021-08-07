package words_in_phone_number

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	phoneNumber := "3662277"
	words := []string{"foo", "bar", "baz", "foobar", "emo", "cap", "car", "cat"}
	expected := []string{"bar", "cap", "car", "emo", "foo", "foobar"}
	actual := wordsInPhoneNumber(phoneNumber, words)
	require.ElementsMatch(t, expected, actual)
}

func TestCase2(t *testing.T) {
	phoneNumber := "36612277"
	words := []string{"foobar", "foo", "bar", "baz",  "emo", "cap", "car", "cat"}
	expected := []string{"bar", "cap", "car", "emo", "foo"}
	actual := WordsInPhoneNumber(phoneNumber, words)
	require.ElementsMatch(t, expected, actual)
}
