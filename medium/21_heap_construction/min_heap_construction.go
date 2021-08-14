package heap_construction
// important question

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	//ptr.buildHeap()
	return ptr
}

// BuildHeap
// From the last parent index
// Swap down to the child node use SiftDown method
// Loop over all the parent index
// Node that the last parent index - 1 is equal to the previous children's parent index
func (h *MinHeap) BuildHeap(array []int) {
	firstParentIdx := (len(array) - 2) / 2
	for currentIndex := firstParentIdx; currentIndex >= 0; currentIndex-- {
		h.SiftDown(currentIndex, len(array)-1)
	}
}

// SiftDown
// From the target parent index get last child index
// Swap the target parent down to the last min child
// Used for remove value
func (h *MinHeap) SiftDown(parentIdx, endIndex int) {
	childOneIdx := parentIdx*2 + 1
	for childOneIdx <= endIndex {
		childTwoIdx := -1
		if parentIdx*2+2 <= endIndex {
			childTwoIdx = parentIdx*2 + 2
		}
		childIdxToSwap := childOneIdx
		if childTwoIdx > -1 && (*h)[childTwoIdx] < (*h)[childOneIdx] {
			childIdxToSwap = childTwoIdx
		}
		if (*h)[childIdxToSwap] < (*h)[parentIdx] {
			h.swap(parentIdx, childIdxToSwap)
			parentIdx = childIdxToSwap
			childOneIdx = parentIdx*2 + 1
		} else {
			return
		}
	}
}

// SiftUp
// From last index and swap up to the top max parent
// Used for append new value
func (h *MinHeap) SiftUp() {
	currentChildIdx := h.length() - 1
	parentIndex := (currentChildIdx - 1) / 2
	for currentChildIdx > 0 {
		current, parent := (*h)[currentChildIdx], (*h)[parentIndex]
		if current < parent {
			h.swap(currentChildIdx, parentIndex)
			currentChildIdx = parentIndex
			parentIndex = (currentChildIdx - 1) / 2
		} else {
			break
		}
	}
}

func (h MinHeap) Peek() int {
	if len(h) == 0 {
		return -1
	}
	return h[0]
}

func (h *MinHeap) Remove() int {
	l := h.length()
	h.swap(0, l-1)
	peeked := (*h)[l-1]
	*h = (*h)[0 : l-1]
	h.SiftDown(0, l-2)
	return peeked
}

func (h *MinHeap) Insert(value int) {
	*h = append(*h, value)
	h.SiftUp()
}

func (h MinHeap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) length() int {
	return len(h)
}

func (h MinHeap) IsEmpty() bool {
	return h.length() == 0
}
