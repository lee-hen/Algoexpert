package heap_sort
// important question

//[9 8 6 5 5 2 3]
//[8 5 6 3 5 2 9]
//[6 5 2 3 5 8 9]
//[5 5 2 3 6 8 9]
//[5 3 2 5 6 8 9]
//[3 2 5 5 6 8 9]
//[2 3 5 5 6 8 9]

func HeapSort(array []int) []int {
	for currentIdx := len(array)/2; currentIdx >= 0; currentIdx-- {
		siftDown(currentIdx, len(array)-1, array)
	}
	endIdx := len(array)-1
	for endIdx > 0 {
		array[0], array[endIdx] = array[endIdx], array[0]
		siftDown(0, endIdx-1, array)
		endIdx--
	}
	return array
}

func siftDown(parentIdx, endIdx int, array []int) {
	currentIdx := parentIdx
	childOneIdx := currentIdx*2 + 1

	for childOneIdx <= endIdx {
		childTwoIdx := -1

		if currentIdx*2+2 <= endIdx {
			childTwoIdx = currentIdx*2+2
		}

		idxToSwap := childOneIdx
		if childTwoIdx > -1 && array[childTwoIdx] > array[childOneIdx] {
			idxToSwap = childTwoIdx
		}

		if array[currentIdx] > array[idxToSwap] {
			break
		}

		array[currentIdx], array[idxToSwap] = array[idxToSwap], array[currentIdx]

		currentIdx = idxToSwap
		childOneIdx = currentIdx*2 + 1
	}
}
