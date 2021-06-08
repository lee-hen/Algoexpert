package heap_construction

import (
	"math"
)

//			  48
//		12                       24
//	7           8           -5            24
//391   24   56    2      6     8      41

//			  -5
//		2                       48
//	7           12           24           24
//391   24   56    8      6     8      41

//			  -5
//		2                         6
//	7           8           8           24
//391   24   56    12    24     48   41
func (h *MinHeap) buildHeap() {
	minHeap := *h
	for i := 1; i < len(minHeap); i++ {
		parentIdx := (i - 1) / 2
		if parentIdx < 0 {
			break
		}

		if minHeap[parentIdx] >= minHeap[i] {
			h.siftUp()
		}
	}
}

// left = 2i+1
// right = 2i+2
//  0  1   2   3  4   5  6   7    8   9  10  11  12  13
// 48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2,  6,  8, 41
func (h *MinHeap) siftDown(currentIdx, endIdx int) {
	minHeap := *h
	for {
		left := 2*currentIdx + 1
		right := 2*currentIdx + 2
		if left > endIdx || right > endIdx {
			break
		}
		minIdx := right
		if minHeap[left] < minHeap[right] {
			minIdx = left
		}
		if minHeap[currentIdx] > minHeap[minIdx] {
			minHeap[currentIdx], minHeap[minIdx] = minHeap[minIdx], minHeap[currentIdx]
		} else {
			break
		}

		currentIdx = minIdx
	}
}

// (endIndex-2)/2 == (endIndex-1)/2 has right node
// because left is always exist
// left parent is same as right parent
// parent = (endIndex-1)/2
//  0  1   2   3  4   5  6   7    8   9  10  11  12  13
//  48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41
func (h *MinHeap) siftUp() {
	left, right := len(*h)-1, len(*h)-1
	parentIdx := (left - 1) / 2
	if (right-2)/2 == parentIdx {
		right = left
		left = right - 1
	}

	minHeap := *h
	for parentIdx >= 0 {
		currentIdx := right
		if minHeap[left] < minHeap[currentIdx] {
			currentIdx = left
		}

		if minHeap[currentIdx] < minHeap[parentIdx] {
			minHeap[parentIdx], minHeap[currentIdx] = minHeap[currentIdx], minHeap[parentIdx]
		}

		right = left - 1
		left = right - 1
		parentIdx = (left - 1) / 2
	}
}

func (h MinHeap) peek() int {
	return h[0]
}

// 		    	      2
// 		    7                        6
// 	    24           8            8         24
// 391      41   56     12    48    24   41
// remove
// 2 7 6 24 8 8 24 391 41 56 12 48 24 41
func (h *MinHeap) remove() int {
	minHeap := *h
	min := math.MinInt32
	if len(minHeap) > 0 {
		min = h.Peek()
		minHeap[0] = minHeap[len(minHeap)-1]
		minHeap.siftDown(0, len(minHeap)-2)
		minHeap = minHeap[:len(minHeap)-1]
	}
	return min
}

func (h *MinHeap) insert(value int) {
	*h = append(*h, value)
	h.siftUp()
}
