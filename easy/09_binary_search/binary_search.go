package binary_search
// important question

func BinarySearch(array []int, target int) int {
	return binarySearch(array, target, 0, len(array)-1)
}

func binarySearch(array []int, target int, left int, right int) int {
	for left <= right {
		middle := (left + right) / 2
		if target == array[middle] {
			return middle
		} else if target > array[middle] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func binarySearch1(array []int, target int, left int, right int) int {
	middle := (left + right) / 2
	if target == array[middle] {
		return middle
	} else if middle < right && target > array[middle] {
		return binarySearch1(array, target, middle+1, right)
	} else if middle > left && target < array[middle] {
		return binarySearch1(array, target, left, middle-1)
	}

	return -1
}
