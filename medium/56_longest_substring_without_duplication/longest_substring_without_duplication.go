package longest_substring_without_duplication
// important question

import (
	"strings"
)

type substring struct {
	left  int
	right int
}

func (ss substring) length() int { return ss.right - ss.left }

// LongestSubstringWithoutDuplication
// O(n) time | O(min(n, a)) space
func LongestSubstringWithoutDuplication(str string) string {
	lastSeen := map[rune]int{}
	longest := substring{0, 1}
	startIndex := 0
	for i, char := range str {
		if seenIndex, found := lastSeen[char]; found && startIndex < seenIndex+1 {
			startIndex = seenIndex + 1
		}
		if longest.length() < i+1-startIndex {
			longest = substring{startIndex, i + 1}
		}
		lastSeen[char] = i
	}
	return str[longest.left:longest.right]
}

// c l e m e n t i s a c a p
// m e n t i s a c

// a b c d e a b c d e f c
// b c d e a b c d e f

// a b c d a b c e f

func longestSubstringWithoutDuplication(str string) string {
	letters := strings.Split(str, "")

	subStrings := make([]string, 0)
	seen := make(map[string]int)

	var leftIdx int
	for rightIdx, letter := range letters {
		if letterIdx, ok := seen[letter]; ok {
			subString := strings.Join(letters[leftIdx:rightIdx], "")
			subStrings = append(subStrings, subString)
			leftIdx = letterIdx+1
			seen = make(map[string]int)
			for i := leftIdx; i < rightIdx; i++ {
				seen[letters[i]] = i
			}
		}

		seen[letter] = rightIdx
	}
	subStrings = append(subStrings, strings.Join(letters[leftIdx:len(str)], ""))

	var maxIdx, maxLen int
	for idx, subString := range subStrings {
		if len(subString) > maxLen {
			maxLen = len(subString)
			maxIdx = idx
		}
	}
	return subStrings[maxIdx]
}
