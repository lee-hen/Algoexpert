package quick_sort
// important question

func QuickSort(array []int) []int {
	sort(0, len(array)-1, array)
	return array
}

func sort(lo, hi int, array []int) {
	if lo >= hi {
		return
	}

	pivot := partition(lo, hi, array)
	sort(lo, pivot-1, array)
	sort(pivot+1, hi, array)

	return
}

func partition(pivot, hi int, array []int) int {
	i, j := pivot+1, hi

	for j >= i {
		if array[i] > array[pivot] && array[j] < array[pivot] {
			array[i], array[j] = array[j], array[i]
		}

		if array[i] <= array[pivot] {
			i++
		}

		if array[j] >= array[pivot] {
			j--
		}
	}

	array[pivot], array[j] = array[j], array[pivot]
	return j
}

func quickSort(array []int) []int {
	sortHelper(0, len(array)-1, array)
	return array
}

func sortHelper(startIdx, endIdx int, array []int) {
	if startIdx >= endIdx {
		return
	}

	pivotElement := array[startIdx]
	left, middle, right := startIdx, startIdx+1, endIdx

	for middle <= right {
		if array[middle] < pivotElement {
			array[middle], array[left] = array[left], array[middle]
			middle++
			left++
		} else if array[middle] > pivotElement {
			array[middle], array[right] = array[right], array[middle]
			right--
		} else {
			middle++
		}
	}

	sortHelper(startIdx, left-1, array)
	sortHelper(right+1, endIdx, array)
	return
}

func QuickSort1(array []int) []int {
	return helper(array, 0, len(array)-1)
}

func helper(array []int, start, end int) []int {
	if start >= end {
		return array
	}

	pivot := start
	left := start + 1
	right := end

	for right >= left {
		if array[left] > array[pivot] && array[right] < array[pivot] {
			array[left], array[right] = array[right], array[left]
		}
		if array[left] <= array[pivot] {
			left += 1
		}
		if array[right] >= array[pivot] {
			right -= 1
		}
	}

	array[pivot], array[right] = array[right], array[pivot]

	if right-1-start < end-(right+1) {
		helper(array, start, right-1)
		helper(array, right+1, end)
	} else {
		helper(array, right+1, end)
		helper(array, start, right-1)
	}
	return array
}
