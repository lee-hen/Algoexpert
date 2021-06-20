package is_valid_subsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0

	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}

		arrIdx++
	}

	return seqIdx == len(sequence)
}

func isValidSubsequence2(array []int, sequence []int) bool {
	sequenceSize := len(sequence)
	counter := 0
	for _, a := range array {
		if counter == sequenceSize {
			break
		}

		if a == sequence[counter] {
			counter++
		}
	}

	return counter == sequenceSize
}

// My solution
func isValidSubsequence1(array []int, sequence []int) bool {
	sequenceSize := len(sequence)
	counter := 0
	for _, a := range array {
		if counter < sequenceSize && a == sequence[counter] {
			counter++
		}
	}

	return counter == sequenceSize
}
