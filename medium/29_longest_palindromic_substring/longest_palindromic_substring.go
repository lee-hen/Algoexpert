package longest_palindromic_substring

//Notes
//In the video explanation of this question, we mistakenly state that the optimal solution runs with constant space.
//
//It's true that throughout our traversal of the input string, we only store an array of length 2; however, we ultimately still need to slice the longest palindromic substring out of the string, and this longest palindromic substring can be as large as the string itself, in the worst case.
//
//Thus, the algorithm runs with linear space.

type substring struct {
	left  int
	right int
}
func (ss substring) length() int {
	return ss.right - ss.left
}

func LongestPalindromicSubstring(str string) string {
	result := substring{0, 1}
	for i := 1; i < len(str); i++ {
		odd := getLongestPalindromeFrom(str, i-1, i+1)
		even := getLongestPalindromeFrom(str, i-1, i)
		longest := even
		if odd.length() > even.length() {
			longest = odd
		}
		if longest.length() > result.length() {
			result = longest
		}
	}
	return str[result.left:result.right]
}

func getLongestPalindromeFrom(str string, leftIndex, rightIndex int) substring {
	for leftIndex >= 0 && rightIndex < len(str) {
		if str[leftIndex] != str[rightIndex] {
			break
		}
		leftIndex -= 1
		rightIndex += 1
	}
	return substring{leftIndex + 1, rightIndex}
}


func LongestPalindromicSubstring1(str string) string {
	longest := ""
	for i := range str {
		for j := i; j < len(str); j++ {
			substr := str[i : j+1]
			if len(substr) > len(longest) && IsPalindrome(substr) {
				longest = substr
			}
		}
	}
	return longest
}

func IsPalindrome(str string) bool {
	for i := range str {
		j := len(str) - i - 1
		if str[i] != str[j] {
			return false
		}
	}
	return true
}

func longestPalindromicSubstring(str string) string {
	var subStringLength int
	var palindromicSubstring string
	for idx, r := range str {
		sub := getSubString(r, []rune(str)[:idx+1])
		if isPalindrome(sub) && len(sub) > subStringLength {
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

func isPalindrome(str []rune) bool {
	for i, j := 0, len(str)-1; i <= j; {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}
	return true
}
