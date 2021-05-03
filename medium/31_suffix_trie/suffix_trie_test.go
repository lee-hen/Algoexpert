package suffix_trie

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func (trie SuffixTrie) Equals(other SuffixTrie) bool {
	if len(trie) != len(other) {
		return false
	}
	for key, child := range trie {
		otherChild, found := other[key]
		if !found {
			return false
		} else if child != nil && !child.Equals(otherChild) {
			return false
		}
	}
	return true
}

func TrieFromString(str string) SuffixTrie {
	trie := SuffixTrie{}
	trie.PopulateSuffixTrieFrom(str)
	return trie
}

func TestCase1(t *testing.T) {
	trie := TrieFromString("babc")
	expected := SuffixTrie{
		'c': {'*': nil},
		'b': {
			'c': {'*': nil},
			'a': {'b': {'c': {'*': nil}}},
		},
		'a': {'b': {'c': {'*': nil}}},
	}
	require.True(t, trie.Equals(expected))
	require.True(t, trie.Contains("abc"))
}
