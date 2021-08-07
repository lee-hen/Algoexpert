package words_in_phone_number

import (
	"strconv"
	"strings"
)

//  ----- ----- -----
// |     |     |     |
// |  1  |  2  |  3  |
// |     | abc | def |
// ----- ----- -----
// |     |     |     |
// |  4  |  5  |  6  |
// | ghi | jkl | mno |
// ----- ----- -----
// |     |     |     |
// |  7  |  8  |  9  |
// | pqrs| tuv | wxyz|
//  ----- ----- -----
// |     |
// |  0  |
// |     |
//  -----

// phoneNumber = "3662277"
// words = ["foo", "bar", "baz", "foobar", "emo", "cap", "car", "cat"]

// ["bar", "cap", "car", "emo", "foo", "foobar"]
// The words could be ordered differently.

// def mno mno abc abc pqrs pqrs

func wordsInPhoneNumber(phoneNumber string, words []string) []string {
	letters := []byte{
		byte('a'), byte('b'), byte('c'),     //  2
		byte('d'), byte('e'), byte('f'),     //  3
		byte('g'), byte('h'), byte('i'),     //  4
		byte('j'), byte('k'), byte('l'),     //  5
		byte('m'), byte('n'), byte('o'),     //  6
		byte('p'), byte('q'), byte('r'), byte('s'),  // 7
		byte('t'), byte('u'), byte('v'),             // 8
		byte('w'), byte('x'), byte('y'), byte('z'),  // 9
	}

	charOfPhoneNumber := make(map[byte]int)

	var j int
	for i := 0; i < 10; i++ {
		if i == 0  {
			charOfPhoneNumber[byte('0')] = i
		} else if i == 1 {
			charOfPhoneNumber[byte('0')] = i
		} else if i < 7 || i == 8 {
			chars := letters[j:j+3]
			for _, char := range chars {
				charOfPhoneNumber[char] = i
			}
			j = j+3
		} else if i == 7 || i == 9 {
			chars := letters[j:j+4]
			for _, char := range chars {
				charOfPhoneNumber[char] = i
			}
			j = j+4
		}
	}

	result := make([]string, 0)

	for _, word := range words {
		var generatePhoneNumber string
		for _, char := range []byte(word) {
			if num, ok := charOfPhoneNumber[char]; ok {
				generatePhoneNumber += strconv.Itoa(num)
			}
		}

		if strings.Index(phoneNumber, generatePhoneNumber) >= 0 {
			result = append(result, word)
		}
	}

	return result
}
