package search_in_sorted_matrix
// important question

// SearchInSortedMatrix
// better solution
func SearchInSortedMatrix(matrix [][]int, target int) []int {
	row := 0
	col := len(matrix[row])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] > target {
			col--
		} else if matrix[row][col] < target {
			row++
		} else {
			return []int{row, col}
		}
	}
	return []int{-1, -1}
}

// SearchInSortedMatrix 44
//{1, 4, 7, 12, 15, 1000},
//{2, 5, 19, 31, 32, 1001},
//{3, 8, 24, 33, 35, 1002},
//{40, 41, 42, 44, 45, 1003},
//{99, 100, 103, 106, 128, 1004},
func searchInSortedMatrix(matrix [][]int, target int) []int {
	for row := range matrix {
		col := binarySearch(target, 0, len(matrix[row])-1, matrix[row])
		if col != -1 {
			return []int{row, col}
		}
	}
	return []int{-1, -1}
}

func binarySearch(target, left, right int, array []int) int {
	middle := (left + right) / 2
	if target == array[middle] {
		return middle
	} else if middle < right && target > array[middle] {
		return binarySearch(target, middle+1, right, array)
	} else if middle > left && target < array[middle] {
		return binarySearch(target, left, middle-1, array)
	}

	return -1
}
