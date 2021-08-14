package search_for_range
// important question

// 0  1  2   3    4  5    6   7  8    9  10  11  12
// 0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73
// 45
// 4, 9

// SearchForRange
// O(log(n)) time | O(log(n)) space
func SearchForRange(array []int, target int) []int {
	finalRange := []int{-1, -1}
	alteredBinarySearch(array, target, 0, len(array)-1, finalRange, true)
	alteredBinarySearch(array, target, 0, len(array)-1, finalRange, false)
	return finalRange
}

func alteredBinarySearch(array []int, target, left, right int, finalRange []int, goLeft bool) {
	if left > right {
		return
	}

	mid := (left + right) / 2
	if array[mid] < target {
		alteredBinarySearch(array, target, mid+1, right, finalRange, goLeft)
	} else if array[mid] > target {
		alteredBinarySearch(array, target, left, mid-1, finalRange, goLeft)
	} else {
		if goLeft {
			if mid == 0 || array[mid-1] != target {
				finalRange[0] = mid
			} else {
				alteredBinarySearch(array, target, left, mid-1, finalRange, goLeft)
			}
		} else {
			if mid == len(array)-1 || array[mid+1] != target {
				finalRange[1] = mid
			} else {
				alteredBinarySearch(array, target, mid+1, right, finalRange, goLeft)
			}
		}
	}
}

// SearchForRange1
// O(log(n)) time | O(1) space
func SearchForRange1(array []int, target int) []int {
	finalRange := []int{-1, -1}
	alteredBinarySearch1(array, target, 0, len(array)-1, finalRange, true)
	alteredBinarySearch1(array, target, 0, len(array)-1, finalRange, false)
	return finalRange
}
func alteredBinarySearch1(array []int, target, left, right int, finalRange []int, goLeft bool) {
	for left <= right {
		mid := (left + right) / 2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		} else {
			if goLeft {
				if mid == 0 || array[mid-1] != target {
					finalRange[0] = mid
					return
				} else {
					right = mid - 1
				}
			} else {
				if mid == len(array)-1 || array[mid+1] != target {
					finalRange[1] = mid
					return
				} else {
					left = mid + 1
				}
			}
		}
	}
}

func searchForRange(array []int, target int) []int {
	result := make([]int, 0)
	binarySearch(&result, array, target, 0, len(array)-1)

	if len(result) == 0 {
		return []int{-1, -1}
	}

	if len(result) == 1 {
		result = append(result, result[0])
	}

	return result
}

func binarySearch(result *[]int, array []int, target, left, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2
	if target == array[middle] {

		outOfRangeIdx := -1
		if len(*result) == 0 {
			outOfRangeIdx = binarySearch(result, array, target, left, middle-1)

			if outOfRangeIdx == -1 {
				*result = append(*result, middle)
			}
		}

		if outOfRangeIdx != -1 {
			if target == array[right] {
				*result = append(*result, right)
			} else {
				*result = append(*result, binarySearch(result, array, target, middle+1, right))
				if (*result)[1] == -1 {
					(*result)[1] = middle
				}
			}
		}

		return middle
	}

	if middle < right && target > array[middle] {
		return binarySearch(result, array, target, middle+1, right)
	}

	if middle > left && target < array[middle] {
		return binarySearch(result, array, target, left, middle-1)
	}

	return -1
}
