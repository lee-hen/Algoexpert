package reverse_words_in_string

import "strings"

func ReverseWordsInString(str string) string {
	characters := make([]byte, 0)
	for _, char := range []byte(str) {
		characters = append(characters, char)
	}
	reverseListRange(characters, 0, len(characters)-1)
	startOfWord := 0

	for startOfWord < len(characters) {
		endOfWord := startOfWord
		for endOfWord < len(characters) && characters[endOfWord] != ' ' {
			endOfWord += 1
		}

		reverseListRange(characters, startOfWord, endOfWord-1)
		startOfWord = endOfWord + 1
	}
	return string(characters)
}

func reverseListRange(list []byte, rangeStart, rangeEnd int) {
	start := rangeStart
	end := rangeEnd

	for start < end {
		list[start], list[end] = list[end], list[start]
		start += 1
		end -= 1
	}
}

func ReverseWordsInString1(str string) string {
	slice := make([]string, 0)
	startOfWord := 0
	for idx, character := range str {
		if character == ' ' {
			slice = append(slice, str[startOfWord:idx])
			startOfWord = idx
		} else if str[startOfWord] == ' ' {
			slice = append(slice, " ")
			startOfWord = idx
		}
	}

	slice = append(slice, str[startOfWord:])
	ReverseSlice(slice)
	return strings.Join(slice, "")
}

func ReverseSlice(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
