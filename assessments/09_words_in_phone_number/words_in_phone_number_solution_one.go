package words_in_phone_number

var DigitLetters = map[rune][]rune{
	'0': {},
	'1': {},
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

// O(m * w + n * 3^n) time | O(m * w + n) space - where n is the length of the
// phone number, m is the number of words, and w is the length of the longest word
func wordsInPhoneNumber1(phoneNumber string, words []string) []string {
	wordTrie := getWordTrie(words)
	finalWords := map[string]bool{}
	for i := 0; i < len(phoneNumber); i++ {
		exploreTrie(phoneNumber, i, wordTrie, finalWords)
	}

	result := make([]string, 0)

	for word := range finalWords {
		result = append(result, word)
	}
	return result
}

func getWordTrie(words []string) Trie {
	trie := Trie{children: map[rune]Trie{}}
	for _, word := range words {
		trie.Add(word)
	}
	return trie
}

func exploreTrie(phoneNumber string, digitIdx int, trieNode Trie, finalWords map[string]bool) {
	if child, found := trieNode.children['*']; found {
		finalWords[child.word] = true
	}

	if digitIdx == len(phoneNumber) {
		return
	}

	currentDigit := rune(phoneNumber[digitIdx])
	possibleLetters := DigitLetters[currentDigit]
	for _, letter := range possibleLetters {
		child, found := trieNode.children[letter]
		if !found {
			continue
		}
		exploreTrie(phoneNumber, digitIdx+1, child, finalWords)
	}
}

type Trie struct {
	children map[rune]Trie

	word string
}

func (t Trie) Add(word string) {
	current := t
	for _, letter := range word {
		if _, found := current.children[letter]; !found {
			current.children[letter] = Trie{
				children: map[rune]Trie{},
			}
		}
		current = current.children[letter]
	}
	current.children['*'] = Trie{
		children: map[rune]Trie{},
		word:     word,
	}
}
