package merge_sorted_arrays
// important question

import (
	heapSort "github.com/lee-hen/Algoexpert/hard/39_heap_sort"

	"container/heap"
)

type Item struct {
	ArrayIdx   int
	ElementIdx int
	Num        int
}

type ItemHeap []Item

func (h ItemHeap) Len() int           { return len(h) }
func (h ItemHeap) Less(i, j int) bool { return h[i].Num < h[j].Num }
func (h ItemHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }


func (h *ItemHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Item))
}

func (h *ItemHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MergeSortedArrays
// O(nlog(k) + k) time | O(n + k) space - where n is the total
// number of array elements and k is the number of arrays
func MergeSortedArrays(arrays [][]int) []int {
	sortedList := make([]int, 0)
	smallestItems := ItemHeap{}
	for arrayIdx := 0; arrayIdx < len(arrays); arrayIdx++ {
		smallestItems = append(smallestItems, Item{
			ArrayIdx:   arrayIdx,
			ElementIdx: 0,
			Num:        arrays[arrayIdx][0],
		})
	}

	heap.Init(&smallestItems)
	for len(smallestItems) != 0 {
		smallestItem := heap.Pop(&smallestItems)

		sortedList = append(sortedList, smallestItem.(Item).Num)
		if smallestItem.(Item).ElementIdx == len(arrays[smallestItem.(Item).ArrayIdx])-1 {
			continue
		}

		heap.Push(&smallestItems, Item{
			ArrayIdx:   smallestItem.(Item).ArrayIdx,
			ElementIdx: smallestItem.(Item).ElementIdx + 1,
			Num:        arrays[smallestItem.(Item).ArrayIdx][smallestItem.(Item).ElementIdx+1],
		})
	}
	return sortedList
}

func mergeSortedArrays(arrays [][]int) []int {
	result := make([]int, 0)
	for _, arr := range arrays {
		for _, number := range arr {
			result = append(result, number)
		}
	}
	heapSort.HeapSort(result)
	return result
}
