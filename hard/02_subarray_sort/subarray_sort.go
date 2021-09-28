package subarray_sort
// important question

import (
	"math"
)


func SubarraySort(array []int) []int {
	minOutOfOrder, maxOutOfOrder := math.MaxInt32, math.MinInt32
	for i, num := range array {
		if isOutOfOrder(i, num, array) {
			minOutOfOrder = min(minOutOfOrder, num)
			maxOutOfOrder = max(maxOutOfOrder, num)
		}
	}

	if minOutOfOrder == math.MaxInt32 {
		return []int{-1, -1}
	}

	subarrayLeft := 0
	for minOutOfOrder >= array[subarrayLeft] {
		subarrayLeft += 1
	}

	subarrayRight := len(array) - 1
	for maxOutOfOrder <= array[subarrayRight] {
		subarrayRight -= 1
	}

	return []int{subarrayLeft, subarrayRight}
}

func isOutOfOrder(i int, num int, array []int) bool {
	if i == 0 {
		return num > array[i+1]
	}

	if i == len(array)-1 {
		return num < array[i-1]
	}

	return num > array[i+1] || num < array[i-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}


// 0  1  2  3  4   5   6   7   8  9  10  11  12
// 1, 2, 4, 7, 10, 11, 7, 12,  6, 7, 16, 18, 19

// 3, 9
// my solution
func subarraySort(slice []int) []int {
	if slice[0] > slice[len(slice)-1] {
		return []int{0, len(slice)-1}
	}

	minVal, maxVal := math.MaxInt32, math.MinInt32
	leftIdx, rightIdx := -1, -1
	for i, j := 0, len(slice)-1; i < len(slice) && j >= 0; i, j = i+1, j-1 {
		// iが最後、jが最初かつleftIdx=-1, rightIdx=-1の時、slice is already sorted
		if leftIdx == -1 && rightIdx == -1 && i == len(slice)-1 && j == 0 {
			return []int{leftIdx, rightIdx}
		}

		if leftIdx == -1 && slice[i] > slice[i+1] {
			leftIdx = i
		}

		if rightIdx == -1 && slice[j-1] > slice[j] {
			rightIdx = j
		}

		if leftIdx != -1 && i > leftIdx && slice[i] < minVal  {
			minVal = slice[i]
		}

		if rightIdx != -1 && j < rightIdx &&  slice[j] > maxVal {
			maxVal = slice[j]
		}
	}

	var i int
	for ; i < len(slice); i++ {
		if slice[i] > minVal {
			break
		}
	}

	j := len(slice)-1
	for ; j >= 0; j-- {
		if slice[j] < maxVal {
			break
		}
	}

	return []int{i, j}
}
