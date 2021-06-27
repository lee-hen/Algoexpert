package multi_string_search

// MultiStringSearch
// O(ns + bs) time | O(ns) space
func MultiStringSearch(bigString string, smallStrings []string) []bool {
	trie := Trie{children: map[byte]Trie{}}
	for i, str := range smallStrings {
		trie.Add(str, i)
	}

	containedStrings := map[string]bool{}
	for i := range bigString {
		findSmallStringsIn(bigString, i, trie, containedStrings)
	}

	output := make([]bool, len(smallStrings))
	for i, str := range smallStrings {
		output[i] = containedStrings[str]
	}
	return output
}

func findSmallStringsIn(str string, startIdx int, trie Trie, containedStrings map[string]bool) {
	current := trie
	for i := startIdx; i < len(str); i++ {
		currentChar := str[i]
		if _, found := current.children[currentChar]; !found {
			break
		}
		current = current.children[currentChar]
		if end, found := current.children['*']; found {
			containedStrings[end.word] = true
		}
	}
}

// 0 1 2 3 4 5 6 7 8 9 10  11 12 13 14 15 16 17
// a b c d e f g h i j  k  l  m  n  o  p  q   r s t u v w x y z
// "abc", "m n o p q r", "wyz", "no", "e", "t u u v"
// [true, true,           false, true, true, false]
//  true, true,           false, false, true, false

// "this is a big string"
// "this", "yo",  "is",  "a",  "bigger", "string", "kappa"
//    0      1      2     3       4         5         6
//  true,  false, true,  true,  false,    true,     false


// 0 1 2 3 4 5 6 7 8 9 10  11  12  13  14  15  16  17  18 19  20  21
// h j ! ) ! % H j 1 j h   8    f   1   9   8  5    n  !   )   5   1
// my solution
func multiStringSearch(bigString string, smallStrings []string) []bool {
	result := make([]bool, len(smallStrings), len(smallStrings))

	trie := NewTrie()
	trie.AddSlice(smallStrings)

	bigStringByte := []byte(bigString)

	for i := 0; i < len(bigStringByte); i++ {
		char := bigStringByte[i]
		if current, found := trie.children[char]; found {

			if endNode, isEnd := current.children['*']; isEnd {
				result[endNode.idx] = true
			}

			for j := i+1; j < len(bigStringByte); j++ {
				if next, found := current.children[bigStringByte[j]]; found {
					current = next
					if endNode, isEnd := current.children['*']; isEnd {
						result[endNode.idx] = true
					}
				} else {
					break
				}
			}
		}
	}

	return result
}

type Trie struct {
	children map[byte]Trie

	word string
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
		idx:     i,
		word:     str,
	}
}

// MultiStringSearch1
// O(b^2 + ns) time | O(b^2 + n) space
func MultiStringSearch1(bigString string, smallStrings []string) []bool {
	trie := NewSuffixTrie(bigString)
	output := make([]bool, len(smallStrings))

	for i, smallString := range smallStrings {
		output[i] = trie.Contains(smallString)
	}
	return output
}

type ModifiedSuffixTrie map[byte]ModifiedSuffixTrie

func NewSuffixTrie(str string) ModifiedSuffixTrie {
	trie := ModifiedSuffixTrie{}

	for i := range str {
		trie.Add(str, i)
	}
	return trie
}

func (trie ModifiedSuffixTrie) Add(str string, startIndex int) {
	node := trie
	for j := startIndex; j < len(str); j++ {
		letter := str[j]
		if _, found := node[letter]; !found {
			node[letter] = ModifiedSuffixTrie{}
		}
		node = node[letter]
	}
}

func (trie ModifiedSuffixTrie) Contains(str string) bool {
	node := trie

	for i := range str {
		letter := str[i]
		if _, found := node[letter]; !found {
			return false
		}
		node = node[letter]
	}
	return true
}
