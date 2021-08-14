package max_sumIncreasing_subsequence
// important question

import (
	"math"
)

//  0  1    2   3   4  5  6   7  8  9  10
//  5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35

// sums
//  0  1    2   3   4   5  6   7   8   9   10
//  5  12  -24  24  22  2  5  34  10   16  69

// traceIdx
// -1  0   -1   1   1 -1  5   4  6  8   7

// subsequence
// 5, 7, 10, 12, 35

// MaxSumIncreasingSubsequence
// I can understand it but it's difficult to come up this idea
func MaxSumIncreasingSubsequence(array []int) (int, []int) {
	sequences := make([]int, len(array))
	sums := make([]int, len(array))

	for i := range sequences {
		sequences[i] = -1
		sums[i] = array[i]
	}

	maxSumIndex := 0
	for i, currentNum := range array {
		for j := 0; j < i; j++ {
			if array[j] < currentNum && sums[j]+currentNum >= sums[i] {
				sums[i] = sums[j] + currentNum
				// sequences store the previous index
				sequences[i] = j
			}
		}
		if sums[i] > sums[maxSumIndex] {
			maxSumIndex = i
		}
	}

	return sums[maxSumIndex], buildSequence(array, sequences, maxSumIndex)
}

func buildSequence(array []int, sequences []int, currentIdx int) []int {
	sequence := make([]int, 0)
	for currentIdx != -1 {
		sequence = append(sequence, array[currentIdx])
		// sequences store the previous index
		currentIdx = sequences[currentIdx]
	}
	reverse(sequence)
	return sequence
}

func reverse(numbers []int) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

// maxSumIncreasingSubsequence well my solution not work very well
// O(n^3) time O(n^3) space because of traverse back the slice
func maxSumIncreasingSubsequence(slice []int) (int, []int) {
	var maxVal = math.MinInt32
	var maxSubsequence []int

	for i := 0; i < len(slice); i++ {
		sum := slice[i]
		subsequence := []int{slice[i]}

		for j := i+1; j < len(slice); j++ {
			if slice[i] < slice[j]  {
				if subsequence[len(subsequence)-1] < slice[j] {
					subsequence = append(subsequence, slice[j])
					sum+=slice[j]
				} else {
					previousSum := sum
					previousSubsequence := make([]int, len(subsequence))
					copy(previousSubsequence, subsequence)

					l := j
					for ; l < len(slice); l++ {
						ll := len(subsequence)-1
						for ; ll > 0; ll-- {
							if subsequence[ll] < slice[l] {
								break
							}
							sum-= subsequence[ll]
						}
						subsequence = subsequence[:ll+1]

						if subsequence[len(subsequence)-1] < slice[l] {
							subsequence = append(subsequence, slice[l])
							sum+=slice[l]
						}

						if previousSum < sum {
							previousSum = sum
							previousSubsequence = make([]int, len(subsequence))
							copy(previousSubsequence, subsequence)
						}
					}

					if sum < previousSum {
						sum = previousSum
						subsequence = previousSubsequence
					}

					j = l
				}
			}
		}

		if maxVal < sum {
			maxVal = sum
			maxSubsequence = subsequence
		}
	}

	return maxVal, maxSubsequence
}
