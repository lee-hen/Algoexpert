package longest_palindromic_substring
// important question

import p "github.com/lee-hen/Algoexpert/easy/14_is_palindrome"

//Notes
//In the video explanation of this question, we mistakenly state that the optimal solution runs with constant space.
//
//It's true that throughout our traversal of the input string, we only store an array of length 2; however, we ultimately still need to slice the longest palindromic substring out of the string, and this longest palindromic substring can be as large as the string itself, in the worst case.
//
//Thus, the algorithm runs with linear space.

type substring struct {
	Left  int
	Right int
}
func (ss substring) Length() int {
	return ss.Right - ss.Left
}

func LongestPalindromicSubstring(str string) string {
	result := substring{0, 1}
	for i := 1; i < len(str); i++ {
		odd := GetLongestPalindromeFrom(str, i-1, i+1)
		even := GetLongestPalindromeFrom(str, i-1, i)
		longest := even
		if odd.Length() > even.Length() {
			longest = odd
		}
		if longest.Length() > result.Length() {
			result = longest
		}
	}
	return str[result.Left:result.Right]
}

func GetLongestPalindromeFrom(str string, leftIndex, rightIndex int) substring {
	for leftIndex >= 0 && rightIndex < len(str) {
		if str[leftIndex] != str[rightIndex] {
			break
		}
		leftIndex -= 1
		rightIndex += 1
	}
	return substring{leftIndex + 1, rightIndex}
}


// longestPalindromicSubstring
// my solution
// O(n^2) time complexity
func longestPalindromicSubstring(str string) string {
	var subStringLength int
	var palindromicSubstring string
	for idx, r := range str {
		sub := getSubString(r, []rune(str)[:idx+1])
		if p.IsPalindromeRune(sub) && len(sub) > subStringLength {
			subStringLength = len(sub)
			palindromicSubstring = string(sub)
		}
	}
	return palindromicSubstring
}

func getSubString(target rune, runes []rune) []rune {
	right := len(runes)-1
	left := right
	current := right-1

	var prev int
	for current > 1 && runes[current] == runes[right] {
		current--
	}
	prev = current+1

	var diff int
	if right > current {
		diff = right-current-1
	}
	for current > 0 && runes[current] != target {
		current--
	}

	if left = current - diff; left < 0 {
		left = prev
	}

	return runes[left:right+1]
}

// LongestPalindromicSubstring2
// O(n^3) time complexity
func LongestPalindromicSubstring2(str string) string {
	longest := ""
	for i := range str {
		for j := i; j < len(str); j++ {
			substr := str[i : j+1]
			if len(substr) > len(longest) && p.IsPalindrome(substr) {
				longest = substr
			}
		}
	}
	return longest
}
