package longest_increasing_subsequence
// important question

// LongestIncreasingSubsequence
// this implementation is very hard to understand.
// O(nlogn) time | O(n) space
func LongestIncreasingSubsequence(array []int) []int {
	sequences := make([]int, len(array))

	// represent the length of subsequence
	// the values is the smallest value's index of subsequence from input array
	indices := make([]int, len(array)+1)
	for i := range array {
		sequences[i] = -1
		indices[i] = -1
	}
	length := 0
	for i, num := range array {
		newLength := binarySearch(1, length, indices, array, num)
		// put to the correct position
		indices[newLength] = i
		// store previous index
		sequences[i] = indices[newLength-1]
		length = max(length, newLength)
	}
	return buildSequence(array, sequences, indices[length])
}

func binarySearch(startIndex, endIndex int, indices, array []int, num int) int {
	if startIndex > endIndex {
		return startIndex
	}
	middleIndex := (startIndex + endIndex) / 2
	if array[indices[middleIndex]] < num {
		startIndex = middleIndex + 1
	} else {
		endIndex = middleIndex - 1
	}
	return binarySearch(startIndex, endIndex, indices, array, num)
}

//  0  1    2   3   4  5  6   7  8  9  10
//  5, 7, -24, 12, 10, 2, 3, 12, 5, 6, 35

// longest
//  0  1    2   3   4  5  6   7  8  9  10
//  1  2    1   3   3  2  3   4  4  5  6

// traceSequences(index)
//  0  1    2   3   4  5  6   7  8  9  10
// -1  0   -1   1   1  2  5   6  6  8  9

// subsequence
// -24 2 3 5 6 35

// longestIncreasingSubsequence
// my solution
// O(n^2) time | O(n) space
func longestIncreasingSubsequence(input []int) []int {
	longest := make([]int, len(input), len(input))
	traceSequences:= make([]int, len(input), len(input))
	for i := range traceSequences {
		traceSequences[i] = -1
	}

	for i := range longest {
		longest[i] = 1
	}

	var longestIndex int
	for i, currentNum := range input {
		// find previous longest subsequence number and index
		// the implementation below is the better way
		for j := 0; j < i; j++ {
			if input[j] < currentNum && longest[j]+1 >= longest[i] {
				longest[i] = longest[j] + 1
				// sequences store the previous index
				traceSequences[i] = j
			}
		}

		if longest[i] > longest[longestIndex] {
			longestIndex = i
		}
	}

	return buildSequence(input, traceSequences, longestIndex)
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

func max(arg int, rest ...int) int {
	for _, num := range rest {
		if num > arg {
			arg = num
		}
	}
	return arg
}
