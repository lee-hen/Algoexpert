package first_duplicate_value
// important question

// FirstDuplicateValue
// mutate the input array
func FirstDuplicateValue(array []int) int {
	for _, value := range array {
		absValue := abs(value)
		if array[absValue-1] < 0 {
			return absValue
		}
		array[absValue-1] *= -1
	}
	return -1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// FirstDuplicateValue
// does not mutate the input array
func firstDuplicateValue(arr []int) int {
	idxArr := make([]int, len(arr)+1)
	for idx := 0; idx < len(arr); idx++ {
		idxArr[arr[idx]] = -1
	}

	for idx := 0; idx < len(arr); idx++ {
		if idxArr[arr[idx]] >= 0 {
			return arr[idx]
		}

		idxArr[arr[idx]] = idx
	}

	return -1
}
