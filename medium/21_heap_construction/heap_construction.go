package heap_construction
// important question

type Heap struct {
	comp   HeapFunc
	values []int
}

type HeapFunc func(int, int) bool

var MinHeapFunc = func(a, b int) bool { return a < b }
var MaxHeapFunc = func(a, b int) bool { return a > b }

func NewHeap(fn HeapFunc) *Heap {
	return &Heap{
		comp:   fn,
		values: []int{},
	}
}

func (h *Heap) Length() int {
	return len(h.values)
}

func (h *Heap) Peek() int {
	if len(h.values) == 0 {
		return -1
	}
	return h.values[0]
}

func (h *Heap) Insert(value int) {
	h.values = append(h.values, value)
	h.siftUp()
}

func (h *Heap) Update(value int) {
	h.values[0] = value
	h.siftDown()
}

func (h *Heap) Remove() int {
	l := h.Length()
	h.swap(0, l-1)
	peeked := h.values[l-1]
	h.values = h.values[0 : l-1]
	h.siftDown()
	return peeked
}

func (h *Heap) siftUp() {
	currentIndex := h.Length() - 1
	parentIndex := (currentIndex - 1) / 2

	for currentIndex > 0 {
		current, parent := h.values[currentIndex], h.values[parentIndex]

		if !h.comp(current, parent) {
			return
		}

		h.swap(currentIndex, parentIndex)
		currentIndex = parentIndex
		parentIndex = (currentIndex - 1) / 2
	}
}

func (h *Heap) siftDown() {
	currentIndex := 0
	childOneIdx := currentIndex*2 + 1
	endIndex := h.Length() - 1

	for childOneIdx <= endIndex {
		childTwoIdx := -1

		if currentIndex*2+2 <= endIndex {
			childTwoIdx = currentIndex*2 + 2
		}

		indexToSwap := childOneIdx
		if childTwoIdx > -1 && h.comp(h.values[childTwoIdx], h.values[childOneIdx]) {
			indexToSwap = childTwoIdx
		}

		if !h.comp(h.values[indexToSwap], h.values[currentIndex]) {
			return
		}

		h.swap(currentIndex, indexToSwap)
		currentIndex = indexToSwap
		childOneIdx = currentIndex*2 + 1
	}
}

func (h *Heap) swap(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
