package index_equals_value
// important question

// IndexEqualsValue
// O(log(n)) time | O(log(n)) space - where n is the length of the input array
func IndexEqualsValue(array []int) int {
	return indexEqualsValueHelper(array, 0, len(array)-1)
}

func indexEqualsValueHelper(array []int, leftIndex int, rightIndex int) int {
	if leftIndex > rightIndex {
		return -1
	}

	middleIndex := leftIndex + (rightIndex-leftIndex)/2
	middleValue := array[middleIndex]

	if middleValue == middleIndex && middleIndex == 0 {
		return middleIndex
	}

	if middleValue == middleIndex && array[middleIndex-1] < middleIndex-1 {
		return middleIndex
	}

	if middleValue < middleIndex {
		return indexEqualsValueHelper(array, middleIndex+1, rightIndex)
	}

	return indexEqualsValueHelper(array, leftIndex, middleIndex-1)
}


func IndexEqualsValue1(array []int) int {
	leftIndex, rightIndex := 0, len(array) - 1

	for leftIndex <= rightIndex {
		middleIndex := leftIndex + (rightIndex-leftIndex)/2
		middleValue := array[middleIndex]

		if middleValue < middleIndex {
			leftIndex = middleIndex + 1
		} else if middleValue == middleIndex && middleIndex == 0 {
			return middleIndex
		} else if middleValue == middleIndex && array[middleIndex-1] < middleIndex-1 {
			return middleIndex
		} else {
			rightIndex = middleIndex - 1
		}
	}

	return -1
}

// -5, -3, 0, 3, 4, 5, 9
// 3

func indexEqualsValue(array []int) int {
	return binarySearch(array, 0, len(array)-1)
}

func binarySearch(array []int, left, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2
	if middle == array[middle] {
		if idx := binarySearch(array, left, middle-1); idx == -1 {
			return middle
		} else {
			return idx
		}
	}

	if middle < right && middle > array[middle] {
		return binarySearch(array, middle+1, right)
	}

	if middle > left && middle < array[middle] {
		return binarySearch(array, left, middle-1)
	}

	return -1
}

// worst case nlog(n)
func indexEqualsValue1(array []int) int {
	return binarySearch1(array, 0, 0, len(array)-1)
}

func binarySearch1(array []int, target, left, right int) int {
	if target >= len(array) {
		return -1
	}

	middle := (left + right) / 2
	if target == array[middle] && middle == target {
		return middle
	}

	if middle < right && target > array[middle] {
		return binarySearch1(array, target, middle+1, right)
	}

	if middle > left && target < array[middle] {
		return binarySearch1(array, target, left, middle-1)
	}

	return binarySearch1(array, target+1, 0, len(array)-1)
}
