package insertion_sort

func InsertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

// insertionSort1
// easy to understand
// best case does compare n-1 compares
func insertionSort1(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i; j > 0; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
			}
		}
	}

	return array
}

// My solution
// well best case is not valuable
func insertionSort2(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < i+1; j++ {
			if array[j] > array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array

}
