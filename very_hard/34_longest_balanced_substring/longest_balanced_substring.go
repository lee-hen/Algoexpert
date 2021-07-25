package longest_balanced_substring

const left = '('
const right = ')'

// LongestBalancedSubstring
// O(n) time | O(1) space - where n is the length of the input string
func LongestBalancedSubstring(str string) int {
	return max(getLongestBalancedInDirection(str, true), getLongestBalancedInDirection(str, false))
}

func getLongestBalancedInDirection(str string, leftToRight bool) int {
	openingParens := left
	startIdx := 0
	step := 1
	if !leftToRight {
		openingParens = right
		startIdx = len(str) - 1
		step = -1
	}

	maxLength := 0

	openingCount := 0
	closingCount := 0

	idx := startIdx
	for idx >= 0 && idx < len(str) {
		char := str[idx]

		if rune(char) == openingParens {
			openingCount++
		} else {
			closingCount++
		}

		if openingCount == closingCount {
			maxLength = max(maxLength, closingCount*2)
		} else if closingCount > openingCount {
			openingCount = 0
			closingCount = 0
		}

		idx += step
	}

	return maxLength
}


// LongestBalancedSubstring2
// O(n) time | O(1) space - where n is the length of the input string
func LongestBalancedSubstring2(str string) int {
	maxLength := 0

	openingCount := 0
	closingCount := 0

	for _, char := range str {
		if char == left {
			openingCount++
		} else {
			closingCount++
		}

		if openingCount == closingCount {
			maxLength = max(maxLength, closingCount*2)
		} else if closingCount > openingCount {
			openingCount = 0
			closingCount = 0
		}
	}

	openingCount = 0
	closingCount = 0

	for i := len(str) - 1; i >= 0; i-- {
		char := str[i]

		if char == left {
			openingCount++
		} else {
			closingCount++
		}

		if openingCount == closingCount {
			maxLength = max(maxLength, openingCount*2)
		} else if openingCount > closingCount {
			openingCount = 0
			closingCount = 0
		}
	}

	return maxLength
}

func LongestBalancedSubstring1(str string) int {
	maxLength := 0
	idxStack := []int{-1}

	for i := range str {
		if str[i] == left {
			idxStack = append(idxStack, i)
		} else {
			idxStack = idxStack[:len(idxStack)-1]
			if len(idxStack) == 0 {
				idxStack = append(idxStack, i)
			} else {
				balancedSubstringStartIdx := idxStack[len(idxStack)-1]
				currentLength := i - balancedSubstringStartIdx
				maxLength = max(maxLength, currentLength)
			}
		}
	}

	return maxLength
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 0 1 2 3 4 5
// ( ( ) ) ) (

// 4 5

// 0 1 2 3 4 5 6 7 8 9 10
// ( ) ) ( ) ( ( ) ( )  )

// 2


// 0 1 2 3 4 5 6 7 8 9 10 11 12 13  14 15 16 17 18 19
// ( ) ) ) ) ) ) ( ) ) (  (  (  (   )  )  (  (  (   )

// 2 3 4 5 6 9 10 11 16 17 19

func longestBalancedSubstring(str string) int {
	var indices []int

	for idx := 0; idx < len(str); idx++ {
		var removed bool
		if len(indices) != 0 {
			open := indices[len(indices)-1]

			if str[idx] == right && str[open] == left {
				indices = indices[:len(indices)-1]
				removed = true
			}
		}

		if !removed {
			if str[idx] == left || str[idx] == right {
				indices = append(indices, idx)
			}
		}
	}


	if len(indices) == 0 {
		return len(str)
	}

	longest := indices[0]
	for i := 0; i < len(indices)-1; i++ {
		if indices[i+1]-indices[i] > 2 {
			if indices[i+1]-indices[i]-1 > longest {
				longest = indices[i+1]-indices[i]-1
			}
		}
	}

	if len(str)-1-indices[len(indices)-1] > longest {
		longest = len(str)-1-indices[len(indices)-1]
	}

	return longest
}
