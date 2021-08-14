package largest_range
// important question

import (
	"math"
)

// LargestRange
// 0  1   2  3  4   5  6  7  8   9  10  11
// 1, 11, 3, 0, 15, 5, 2, 4, 10, 7, 12,  6
// 0 7  because 0 1 2 3 4 5 6 7 is in the input array
// the range is means find the largest continuous range of numbers
// and return the range of the first and the second
func LargestRange(array []int) []int {
	best := make([]int, 0)
	longestLength := 0

	nums := map[int]bool{}
	for _, num := range array {
		nums[num] = true
	}

	for _, num := range array {
		if !nums[num] {
			continue
		}

		nums[num] = false
		currentLength, left, right := 1, num-1, num+1

		for nums[left] {
			nums[left] = false
			currentLength += 1
			left -= 1
		}

		for nums[right] {
			nums[right] = false
			currentLength += 1
			right += 1
		}

		if currentLength > longestLength {
			longestLength = currentLength
			best = []int{left + 1, right - 1}
		}
	}
	return best
}

// my solution
func largestRange(slice []int) []int {
	ranges := make(map[int]int)

	for _, num := range slice {
		if _, found := ranges[num]; !found {
			ranges[num] = num
		}
	}

	maxCounter := math.MinInt32
	var first, second int
	for _, num := range slice {
		if _, found := ranges[num+1]; found {
			ranges[num] = num+1
		} else {
			temp, current, counter := num, num, 0
			for {
				if _, found = ranges[current]; found {
					current--
					counter++
				} else {
					break
				}
			}
			if counter > maxCounter {
				maxCounter = counter
				first = current+1
				second = temp
			}
		}
	}

	return []int{first, second}
}
