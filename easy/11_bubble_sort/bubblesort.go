package bubblesort

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
