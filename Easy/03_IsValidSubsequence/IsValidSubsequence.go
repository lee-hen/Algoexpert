package isvalidsubsequence

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

func IsValidSubsequence2(array []int, sequence []int) bool {
	sequence_size := len(sequence)
	counter := 0
	for _, a := range array {
		if counter == sequence_size {
			break
		}

		if a == sequence[counter] {
			counter++
		}
	}

	return counter == sequence_size
}

// My solution
func IsValidSubsequence1(array []int, sequence []int) bool {
	sequence_size := len(sequence)
	counter := 0
	for _, a := range array {
		if counter < sequence_size && a == sequence[counter] {
			counter++
		}
	}

	return counter == sequence_size
}
