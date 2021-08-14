package smallest_substring_containing
// important question

import (
	"math"
	"unicode/utf8"
)

// SmallestSubstringContaining
// O(b + s) time | O(b + s) space - where b is the length of the big
// input string and s is the length of the small input string
func SmallestSubstringContaining(bigString, smallString string) string {
	targetCharCounts := getCharCounts(smallString)
	substringBounds := getSubstringBounds(bigString, targetCharCounts)
	return getStringFromBounds(bigString, substringBounds)
}

func getCharCounts(str string) map[byte]int {
	charCounts := map[byte]int{}
	for _, char := range str {
		increaseCharCount(byte(char), charCounts)
	}
	return charCounts
}

func getSubstringBounds(str string, targetCharCounts map[byte]int) []int {
	substringBounds := []int{0, math.MaxInt32}
	substringCharCounts := map[byte]int{}
	var leftIdx, rightIdx, numUniqueCharsDone int

	// Move the rightIdx to the right in the string until you've counted
	// all of the target characters enough times.
	for rightIdx < len(str) {
		rightChar := str[rightIdx]
		if _, found := targetCharCounts[rightChar]; !found {
			rightIdx++
			continue
		}
		increaseCharCount(rightChar, substringCharCounts)
		if substringCharCounts[rightChar] == targetCharCounts[rightChar] {
			numUniqueCharsDone++
		}

		// Move the leftIdx to the right in the string until you no longer
		// have enough of the target characters in between the leftIdx and
		// the rightIdx. Update the substringBounds accordingly.
		for numUniqueCharsDone == len(targetCharCounts) && leftIdx <= rightIdx {
			substringBounds = getCloserBounds(leftIdx, rightIdx, substringBounds[0], substringBounds[1])
			leftChar := str[leftIdx]
			if _, found := targetCharCounts[leftChar]; !found {
				leftIdx++
				continue
			}
			if substringCharCounts[leftChar] == targetCharCounts[leftChar] {
				numUniqueCharsDone--
			}
			decreaseCharCount(leftChar, substringCharCounts)
			leftIdx++
		}
		rightIdx++
	}
	return substringBounds
}

func getCloserBounds(idx1, idx2, idx3, idx4 int) []int {
	if idx2-idx1 < idx4-idx3 {
		return []int{idx1, idx2}
	}
	return []int{idx3, idx4}
}

func getStringFromBounds(str string, bounds []int) string {
	start, end := bounds[0], bounds[1]
	if end == math.MaxInt32 {
		return ""
	}
	return str[start : end+1]
}

func increaseCharCount(char byte, charCounts map[byte]int) {
	charCounts[char]++
}

func decreaseCharCount(char byte, charCounts map[byte]int) {
	charCounts[char]--
}

// 0 1 2 3 4 5 6 7 8 9 10 11 12 13
// a b c d $ e f $ a x  b  $  c  $
//             l                 r


// $  $  a  b  f
// 4  7  8  10 6

// leftIdx 4
// rightIdx 10


// f $ a x b $   6-11

// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33
// a $ f u u + a f f f  +  a  f  f  a  f  f  a  +  a  $  A  f  f  a  b  +  a  +  a  +  $  a  $
// l                                                                       r

//
// a 7   0 6 11 14 17 19 24
// + 3   10 18 26
// A 1   21
// $ 2   1 20


// 0 24   6 27   11 29  14 32
// 5 18  10 26   18 27  26 30
// 21
// 1 20  20 31   31 33

// a a a a a a a + + + A $ $

// a f f a + a $ A f f a b + a + a + $ a

// smallestSubstringContaining
// my solution
func smallestSubstringContaining(bigString, smallString string) string {
	smallStr := make(map[rune][]int)

	for _, str := range smallString {
		smallStr[str] = append(smallStr[str], -1)
	}

	var leftIdx int
	rightIdx := len(bigString)-1
	for i, str := range bigString {
		if _, ok := smallStr[str]; ok {
			firstMatchStr := firstMatchSmallStr(i, str, smallStr)
			if !firstMatchStr {
				leftStr,_ := utf8.DecodeRuneInString(bigString[leftIdx:])

				smallStr[str] = smallStr[str][1:]
				smallStr[str] = append(smallStr[str], i)

				if str == leftStr {
					nextLeftIdx := findNextLeftPosition(smallStr)
					nextRightIdx := i

					if nextRightIdx-nextLeftIdx <= rightIdx-leftIdx {
						rightIdx = nextRightIdx
						leftIdx = nextLeftIdx
					}
				}
			} else {
				nextLeftIdx := findNextLeftPosition(smallStr)
				if nextLeftIdx != -1 {
					leftIdx = nextLeftIdx
				}
				rightIdx = i
			}
		}
	}

	for _, indices := range smallStr {
		for _, idx := range indices {
			if idx == -1 {
				return ""
			}
		}
	}

	return bigString[leftIdx:rightIdx+1]
}

func firstMatchSmallStr(i int, str rune, smallStr map[rune][]int) bool {
	for idx := range smallStr[str] {
		if smallStr[str][idx] == -1 {
			smallStr[str][idx] = i
			return true
		}
	}
	return false
}

func findNextLeftPosition(smallStr map[rune][]int) int {
	nextIdx := math.MaxInt32
	for _, indices := range smallStr {
		nextIdx = min(nextIdx, indices[0])
	}

	return nextIdx
}

func min(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num < curr {
			curr = num
		}
	}
	return curr
}
