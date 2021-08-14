package find_three_largest_numbers
// important question

import "math"

func FindThreeLargestNumbers(array []int) []int {
	three := []int{math.MinInt64, math.MinInt64, math.MinInt64}
	for _, num := range array {
		if num > three[2] {
			shiftAndUpdate(three, num, 2)
		} else if num > three[1] {
			shiftAndUpdate(three, num, 1)
		} else if num > three[0] {
			shiftAndUpdate(three, num, 0)
		}
	}
	return three
}

func shiftAndUpdate(three []int, num, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			three[i] = num
		} else {
			three[i] = three[i+1]
		}
	}
}

// My solution
func FindThreeLargestNumbers1(array []int) []int {
	const boundary = math.MinInt64 //-9223372036854775808
	largestNumbers := [3]int{boundary, boundary, boundary}
	for _, el := range array {
		if el > largestNumbers[2] {
			if largestNumbers[2] > largestNumbers[1] {
				if largestNumbers[1] > largestNumbers[0] {
					largestNumbers[0] = largestNumbers[1]
				}
				largestNumbers[1] = largestNumbers[2]
			}
			largestNumbers[2] = el
		} else if el > largestNumbers[1] {
			if largestNumbers[1] > largestNumbers[0] {
				largestNumbers[0] = largestNumbers[1]
			}
			largestNumbers[1] = el
		} else if el > largestNumbers[0] {
			largestNumbers[0] = el
		}
	}
	return largestNumbers[:]
}
