package words_in_phone_number

import (
	"fmt"
	"strings"
)


var LetterDigits = map[rune]rune{
	'a': '2',
	'b': '2',
	'c': '2',
	'd': '3',
	'e': '3',
	'f': '3',
	'g': '4',
	'h': '4',
	'i': '4',
	'j': '5',
	'k': '5',
	'l': '5',
	'm': '6',
	'n': '6',
	'o': '6',
	'p': '7',
	'q': '7',
	'r': '7',
	's': '7',
	't': '8',
	'u': '8',
	'v': '8',
	'w': '9',
	'x': '9',
	'y': '9',
	'z': '9',
}

// WordsInPhoneNumber
// O(n^2 + m * w) time | O(n^2 + m * w) space - where n is the length of the
// phone number, m is the number of words, and w is the length of the longest word
func WordsInPhoneNumber(phoneNumber string, words []string) []string {
	phoneNumberSuffixTrie := NewModifiedSuffixTrie()
	// build Suffix Trie
	phoneNumberSuffixTrie.PopulateModifiedSuffixTrieFrom(phoneNumber)
	output := make([]string, 0)
	for _, word := range words {
		digitWord := wordToDigits(word)
		if phoneNumberSuffixTrie.Contains(digitWord) {
			output = append(output, word)
		}
	}
	return output
}

func wordToDigits(word string) string {
	builder := strings.Builder{}
	for i := 0; i < len(word); i++ {
		char := LetterDigits[rune(word[i])]
		fmt.Fprintf(&builder, string(char))
	}
	return builder.String()
}

type ModifiedSuffixTrie map[byte]ModifiedSuffixTrie

func NewModifiedSuffixTrie() ModifiedSuffixTrie {
	trie := ModifiedSuffixTrie{}
	return trie
}

// PopulateModifiedSuffixTrieFrom
// O(n^2) time | O(n^2) space
func (trie ModifiedSuffixTrie) PopulateModifiedSuffixTrieFrom(str string) {
	for i := range str {
		node := trie
		for j := i; j < len(str); j++ {
			letter := str[j]
			if _, found := node[letter]; !found {
				node[letter] = NewModifiedSuffixTrie()
			}
			node = node[letter]
		}
		node['*'] = nil
	}
}

// Contains
// O(m) time | O(1) space
func (trie ModifiedSuffixTrie) Contains(str string) bool {
	node := trie
	for i := 0; i < len(str); i++ {
		letter := str[i]
		if _, found := node[letter]; !found {
			return false
		}
		node = node[letter]
	}
	return true
}
