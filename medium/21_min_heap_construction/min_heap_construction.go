package min_heap_construction

type MinHeap []int

func NewMinHeap(array []int) *MinHeap {
	heap := MinHeap(array)
	ptr := &heap
	ptr.BuildHeap(array)
	return ptr
}

func (h *MinHeap) BuildHeap(array []int) {
}

func (h *MinHeap) siftDown(currentIndex, endIndex int) {
}

func (h *MinHeap) siftUp() {
}

func (h MinHeap) Peek() int {
	return -1
}

func (h *MinHeap) Remove() int {
	return -1
}

func (h *MinHeap) Insert(value int) {
}
