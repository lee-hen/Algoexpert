package single_cycle_check
// important question

// HasSingleCycle
// input [2, 3, 1, -4, -4, 2]
// output true
// better solution
//Notes
//In the video explanation of this question, we explain that we need to handle negative values for the nextIdx calculated in our helper method.
//
//nextIdx = (currentIdx + jump) % len(array)
//return nextIdx if nextIdx >= 0 else nextIdx + len(array)
//In most programming languages, this is necessary because, if currentIdx + jump is negative, then (currentIdx + jump) % len(array) will also be negative.
//
//In Python, however, "the modulo operator always yields a result with the same sign as its second operand (or zero)" [Python Docs]. In other words, in Python, the modulo operation to compute the nextIdx will always return a number with the sign of len(array), which is naturally positive.
//
//More specifically, the modulo operator in Python behaves as follows when used with a negative first operand:
//
//-x % y == -(x % y) + y
//The Python modulo operator effectively does, by default, what we do in our code to handle negative values.
//
//Thus, in Python, we can just return (currentIdx + jump) % len(array) for the nextIdx, without needing to handle negative values.
func HasSingleCycle(array []int) bool {
	numElementsVisited := 0
	currentIdx := 0
	for numElementsVisited < len(array) {
		if numElementsVisited > 0 && currentIdx == 0 {
			return false
		}
		numElementsVisited++
		currentIdx = getNextIdx(currentIdx, array)
	}
	return currentIdx == 0
}
func getNextIdx(currentIdx int, array []int) int {
	jumpIdx := array[currentIdx]
	nextIdx := (currentIdx + jumpIdx) % len(array)
	if nextIdx >= 0 {
		return nextIdx
	}
	return nextIdx + len(array)
}

//-1 = len(array) -1
//-2 = len(array) -2
//-k = len(array) -k

// my solution
func hasSingleCycle(array []int) bool {
	expect := (len(array) - 1) * len(array) / 2
	var total int
	for idx, arr := range array {
		nextIdx := jump(arr, idx, array)
		nextJump := array[nextIdx]
		nextJumpIdx := jump(nextJump, nextIdx, array)
		total += nextJumpIdx
		if len(array) > 2 && nextJumpIdx == idx && arr == nextJump*-1 {
			break
		}
	}
	if expect == total {
		return true
	}
	return false
}

func jump(k, idx int, array []int) int {
	jumpIdx := (idx + k) % len(array)
	if jumpIdx < 0 {
		return jumpIdx + len(array)
	}

	return jumpIdx
}

// similar codility problem
func cyclicRotation(array []int, k int) []int {
	result := make([]int, len(array), len(array))
	for idx := range array {
		jumpIdx := (idx + k) % len(array)
		if jumpIdx < 0 {
			jumpIdx += len(array)
		}
		result[jumpIdx] = array[idx]
	}
	return result
}
