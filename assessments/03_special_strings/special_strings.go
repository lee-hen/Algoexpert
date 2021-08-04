package special_strings

// SpecialStrings
// Average case: when there aren't too many strings with
// identical prefixes that are found in another string
// O(n * m) time | O(n * m) space - where n is the number
// of strings and m is the length of the longest string
// --------
// See the Notes section in the Explanation tab for more info.
func SpecialStrings(strings []string) []string {
	trie := newTrie()

	for _, str := range strings {
		trie.insert(str)
	}

	output := make([]string, 0)
	for _, str := range strings {
		if isSpecial(str, trie.root, 0, 0, trie) {
			output = append(output, str)
		}
	}
	return output
}

func isSpecial(str string, node trieNode, idx int, count int, t *trie) bool {
	char := rune(str[idx])
	nextTrieNode, found := node[char]
	if !found {
		return false
	}

	atEndOfString := idx == len(str)-1
	_, isComplete := nextTrieNode[t.endSymbol]
	if atEndOfString {
		return isComplete && count+1 >= 2
	}

	if isComplete {
		restIsSpecial := isSpecial(str, t.root, idx+1, count+1, t)
		if restIsSpecial {
			return true
		}
	}

	return isSpecial(str, nextTrieNode, idx+1, count, t)
}

type trie struct {
	endSymbol rune
	root      trieNode
}

type trieNode map[rune]trieNode

func newTrie() *trie {
	return &trie{
		endSymbol: '*',
		root:      trieNode{},
	}
}

func (t *trie) insert(str string) {
	currentTrieNode := t.root
	for i := 0; i < len(str); i++ {
		char := str[i]
		child, found := currentTrieNode[rune(char)]
		if !found {
			child = trieNode{}
			currentTrieNode[rune(char)] = child
		}
		currentTrieNode = child
	}
	currentTrieNode[t.endSymbol] = trieNode{}
}


// my solution

type Trie struct {
	children map[byte]Trie

	idx int
}

func NewTrie() Trie {
	trie := Trie{children: map[byte]Trie{}, idx: -1}
	return trie
}

func (t Trie) AddSlice(slice []string) {
	for i := range slice {
		str := slice[i]
		t.Add(str, i)
	}
}

func (t Trie) Add(str string, i int) {
	node := t

	for j := 0; j < len(str); j++ {
		letter := str[j]
		if _, found := node.children[letter]; !found {
			node.children[letter] = NewTrie()
		}
		node = node.children[letter]
	}

	node.children['*'] = Trie{
		children: map[byte]Trie{},
		idx:	i,
	}
}

func specialStrings(strings []string) []string {
	trie := NewTrie()
	trie.AddSlice(strings)
	result := make([]string, 0)

	for currentIdx, str := range strings {
		var isSpecial, isFound bool

		current := &trie
		var previous *Trie
		for _, char := range []byte(str) {
			isSpecial = false

			if previous != nil {
				if next, found := previous.children[char]; found {
					previous = &next
					if endNode, isEnd := previous.children['*']; isEnd {
						if endNode.idx != currentIdx {
							isSpecial = true
							current = &trie
							isFound = true
						}
					}
				}
			}

			if next, found := current.children[char]; found {
				current = &next

				if endNode, isEnd := current.children['*']; isEnd {
					if endNode.idx != currentIdx {
						isSpecial = true
						previous = current
						current = &trie
					}
				}
			} else if !isFound {
				break
			}
		}

		if isSpecial  {
			result = append(result, str)
		}
	}

	return result
}


//strings = [
// "foobarbaz",
// "foo",
// "bar",
// "foobarfoo",
// "baz",
// "foobaz",
// "foofoofoo",
// "foobazar",
//]

// ["foobarbaz", "foobarfoo", "foobaz", "foofoofoo"]

// [foobarbaz foo bar foobarfoo baz foobaz foofoofoo foobazar]

// specialStrings
// naive solution
func specialStringsNaive(strings []string) []string {
	strIndices := make(map[string]int)

	for idx, str := range strings {
		strIndices[str] = idx
	}

	result := make([]string, 0)

	for currentIdx, str := range strings {
		var isSpecial bool
		var i int
		for j := range str {
			if idx, ok := strIndices[str[i:j+1]]; ok && currentIdx != idx {
				isSpecial = true
				i = j+1
			} else if idx1, ok1 := strIndices[str[0:j+1]]; ok1 && currentIdx != idx1 {
				isSpecial = true
				i = j+1
			}
		}

		if isSpecial && i == len(str) {
			result = append(result, str)
		}
	}

	return result
}
