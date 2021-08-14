package run_length_encoding
// important question

import (
	"strconv"
	"unicode/utf8"
)

// RunLengthEncoding
// Better solution
func RunLengthEncoding(str string) string {
	encodedStringCharacters := make([]byte, 0)
	currentRunLength := 1
	for i := 1; i < len(str); i++ {
		currentCharacter := str[i]
		previousCharacter := str[i-1]

		if currentCharacter != previousCharacter || currentRunLength == 9 {
			encodedStringCharacters = append(encodedStringCharacters, strconv.Itoa(currentRunLength)[0])
			encodedStringCharacters = append(encodedStringCharacters, previousCharacter)
			currentRunLength = 0
		}

		currentRunLength++
	}

	encodedStringCharacters = append(encodedStringCharacters, strconv.Itoa(currentRunLength)[0])
	encodedStringCharacters = append(encodedStringCharacters, str[len(str)-1])
	return string(encodedStringCharacters)
}

// My solution
func runLengthEncoding(str string) string {
	strSlice := make([]string, 0, utf8.RuneCountInString(str))
	for _, char := range str {
		strSlice = append(strSlice, string(char))
	}

	var encodingString string
	var idx int
	out:
	for ; idx < len(strSlice); idx++ {
		var encodingStr string
		var nextStr string
		var i int

		if idx + 1 == len(strSlice) {
			i = 1
		} else {
			nextStr = strSlice[idx + 1]
			if strSlice[idx] != nextStr {
				i = 1
			}
		}

		if i == 1 {
			encodingStr = strconv.Itoa(i) + strSlice[idx]
		}

		for ; strSlice[idx] == nextStr; idx++ {
			if i++; i > 9 {
				idx--
				encodingString += encodingStr
				continue out
			}
			encodingStr = strconv.Itoa(i) + strSlice[idx]

			if idx == len(strSlice) - 1 {
				encodingString += encodingStr
				break out
			}
		}

		if i > 1 {
			idx--
		}
		encodingString += encodingStr
	}

	return encodingString
}
