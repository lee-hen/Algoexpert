package count_inversions
// important question

// 0  1  2  3  4  5  6
// 8, 5, 2, 9, 5, 6, 3

// divide
// left             right
// [8 5 2 9]        [5 6 3]
// [8 5] [2 9]      [5 6] [3]
// [8] [5] [2] [9]  [5] [6] [3]

// merge (conquer)
// left             right
// [5 8]  [2 9]     [5 6] [3]
// [2 5 8 9]        [3 5 6]
// [2 3 5 5 6 8 9]

func CountInversions(array []int) int {
	if len(array) == 0 {
		return 0
	}
	auxiliaryArray := make([]int, len(array))
	copy(auxiliaryArray, array)
	return sort(array, 0, len(array)-1, auxiliaryArray)
}

func sort(mainArray []int, startIdx, endIdx int, auxiliaryArray []int) int {
	if startIdx == endIdx {
		return 0
	}
	middleIdx := (startIdx + endIdx) / 2
	leftInversions := sort(auxiliaryArray, startIdx, middleIdx, mainArray)
	rightInversions := sort(auxiliaryArray, middleIdx+1, endIdx, mainArray)
	return leftInversions + rightInversions + merge(mainArray, startIdx, middleIdx, endIdx, auxiliaryArray)
}

func merge(mainArray []int, startIdx, middleIdx, endIdx int, auxiliaryArray []int) int {
	k := startIdx
	i := startIdx
	j := middleIdx + 1

	var inversions int
	for i <= middleIdx && j <= endIdx {
		if auxiliaryArray[i] <= auxiliaryArray[j] {
			mainArray[k] = auxiliaryArray[i]
			i++
		} else {
			// because left side is already sorted and also right side
			// if auxiliaryArray[i] > auxiliaryArray[j]
			// this means the whole left[i:] side is larger than auxiliaryArray[j]
			inversions += middleIdx-i+1
			mainArray[k] = auxiliaryArray[j]
			j++
		}
		k++
	}
	for i <= middleIdx {
		mainArray[k] = auxiliaryArray[i]
		i++
		k++
	}
	for j <= endIdx {
		mainArray[k] = auxiliaryArray[j]
		j++
		k++
	}

	return inversions
}
