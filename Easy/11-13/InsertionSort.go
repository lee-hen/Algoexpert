package main

// My solution
func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		for j := 0; j < i+1; j++ {
			if array[j] > array[i] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}
