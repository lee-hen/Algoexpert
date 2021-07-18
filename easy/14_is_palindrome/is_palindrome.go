package is_palindrome

func IsPalindrome(str string) bool {
	return helper(str, 0)
}

func helper(str string, i int) bool {
	j := len(str) - i - 1
	if j >= i {
		return true
	}

	if str[j] != str[i] {
		return false
	}

	return helper(str, i+1)

}

// My solutions
func isPalindrome2(str string) bool {
	for i, j := 0, len(str)-1; i <= j; {

		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isPalindrome1(str string) bool {
	var middle = -1
	if len(str)%2 == 1 {
		middle = len(str) / 2
	}

	var runeValues []rune
	for i, runeValue := range str {
		if i == middle {
			continue
		}

		if len(runeValues) > 0 && runeValues[len(runeValues)-1] == runeValue {
			_, runeValues = runeValues[len(runeValues)-1], runeValues[:len(runeValues)-1]
		} else {
			runeValues = append(runeValues, runeValue)
		}
	}

	if len(runeValues) == 0 {
		return true
	}

	return false
}

func IsPalindromeRune(str []rune) bool {
	for i, j := 0, len(str)-1; i <= j; {
		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}
	return true
}
