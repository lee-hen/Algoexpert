package suffix_trie

type SuffixTrie map[byte]SuffixTrie

func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

// PopulateSuffixTrieFrom
// better solution
func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	for i := range str {
		node := trie
		for j := i; j < len(str); j++ {
			letter := str[j]
			if _, found := node[letter]; !found {
				node[letter] = NewSuffixTrie()
			}
			node = node[letter]
		}
		node['*'] = nil
	}
}

// Contains
// better solution
func (trie SuffixTrie) Contains(str string) bool {
	node := trie
	for i := 0; i < len(str); i++ {
		letter := str[i]
		if _, found := node[letter]; !found {
			return false
		}
		node = node[letter]
	}
	_, found := node['*']
	return found
}

// babc
//      root
//   b    a    c
// c  a   b
//    b   c
//    c
// my solution
func (trie SuffixTrie) populateSuffixTrieFrom(str string) {
	chars := []rune(str)

	current := trie
	for i := 0; i < len(chars); i++ {
		currentChar := byte(chars[i])

		if _, found := trie[currentChar]; !found {
			trie[currentChar] = NewSuffixTrie()
			current = trie[currentChar]
		} else {
			current = current[currentChar]
		}

		for j := i+1; j < len([]rune(str)); j++ {
			nextChar := byte(chars[j])
			if _, found := current[nextChar]; !found {
				current[nextChar] = NewSuffixTrie()
			}
			current = current[nextChar]

		}
		current[byte('*')] = nil
		current = trie
	}
}

// my solution
func (trie SuffixTrie) contains(str string) bool {
	current := trie

	for _, r := range str {
		char := byte(r)
		if _, found := current[char]; found {
			current = current[char]
		}
	}

	if _, ok := current[byte('*')]; ok {
		return true
	}

	return false
}
