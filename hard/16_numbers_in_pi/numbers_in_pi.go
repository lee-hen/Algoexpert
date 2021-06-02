package numbers_in_pi

import (
	"math"
)

// 3141592653589793238462643383279
// [314159265358979323846, 26433, 8, 3279, 314159265, 35897932384626433832, 79]
// 314159265 | 35897932384626433832 | 79

// 3|1|4|1
// 1 3 4

// [3141, 5, 31, 2, 4159, 9, 42]

// 3,  1,  4,  1,  5,  9,   2
// 3  calculate suffix
// 31  calculate suffix  4159 | 2
// 3141  calculate suffix
// 31415  calculate suffix
// 314159  calculate suffix
// 3141592  calculate suffix

// 31|4159|2

func NumbersInPi(pi string, numbers []string) int {
	favoriteNumbers := make(map[string]struct{})
	for _, number := range numbers {
		favoriteNumbers[number] = struct{}{}
	}

	digitsWithSpace := make([]byte, 0)
	smallestNumberOfSpaces := calculateSmallestNumberOfSpaces(0,  math.MaxInt32, []byte(pi), &digitsWithSpace, favoriteNumbers)

	if  smallestNumberOfSpaces == math.MaxInt32 {
		return -1
	}
	return smallestNumberOfSpaces
}

func calculateSmallestNumberOfSpaces(i, smallestNumberOfSpaces int, digits []byte, digitsWithSpace*[]byte, favoriteNumbers map[string]struct{}) int {
	prefixDigits := make([]byte, 0)
	for j := i; j < len(digits); j++ {
		digit := digits[j]
		prefixDigits = append(prefixDigits, digit)

		if _, ok := favoriteNumbers[string(prefixDigits)]; ok {
			*digitsWithSpace = append(*digitsWithSpace, prefixDigits...)

			if j < len(digits)-1 {
				*digitsWithSpace = append(*digitsWithSpace, '|')
			} else if len(*digitsWithSpace) >= len(digits) {
				pi := make([]byte, 0)
				var numberOfSpaces, idx int
				for _, digitOrSpace := range *digitsWithSpace {
					if digitOrSpace == '|' {
						numberOfSpaces++
					} else {
						idx++
						if digits[idx] == digitOrSpace {
							pi = append(pi, digitOrSpace)
						} else {
							idx--
							numberOfSpaces--
						}
					}
				}

				smallestNumberOfSpaces = min(smallestNumberOfSpaces, numberOfSpaces+1)
				*digitsWithSpace = make([]byte, 0)
			}

			smallestNumberOfSpaces = calculateSmallestNumberOfSpaces(j+1, smallestNumberOfSpaces, digits, digitsWithSpace, favoriteNumbers)
		}
	}

	return smallestNumberOfSpaces
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

