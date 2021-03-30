package insertionsort

func InsertionSort(array []int) []int {
	for i := range array {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

// My solution
func InsertionSort1(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < i+1; j++ {
			if array[j] > array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}
