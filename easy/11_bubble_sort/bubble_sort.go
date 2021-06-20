package bubble_sort

func BubbleSort(array []int) []int {
	isSorted := false
	counter := 0

	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1-counter; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}
		counter++
	}
	return array
}

func bubbleSort(array []int) []int {
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := 0; i < len(array)-1; i++ {
			if array[i] > array[i+1] {
				array[i], array[i+1] = array[i+1], array[i]
				isSorted = false
			}
		}

	}
	return array
}
