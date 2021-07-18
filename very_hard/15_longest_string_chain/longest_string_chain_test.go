package longest_string_chain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := []string{"abde", "abc", "abd", "abcde", "ade", "ae", "1abde", "abcdef"}
	expected := []string{"abcdef", "abcde", "abde", "ade", "ae"}
	output := LongestStringChain(input)
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	input := []string{"abcdefg1", "1234c", "abdefg2", "abdfg", "123", "122", "bgg", "g", "1a2345", "12a345"}
	expected := []string{}
	output := LongestStringChain(input)
	require.Equal(t, expected, output)
}
