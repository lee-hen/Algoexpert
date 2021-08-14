package continuous_median
// important question

import heap "github.com/lee-hen/Algoexpert/medium/21_heap_construction"

type ContinuousMedianHandler struct {
	Median float64

	lowerNumbers   *heap.Heap
	greaterNumbers *heap.Heap

	numbers []int
}

func (handler *ContinuousMedianHandler) GetMedian() float64 {
	return handler.Median
}

func NewContinuousMedianHandler() *ContinuousMedianHandler {
	return &ContinuousMedianHandler{
		Median:         0,
		lowerNumbers:   heap.NewHeap(heap.MaxHeapFunc),
		greaterNumbers: heap.NewHeap(heap.MinHeapFunc),
	}
}

// Insert
// O(log(n)) time | O(n) space
func (handler *ContinuousMedianHandler) Insert(number int) {
	if handler.lowerNumbers.Length() == 0 || number < handler.lowerNumbers.Peek() {
		handler.lowerNumbers.Insert(number)
	} else {
		handler.greaterNumbers.Insert(number)
	}

	handler.ReBalanceHeaps()
	handler.UpdateMedian()
}

func (handler *ContinuousMedianHandler) ReBalanceHeaps() {
	if handler.lowerNumbers.Length()-handler.greaterNumbers.Length() == 2 {
		handler.greaterNumbers.Insert(handler.lowerNumbers.Remove())
	} else if handler.greaterNumbers.Length()-handler.lowerNumbers.Length() == 2 {
		handler.lowerNumbers.Insert(handler.greaterNumbers.Remove())
	}
}

func (handler *ContinuousMedianHandler) UpdateMedian() {
	if handler.lowerNumbers.Length() == handler.greaterNumbers.Length() {
		sum := handler.lowerNumbers.Peek() + handler.greaterNumbers.Peek()
		handler.Median = float64(sum) / 2
	} else if handler.lowerNumbers.Length() > handler.greaterNumbers.Length() {
		handler.Median = float64(handler.lowerNumbers.Peek())
	} else {
		handler.Median = float64(handler.greaterNumbers.Peek())
	}
}
