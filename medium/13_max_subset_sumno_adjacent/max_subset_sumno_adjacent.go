package max_subset_sumno_adjacent

// MaxSubsetSumNoAdjacent
// better solution
func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	}

	first := max(array[0], array[1]) // array[i-1]
	second := array[0] // array[i-2]
	for i := 2; i < len(array); i++ {
		// ex.  [7 10 12 14]
		current := second+array[i] // 10 + 14
		previous := max(first, current) // max(array[i-1], array[i-2] + array[i])
		second = first // 10
		first = previous // 19
	}
	return first
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxSubsetSumNoAdjacent1
// my solution
func MaxSubsetSumNoAdjacent1(array []int) int {
	sum1 := sumNoAdjacent(0, array)
	sum2 := sumNoAdjacent(1, array)
	if sum1 > sum2 {
		return sum1
	}
	return sum2
}

func sumNoAdjacent(idx int, array []int) int {
	if idx > len(array)-1 {
		return 0
	}

	sum1 := array[idx] + sumNoAdjacent(idx+2, array)
	sum2 := array[idx] + sumNoAdjacent(idx+3, array)

	if sum1 > sum2 {
		return sum1
	}

	return sum2
}
