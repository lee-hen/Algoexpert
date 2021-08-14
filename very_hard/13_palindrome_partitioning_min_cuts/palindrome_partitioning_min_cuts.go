package palindrome_partitioning_min_cuts
// important question

import (
	p "github.com/lee-hen/Algoexpert/easy/14_is_palindrome"
	"math"
)

//          n o o o n
//      o o n o o
//    n o o n o o o n

// n | o o | n o o o n
// n | o o n o o | o | n

//     0 1 2 3 4 5 6 7
//   - n o o n o o o n
// 0 0 1 0 0 0 0 0 0 0
// 1 0 1 2 0 0 0 0 0 0
// 2 0 1 2 2 0 0 0 0 0
// 3 0 1 2 2 1 0 0 0 0
// 4 0 1 2 2 1 2 0 0 0
// 5 0 1 2 2 1 2 2 0 0
// 6 0 1 2 2 1 2 2 3 0
// 7 0 1 2 2 1 2 2 3 3

//  0   1  2  3  4  5  6  7  8  9
//      0  1  2  3  4  5  6  7  8
//      n  o  o  n  a  b  b  a  d
// 1 0  1  0  0  0  0  0  0  0  0
// 2 0  1  2  0  0  0  0  0  0  0
// 3 0  1  2  2  0  0  0  0  0  0
// 4 0  1  2  2  1  0  0  0  0  0
// 5 0  1  2  2  1  2  0  0  0  0
// 6 0  1  2  2  1  2  3  0  0  0
// 7 0  1  2  2  1  2  3  3  0  0
// 8 0  1  2  2  1  2  3  3  2  0
// 9 0  1  2  2  1  2  3  3  2  3


//    a b b b
//1 0 1 0 0 0
//2 0 1 2 2 2
//3 0 1 2 2 2
//4 0 1 2 2 2


// PalindromePartitioningMinCuts2
// O(n^2) time | O(n^2) space
func PalindromePartitioningMinCuts2(str string) int {
	palindromes := make([][]bool, len(str))
	for i := range palindromes {
		palindromes[i] = make([]bool, len(str))
	}
	for i := range str {
		palindromes[i][i] = true
	}
	// very interesting.
	for length := 2; length <= len(str); length++ {
		for i := 0; i < len(str)-length+1; i++ {
			j := i + length - 1 // i start j end
			if length == 2 {
				palindromes[i][j] = str[i] == str[j]
			} else {
				palindromes[i][j] = str[i] == str[j] && palindromes[i+1][j-1]
			}
		}
	}
	cuts := make([]int, len(str))
	for i := range cuts {
		cuts[i] = math.MaxInt32
	}
	for i := range str {
		if palindromes[0][i] {
			cuts[i] = 0
		} else {
			cuts[i] = cuts[i-1] + 1
			for j := 1; j < i; j++ {
				if palindromes[j][i] && cuts[j-1]+1 < cuts[i] {
					cuts[i] = cuts[j-1] + 1
				}
			}
		}
	}
	return cuts[len(cuts)-1]
}

// PalindromePartitioningMinCuts
// O(n^2) time | O(n) space
func PalindromePartitioningMinCuts(str string) int {
	if len(str) == 0 {
		return -1
	}
	result := make([]int, len(str)+1, len(str)+1)
	palindromeRangesSoFar := make(map[int]int)

	for i := 0; i < len(str); i++ {
		findRangeIndexOfPalindrome(str, i-1, i+1, palindromeRangesSoFar)
		findRangeIndexOfPalindrome(str, i-1, i, palindromeRangesSoFar)

		var palindromeLeftIdx = -1
		if idx, ok := palindromeRangesSoFar[i]; ok {
			palindromeLeftIdx = idx
		}

		if palindromeLeftIdx > -1 {
			result[i+1] = min(result[i], result[palindromeLeftIdx]) + 1
		} else {
			result[i+1] = result[i] + 1
		}
	}

	return result[len(result)-1]-1
}

func findRangeIndexOfPalindrome(str string, leftIndex, rightIndex int, palindromeResult map[int]int) {
	for leftIndex >= 0 && rightIndex < len(str) {
		if str[leftIndex] != str[rightIndex] {
			break
		}
		if _, ok := palindromeResult[rightIndex]; !ok {
			palindromeResult[rightIndex] = leftIndex
		}

		leftIndex -= 1
		rightIndex += 1
	}
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

func PalindromePartitioningMinCuts1(s string) int {
	palindromes := make([][]bool, len(s))
	for i := range palindromes {
		palindromes[i] = make([]bool, len(s))
	}
	for i := range s {
		for j := i; j < len(s); j++ {
			palindromes[i][j] = p.IsPalindrome(s[i : j+1])
		}
	}
	cuts := make([]int, len(s))
	for i := range cuts {
		cuts[i] = math.MinInt32
	}
	for i := range s {
		if palindromes[0][i] {
			cuts[i] = 0
		} else {
			cuts[i] = cuts[i-1] + 1
			for j := 1; j < i; j++ {
				if palindromes[j][i] && cuts[j-1]+1 < cuts[i] {
					cuts[i] = cuts[j-1] + 1
				}
			}
		}
	}
	return cuts[len(s)-1]
}
