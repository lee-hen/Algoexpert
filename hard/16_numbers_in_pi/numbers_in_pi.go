package numbers_in_pi
// important question

import (
	"math"
)

// [3141, 5, 31, 2, 4159, 9, 42]
// 31|4159|2

// similar to permutations dynamic programming
// 3,  1,  4,  1,  5,  9,   2

//2 calculate suffix            6
//9 calculate suffix            5
//92 calculate suffix
//5 calculate suffix            4
//59 calculate suffix
//592 calculate suffix
//1 calculate suffix            3
//15 calculate suffix
//159 calculate suffix
//1592 calculate suffix
//4 calculate suffix            2
//41 calculate suffix
//415 calculate suffix
//4159 calculate suffix
//41592 calculate suffix
//1 calculate suffix            1
//14 calculate suffix
//141 calculate suffix
//1415 calculate suffix
//14159 calculate suffix
//141592 calculate suffix
//3 calculate suffix            0
//31 calculate suffix
//314 calculate suffix
//3141 calculate suffix
//31415 calculate suffix
//314159 calculate suffix
//3141592 calculate suffix

func NumbersInPi(pi string, numbers []string) int {
	numbersTable := map[string]struct{}{}
	for _, number := range numbers {
		numbersTable[number] = struct{}{}
	}

	cache := map[int]int{}
	for i := len(pi) - 1; i >= 0; i-- {
		getMinSpaces(i, pi, numbersTable, cache)
	}

	if cache[0] == math.MaxInt32 {
		return -1
	}
	return cache[0]
}

func getMinSpaces(idx int, pi string, numbersTable map[string]struct{}, cache map[int]int) int {
	if idx == len(pi) {
		return -1
	} else if val, found := cache[idx]; found {
		return val
	}

	minSpaces := math.MaxInt32
	for i := idx; i < len(pi); i++ {
		prefix := pi[idx : i+1] // because string is immutable this may cause o(n) time complexity

		if _, found := numbersTable[prefix]; found {
			minSpacesInSuffix := getMinSpaces(i+1, pi, numbersTable, cache)
			//fmt.Println(prefix)
			minSpaces = min(minSpaces, minSpacesInSuffix+1)
		}
	}

	cache[idx] = minSpaces
	return cache[idx]
}

// [1 141 14 3]
// 3141
// 3|14|1
// 3|141

// 14|1 141 min(1 0)

func NumbersInPi1(pi string, numbers []string) int {
	numbersTable := map[string]struct{}{}
	for _, number := range numbers {
		numbersTable[number] = struct{}{}
	}

	minSpaces := getMinSpaces(0, pi, numbersTable, map[int]int{})
	if minSpaces == math.MaxInt32 {
		return -1
	}
	return minSpaces
}

// 3141592653589793238462643383279
// [314159265358979323846, 26433, 8, 3279, 314159265, 35897932384626433832, 79]
// 314159265 | 35897932384626433832 | 79

func numbersInPi(pi string, numbers []string) int {
	favoriteNumbers := make(map[string]struct{})
	for _, number := range numbers {
		favoriteNumbers[number] = struct{}{}
	}

	digitsWithSpace := make([]byte, 0)
	smallestNumberOfSpaces := calculateSmallestNumberOfSpaces(0, math.MaxInt32, []byte(pi), &digitsWithSpace, favoriteNumbers)

	if smallestNumberOfSpaces == math.MaxInt32 {
		return -1
	}
	return smallestNumberOfSpaces
}

func calculateSmallestNumberOfSpaces(i, smallestNumberOfSpaces int, digits []byte, digitsWithSpace *[]byte, favoriteNumbers map[string]struct{}) int {
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
