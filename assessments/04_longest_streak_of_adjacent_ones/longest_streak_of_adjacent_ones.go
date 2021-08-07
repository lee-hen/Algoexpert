package longest_streak_of_adjacent_ones
// important question

//  0  1  2  3  4  5  6  7  8  9  10  11 12 13 14 15 16 17
//  1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1,  1, 0, 1  1  1  1   1

// LongestStreakOfAdjacentOnes
// How smart it is!
// O(n) time | O(1) space - where n is the length of the input array
func LongestStreakOfAdjacentOnes(array []int) int {
	longestStreakLength := 0
	longestStreakReplacedZeroIdx := -1

	currentStreakLength := 0
	replacedZeroIdx := -1

	for i := 0; i < len(array); i++ {
		if array[i] == 1 {
			currentStreakLength++
		} else {
			currentStreakLength = i - replacedZeroIdx
			replacedZeroIdx = i
		}


		if currentStreakLength > longestStreakLength {
			longestStreakLength = currentStreakLength
			longestStreakReplacedZeroIdx = replacedZeroIdx
		}
	}

	return longestStreakReplacedZeroIdx
}


// my solution

// -1  0  1  2  3  4  5  6  7  8  9  10  11 12 13 14 15 16 17
//  0  1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1,  1, 0, 1  1  1  1   1
// 12  Replacing the 0 at index 8 with a 1 forms a streak of 6 adjacent 1s.

// 0 0 1
// 1 0 1
// 1 0 0

func longestStreakOfAdjacentOnes(array []int) int {
	var targetIdx = -1

	indices := make([]int, 0)
	for i, num := range array {
		if num == 1 &&  (i == 0 || i > 0 && array[i-1] == 0) {
			indices = append(indices, i)
		}

		if num == 0 && targetIdx == -1 {
			targetIdx = i
		}
	}

	var count int
	for _, idx := range indices {
		zeroIdx := idx-1
		var startIdx = idx

		for idx < len(array) && array[idx] == 1 {
			idx++
		}

		if idx < len(array) {
			zeroIdx = idx
			idx++
		}

		for idx < len(array) && array[idx] == 1 {
			idx++
		}

		if idx-startIdx >= count {
			count = idx-startIdx
			targetIdx = zeroIdx
		}
	}

	return targetIdx
}
