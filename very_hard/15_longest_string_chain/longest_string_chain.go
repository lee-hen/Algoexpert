package longest_string_chain
// important question

import (
	"math"
	"sort"
)

type Chain struct {
	NextString     string
	MaxChainLength int
}

// LongestStringChain
// O(n * m^2 + nlog(n)) time | O(nm) space - where n is the number of strings and
// m is the length of the longest string
func LongestStringChain(strings []string) []string {
	// For every string, imagine the longest string chain that starts with it.
	// Set up every string to point to the next string in its respective longest
	// string chain. Also keep track of the lengths of these longest string chains.
	stringChains := map[string]*Chain{}
	for _, str := range strings {
		stringChains[str] = &Chain{NextString: "", MaxChainLength: 1}
	}

	// Sort the strings based on their length so that whenever we visit a
	// string (as we iterate through them from left to right), we can
	// already have computed the longest string chains of any smaller strings.
	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) < len(strings[j])
	})
	sortedStrings := strings

	for _, str := range sortedStrings {
		findLongestStringChain(str, stringChains)
	}
	return buildLongestStringChain(strings, stringChains)
}

func findLongestStringChain(str string, stringChains map[string]*Chain) {
	// Try removing every letter of the current string to see if the
	// remaining strings form a string chain.
	for i := range str {
		smallerString := getSmallerString(str, i)
		if _, found := stringChains[smallerString]; !found {
			continue
		}
		tryUpdateLongestStringChain(str, smallerString, stringChains)
	}
}

func getSmallerString(str string, index int) string {
	return str[:index] + str[index+1:]
}

func tryUpdateLongestStringChain(currentString, smallerString string, stringChains map[string]*Chain) {
	smallerStringChainLength := stringChains[smallerString].MaxChainLength
	currentStringChainLength := stringChains[currentString].MaxChainLength
	// Update the string chain of the current string only if the smaller string leads
	// to a longer string chain.
	if smallerStringChainLength+1 > currentStringChainLength {
		stringChains[currentString].MaxChainLength = smallerStringChainLength + 1
		stringChains[currentString].NextString = smallerString
	}
}

func buildLongestStringChain(strings []string, stringChains map[string]*Chain) []string {
	// Find the string that starts the longest string chain.
	maxChainLength := 0
	chainStartingString := ""
	for _, str := range strings {
		if stringChains[str].MaxChainLength > maxChainLength {
			maxChainLength = stringChains[str].MaxChainLength
			chainStartingString = str
		}
	}

	// Starting at the string found above, build the longest string chain.
	ourLongestStringChain := []string{}
	currentString := chainStartingString
	for currentString != "" {
		ourLongestStringChain = append(ourLongestStringChain, currentString)
		currentString = stringChains[currentString].NextString
	}
	if len(ourLongestStringChain) == 1 {
		return []string{}
	}
	return ourLongestStringChain
}

//  ["abde", "abc", "abd", "abcde", "ade", "ae", "1abde", "abcdef"]
// ["abcdef", "abcde", "abde", "ade", "ae"]

// 0 abcdef  4
// 1 1abde   3
// 2 abcde   3
// 3 abde    2
// 4 abd     0
// 5 ade     1
// 6 abc     0
// 7 ae      0

//    0       1      2     3     4    5    6    7
//  abcdef  1abde  abcde  abde  abd  ade  abc   ae
//    2      -1       3     5    -1   7   -1    -1
//            2
//                                         -1

// longestStringChain
// my solution recursive approach
func longestStringChain(strings []string) []string {
	sort.Slice(strings, func(i, j int) bool {
		return len(strings[i]) > len(strings[j])
	})

	stringChanTraceIndices := make([]int, len(strings), len(strings))
	for i := range stringChanTraceIndices {
		stringChanTraceIndices[i] = -1
	}

	lengthOfStringChain := make(map[int]int)
	stringIndices := make(map[string]int)

	for i, str := range strings {
		stringIndices[str] = i
	}

	for i, str := range strings {
		if _, ok := lengthOfStringChain[i]; ok {
			continue
		}

		findStringChain(i, str, stringIndices, lengthOfStringChain, stringChanTraceIndices)
	}

	return buildStringChain(strings, stringChanTraceIndices, maxLength(lengthOfStringChain))
}

func findStringChain(i int, str string, stringIndices map[string]int, lengthOfStringChain map[int]int, stringChanTraceIndices []int) {
	for j := 0; j < len(str); j++ {
		char := str[:j]
		char += str[j+1:]

		if idx, ok := stringIndices[char]; ok {
			findStringChain(idx,  char, stringIndices, lengthOfStringChain, stringChanTraceIndices)

			if lengthOfStringChain[i] < lengthOfStringChain[idx] + 1 {
				stringChanTraceIndices[i] = idx
			}

			lengthOfStringChain[i] = max(lengthOfStringChain[i], lengthOfStringChain[idx] + 1)
		}
	}
}

func max(arg int, rest ...int) int {
	for _, num := range rest {
		if num > arg {
			arg = num
		}
	}
	return arg
}

func maxLength(lengthOfStringChain map[int]int) int {
	var idx int
	maxLength := math.MinInt32
	for i, length := range lengthOfStringChain {
		if length > maxLength {
			maxLength = length
			idx = i
		}
	}

	return idx
}

func buildStringChain(strings []string, stringChanTraceIndices []int, currentIdx int) []string {
	stringChan := make([]string, 0)
	for currentIdx != -1 {
		stringChan = append(stringChan, strings[currentIdx])
		currentIdx = stringChanTraceIndices[currentIdx]
	}

	if len(stringChan) == 1 {
		return []string{}
	}

	return stringChan
}
