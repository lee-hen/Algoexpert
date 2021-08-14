package sort_k_sorted_array
// important question

import (
	heap "github.com/lee-hen/Algoexpert/medium/21_heap_construction"
)

// SortKSortedArray
// O(nlog(k)) time | O(k) space - where n is the number
// of elements in the array and k is how far away elements
// are from their sorted position
func SortKSortedArray(array []int, k int) []int {
	if len(array) == 0 || k == 0 {
		return array
	}

	heapArray := make([]int, min(k+1, len(array)))
	copy(heapArray, array[0:min(k+1, len(array))])
	minHeaps := heap.NewMinHeap(heapArray)

	idxToInsert := 0
	for idx := k + 1; idx < len(array); idx++ {
		minElement := minHeaps.Remove()

		array[idxToInsert] = minElement
		idxToInsert++

		minHeaps.Insert(array[idx])
	}

	for !minHeaps.IsEmpty() {
		minElement := minHeaps.Remove()
		array[idxToInsert] = minElement
		idxToInsert++
	}

	return array
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// k = 3
// n*k + k*K time  o(1) space
func sortKSortedArray(array []int, k int) []int {
	if len(array) == 0 {
		return array
	}

	for i := 0; i+k < len(array); i++ {
		heap.NewMinHeap(array[i : i+k+1])
	}

	kth := len(array) - k
	if k > len(array) {
		kth = k - len(array)
	}

	for i := kth; i < len(array); i++ {
		heap.NewMinHeap(array[i:])
	}

	return array
}

func sortKSortedArray1(array []int, k int) []int {
	if len(array) == 0 {
		return array
	}

	if k > len(array) {
		k = k % len(array)
		if k == 1 {
			if array[len(array)-1] <= array[0] {
				return rotation(array, 1)
			} else {
				return rotation(array, len(array)-1)
			}
		}
	}

	for i := 0; i+k < len(array); i++ {
		heap.NewMinHeap(array[i : i+k+1])
	}

	for i := len(array) - k; i < len(array); i++ {
		heap.NewMinHeap(array[i:])
	}

	return array
}

func rotation(array []int, k int) []int {
	result := make([]int, len(array), len(array))
	for idx := range array {
		jumpIdx := (idx + k) % len(array)
		result[jumpIdx] = array[idx]
	}
	return result
}
