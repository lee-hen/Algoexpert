package quick_select
// important question

import (
	"math"
)

// QuickSelect
// Best: O(n) time | O(1) space
// Average: O(n) time | O(1) space
// Worst: O(n^2) time | O(1) space
func QuickSelect(array []int, k int) int {
	return helper(array, 0, len(array)-1, k-1)
}

func helper(array []int, start, end int, position int) int {
	for {
		pivot, left, right := start, start+1, end
		for left <= right {
			if array[left] > array[right] && array[right] < array[pivot] {
				swap(left, right, array)
			}

			if array[left] <= array[pivot] {
				left += 1
			}

			if array[right] >= array[pivot] {
				right -= 1
			}
		}
		swap(pivot, right, array)

		if right == position {
			return array[right]
		}

		if right < position {
			start = right + 1
		} else {
			end = right - 1
		}
	}
}

func swap(one, two int, array []int) {
	array[one], array[two] = array[two], array[one]
}

// 3
// 8, 5, 2, 9, 7, 6, 3
// 5

func quickSelect(array []int, k int) int {
	return array[partition(0, len(array)-1, k, array)]
}

func partition(lo, hi, k int, array []int) int {
	if lo > hi {
		return -1
	}

	pivot := lo
	i, j := lo+1, hi

	for  {
		for i < hi && array[i] < array[pivot] {
			i++
		}

		for j > lo && array[j] > array[pivot] {
			j--
		}

		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
	}

	array[pivot], array[j] = array[j], array[pivot]

	// put the pivot to the j, the index of j is the k-1th smallest value
	if j == k-1 {
		return j
	}

	if j < k { // this means the left is sorted and the pivot number is not kth smallest one so try to the right
		return partition(j+1, hi, k, array)
	}

	return partition(lo, j-1, k, array)
}

// bad time and space complexity
func quickSelect1(array []int, k int) int {
	if k > len(array) {
		return -1
	}

	nums := make(map[int]struct{})
	for _, el := range array {
		nums[el] = struct{}{}
	}

	var result int

	minVal := min(array)
	if k == 1 {
		return minVal
	}

	for i := 1; i < k; i++ {
		target := minVal+i
		if _, found := nums[target]; !found {
			k++
		} else {
			result = target
		}
	}

	return result
}

func min(array []int) int {
	minNum := math.MaxInt32
	for _, el := range array {
		if el < minNum {
			minNum = el
		}
	}
	return minNum
}

