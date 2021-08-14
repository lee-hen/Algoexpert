package merge_sort
// important question

// MergeSort
// Best: O(nlog(n)) time | O(n) space
// Average: O(nlog(n)) time | O(n) space
// Worst: O(nlog(n)) time | O(n) space
func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	auxiliaryArray := make([]int, len(array))
	copy(auxiliaryArray, array)
	mergeSortHelper(array, 0, len(array)-1, auxiliaryArray)
	return array
}

func mergeSortHelper(mainArray []int, startIdx, endIdx int, auxiliaryArray []int) {
	if startIdx == endIdx {
		return
	}
	middleIdx := (startIdx + endIdx) / 2
	// sort left auxiliary array
	mergeSortHelper(auxiliaryArray, startIdx, middleIdx, mainArray)
	// sort right auxiliary array
	mergeSortHelper(auxiliaryArray, middleIdx+1, endIdx, mainArray)
	// then do merge them back to the main array
	doMerge(mainArray, startIdx, middleIdx, endIdx, auxiliaryArray)
}

func doMerge(mainArray []int, startIdx, middleIdx, endIdx int, auxiliaryArray []int) {
	k := startIdx
	i := startIdx
	j := middleIdx + 1
	for i <= middleIdx && j <= endIdx {
		if auxiliaryArray[i] <= auxiliaryArray[j] {
			mainArray[k] = auxiliaryArray[i]
			i++
		} else {
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
}

// MergeSort1
// Best: O(nlog(n)) time | O(nlog(n)) space
// Average: O(nlog(n)) time | O(nlog(n)) space
// Worst: O(nlog(n)) time | O(nlog(n)) space
func MergeSort1(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middleIndex := len(array) / 2
	leftHalf := MergeSort1(array[:middleIndex])
	rightHalf := MergeSort1(array[middleIndex:])
	return mergeSortedArrays(leftHalf, rightHalf)
}


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

func mergeSort(array []int) []int {
	return sort(0, len(array)-1, array)
}

func sort(lo, hi int, array []int) []int {
	if len(array) == 0 {
		return []int{}
	}

	mid := (lo + hi) / 2
	if lo == hi {
		return array[lo : hi+1]
	}

	left := sort(lo, mid, array)
	right := sort(mid+1, hi, array)
	return mergeSortedArrays(left, right)
}

func mergeSortedArrays(leftHalf, rightHalf []int) []int {
	sortedArray := make([]int, len(leftHalf)+len(rightHalf))
	k, i, j := 0, 0, 0
	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] <= rightHalf[j] {
			sortedArray[k] = leftHalf[i]
			i++
		} else {
			sortedArray[k] = rightHalf[j]
			j++
		}
		k++
	}
	for i < len(leftHalf) {
		sortedArray[k] = leftHalf[i]
		i++
		k++
	}
	for j < len(rightHalf) {
		sortedArray[k] = rightHalf[j]
		j++
		k++
	}
	return sortedArray
}
