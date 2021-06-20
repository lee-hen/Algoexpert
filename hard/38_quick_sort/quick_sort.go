package quick_sort

func QuickSort(array []int) []int {
	sortHelper(0, len(array)-1, array)
	return array
}

func sortHelper(startIdx, endIdx int, array []int) {
	if startIdx >= endIdx {
		return
	}

	pivotElement := array[startIdx]
	left, middle, right := startIdx, startIdx+1, endIdx

	for middle <= right && left <= middle {
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

func quickSort(array []int) []int {
	sort(0, len(array)-1, array)
	return array
}

func sort(lo, hi int, array []int) {
	if lo >= hi {
		return
	}

	j := partition(lo, hi, array)
	sort(lo, j-1, array)
	sort(j+1, hi, array)
	return
}

func partition(lo, hi int, array []int) int {
	pivot := lo
	i, j := lo+1, hi

	for  {
		for i < hi && array[i] <= array[pivot] {
			i++
		}

		for j > lo && array[pivot] <= array[j] {
			j--
		}

		if i >= j {
			break
		}
		array[i], array[j] = array[j], array[i]
	}

	array[pivot], array[j] = array[j], array[pivot]
	return j
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
