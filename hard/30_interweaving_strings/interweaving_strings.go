package interweaving_strings
// important question

// InterweavingStrings
// O(nm) time | O(nm) space - where n is the length of the
// first string and m is the length of the second string
func InterweavingStrings(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}
	cache := make([][]*bool, len(one)+1)

	for i := 0; i < len(one)+1; i++ {
		cache[i] = make([]*bool, len(two)+1)
	}

	return areInterwoven(one, two, three, 0, 0, cache)
}

func areInterwoven(one, two, three string, i, j int, cache [][]*bool) bool {
	if cache[i][j] != nil {
		return *cache[i][j]
	}

	k := i + j
	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		result := areInterwoven(one, two, three, i+1, j, cache)
		cache[i][j] = &result
		if result {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		result := areInterwoven(one, two, three, i, j+1, cache)
		cache[i][j] = &result
		if result {
			return true
		}
	}

	result := false
	cache[i][j] = &result
	return result
}

// InterweavingStrings1
// O(2^(n + m)) time | O(n + m) space - where n is the length
// of the first string and m is the length of the second string
func InterweavingStrings1(one, two, three string) bool {
	if len(three) != len(one)+len(two) {
		return false
	}
	return areInterwoven1(one, two, three, 0, 0)
}

func areInterwoven1(one, two, three string, i, j int) bool {
	k := i + j
	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		if areInterwoven1(one, two, three, i+1, j) {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		if areInterwoven1(one, two, three, i, j+1) {
			return true
		}
	}
	return false
}

// a a c a a a a
// a a a b a a a
// a a a a b a c a a a a a a a
// true

// a a c a a a a
// a a a b a a a
// a a a a c a b a a a a a a a
// true

// a l g o
// f r o g
// f r a l g o g o
// true

// o
// o g
// o g o
// true
// i think this one is better because it's time complexity is O(2(n+m))
func interweavingStrings(one, two, three string) bool {
	if len(one)+len(two) != len(three) {
		return false
	}

	trie1 := Trie{children: map[byte]Trie{}}
	trie1.Add(one)

	trie2 := Trie{children: map[byte]Trie{}}
	trie2.Add(two)

	cache := make(map[int]*bool)
	return explore(trie1, trie2, cache, 0, []byte(three))
}

func explore(ptr1, ptr2 Trie, cache map[int]*bool, i int, target []byte) bool {
	if cache[i] != nil {
		return *cache[i]
	}

	_, isEndOfOne := ptr1.children['*']
	_, isEndOfTwo := ptr2.children['*']

	if isEndOfOne && isEndOfTwo {
		return true
	}

	letter := target[i]
	trie1, found1 := ptr1.children[letter]
	trie2, found2 := ptr2.children[letter]

	if !found1 && !found2 {
		return false
	}

	if found1 && found2 {
		var found bool
		if _, found = trie1.children[target[i+1]]; found {
			ptr1 = trie1
		} else if _, found = trie2.children[target[i+1]]; found {
			ptr2 = trie2
		}

		if found {
			result := explore(ptr1, ptr2, cache, i+1, target)
			cache[i] = &result

			return result
		}
	}

	if found1 {
		ptr1 = trie1
	} else if found2 {
		ptr2 = trie2
	}

	return explore(ptr1, ptr2, cache, i+1, target)
}

type Trie struct {
	children map[byte]Trie
}

func (t Trie) Add(word string) {
	current := t

	for _, letter := range []byte(word) {
		if _, found := current.children[letter]; !found {
			current.children[letter] = Trie{
				children: map[byte]Trie{},
			}
		}

		current = current.children[letter]
	}

	current.children['*'] = Trie{
		children: map[byte]Trie{},
	}
}
