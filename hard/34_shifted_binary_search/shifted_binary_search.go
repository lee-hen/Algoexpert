package shifted_binary_search
// important question

// ShiftedBinarySearch
// O(log(n)) time | O(1) space
func ShiftedBinarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		middle := (left + right) / 2

		if target == array[middle] {
			return middle
		}

		if array[left] <= array[middle] {
			if target < array[middle] && target >= array[left] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target > array[middle] && target <= array[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return -1
}


// 73
// 0   1   2   3   4   5  6  7   8   9
// 71, 72, 73, 0, 1, 21, 33, 37, 45, 61

// 45
// 0   1   2   3   4   5  6  7   8   9
// 45, 61, 71, 72, 73, 0, 1, 21, 33, 37

// shiftedBinarySearch
// my solution
func shiftedBinarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	for left <= right {
		middle := (left + right) / 2

		if array[middle] == target {
			return middle
		}

		if array[left] > array[right] {

			// 71, 72, 73, 0
			// left 71  middle 72 right 0
			// target 73
			if array[middle] > array[right] && target > array[middle] ||
				array[left] > target { // 45, 61, 21, 33, 37 target 33
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			if target > array[middle] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		}
	}

	return -1
}

func ShiftedBinarySearch1(array []int, target int) int {
	return helper(array, target, 0, len(array)-1)
}

func helper(array []int, target int, left int, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2
	potentialMatch := array[middle]

	leftNum, rightNum := array[left], array[right]
	newLeft, newRight := middle+1, right

	if target == potentialMatch {
		return middle
	}

	if leftNum <= potentialMatch {
		if target < potentialMatch && target >= leftNum {
			newLeft, newRight = left, middle-1
		}
	} else {
		if !(target > potentialMatch && target <= rightNum) {
			newLeft, newRight = left, middle-1
		}
	}

	return helper(array, target, newLeft, newRight)
}
