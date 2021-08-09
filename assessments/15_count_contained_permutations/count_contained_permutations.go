package count_contained_permutations

// CountContainedPermutations
// O(b + s) time | O(s) space - where b is the length of the big
// input string and s is the length of the small input string
func CountContainedPermutations(bigString, smallString string) int {
	smallStringCharCounts := getCharCounts(smallString)
	numUniqueChars := len(smallStringCharCounts)
	runningCharCounts := map[rune]int{}
	permutationsCount := 0
	numUniqueCharsDone := 0
	leftIdx := 0
	rightIdx := 0

	for rightIdx < len(bigString) {
		rightChar := rune(bigString[rightIdx])
		if _, found := smallStringCharCounts[rightChar]; found {
			increaseCharCount(rightChar, runningCharCounts)

			if runningCharCounts[rightChar] == smallStringCharCounts[rightChar] {
				numUniqueCharsDone++
			}
		}

		if numUniqueCharsDone == numUniqueChars {
			permutationsCount++
		}

		rightIdx++

		shouldIncrementLeftIdx := rightIdx-leftIdx == len(smallString)
		if !shouldIncrementLeftIdx {
			continue
		}

		leftChar := rune(bigString[leftIdx])
		if _, found := smallStringCharCounts[leftChar]; found {
			if runningCharCounts[leftChar] == smallStringCharCounts[leftChar] {
				numUniqueCharsDone--
			}

			decreaseCharCount(leftChar, runningCharCounts)
		}

		leftIdx++
	}
	return permutationsCount
}

func getCharCounts(str string) map[rune]int {
	charCounts := map[rune]int{}
	for _, char := range str {
		increaseCharCount(char, charCounts)
	}
	return charCounts
}

func increaseCharCount(char rune, charCounts map[rune]int) {
	charCounts[char]++
}

func decreaseCharCount(char rune, charCounts map[rune]int) {
	charCounts[char]--
}

// my solution

func countContainedPermutations(bigString, smallString string) int {
	smallStrings := getPermutations([]byte(smallString))
	trie := NewTrie()
	trie.AddSlice(smallStrings)

	var counter int
	for i := range bigString {
		findSmallStringsIn(bigString, i, trie, &counter)
	}

	return counter
}

func findSmallStringsIn(str string, startIdx int, trie Trie, counter *int) {
	current := trie
	for i := startIdx; i < len(str); i++ {
		currentChar := str[i]
		if _, found := current.children[currentChar]; !found {
			break
		}
		current = current.children[currentChar]
		if _, found := current.children['*']; found {
			*counter++
		}
	}
}

type Trie struct {
	children map[byte]Trie

}

func NewTrie() Trie {
	trie := Trie{children: map[byte]Trie{}}
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
	}
}

func getPermutations(array []byte) []string {
	permutations := make([]string, 0)
	permutationsHelper(0, array, &permutations)
	return permutations
}

func permutationsHelper(i int, array []byte, permutations *[]string) {
	if i == len(array)-1 {
		newPerm := make([]byte, len(array))
		copy(newPerm, array)
		*permutations = append(*permutations, string(newPerm))
		return
	}

	for j := i; j < len(array); j++ {
		swap(array, i, j)
		permutationsHelper(i+1, array, permutations)
		swap(array, i, j)
	}
}

func swap(array []byte, i, j int) {
	array[i], array[j] = array[j], array[i]
}

