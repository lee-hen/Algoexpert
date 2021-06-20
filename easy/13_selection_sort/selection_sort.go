package selection_sort

func SelectionSort(array []int) []int {

	currentIndex := 0
	for currentIndex < len(array)-1 {
		smallestIndex := currentIndex

		for i := currentIndex + 1; i < len(array); i++ {
			if array[smallestIndex] > array[i] {
				smallestIndex = i
			}
		}
		array[smallestIndex], array[currentIndex] = array[currentIndex], array[smallestIndex]

		currentIndex++
	}
	return array
}

// My solution
func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		k := i
		for j := i + 1; j < len(array); j++ {
			if array[k] > array[j] {
				k = i
			}
		}

		array[i], array[k] = array[k], array[i]
	}

	return array
}
