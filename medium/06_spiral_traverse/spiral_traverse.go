package spiral_traverse
// important question

// SpiralTraverse
// better solution
func SpiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	result := make([]int, 0)
	spiralFill(array, 0, len(array)-1, 0, len(array[0])-1, &result)

	return result
}

func spiralFill(array [][]int, startRow, endRow, startCol, endCol int, result *[]int) {
	if startRow > endRow || startCol > endCol {

		return

	}
	for col := startCol; col <= endCol; col++ {
		*result = append(*result, array[startRow][col])
	}
	for row := startRow + 1; row <= endRow; row++ {
		*result = append(*result, array[row][endCol])
	}
	for col := endCol - 1; col >= startCol; col-- {
		// Handle the edge case when there's a single row
		// in the middle of the matrix. In this case, we don't
		// want to double-count the values in this row, which
		// we've already counted in the first for loop above.
		// See Test Case 8 for an example of this edge case.
		if startRow == endRow {
			break
		}
		*result = append(*result, array[endRow][col])
	}
	for row := endRow - 1; row >= startRow+1; row-- {
		// Handle the edge case when there's a single column
		// in the middle of the matrix. In this case, we don't
		// want to double-count the values in this column, which
		// we've already counted in the second for loop above.
		// See Test Case 9 for an example of this edge case.
		if startCol == endCol {
			break
		}
		*result = append(*result, array[row][startCol])
	}
	spiralFill(array, startRow+1, endRow-1, startCol+1, endCol-1, result)
}

func spiralTraverse1(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	result := make([]int, 0)
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1
	for startRow <= endRow && startCol <= endCol {
		for col := startCol; col <= endCol; col++ {
			result = append(result, array[startRow][col])
		}
		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, array[row][endCol])
		}
		for col := endCol - 1; col >= startCol; col-- {
			// Handle the edge case when there's a single row
			// in the middle of the matrix. In this case, we don't
			// want to double-count the values in this row, which
			// we've already counted in the first for loop above.
			// See Test Case 8 for an example of this edge case.
			if startRow == endRow {
				break
			}
			result = append(result, array[endRow][col])
		}
		for row := endRow - 1; row > startRow; row-- {
			// Handle the edge case when there's a single column
			// in the middle of the matrix. In this case, we don't
			// want to double-count the values in this column, which
			// we've already counted in the second for loop above.
			// See Test Case 9 for an example of this edge case.
			if startCol == endCol {
				break
			}
			result = append(result, array[row][startCol])
		}
		startRow++
		endRow--
		startCol++
		endCol--
	}
	return result
}


// my solution
func spiralTraverse(array [][]int) []int {
	var rowSize, colSize int
	rowSize = len(array)
	colSize = len(array[0])

	result := make([]int, rowSize * colSize, rowSize * colSize)
	traverseMatrix(0, 0, 0, 0, 0, rowSize, colSize, result, array)
	return result
}

func traverseMatrix(idxRow, idxCol, i, rowStart, colStart, rowSize, colSize int, result []int, matrix [][]int)  {
	if i > len(result) -1 {
		return
	}
	if idxCol == colSize-1 && idxRow == rowStart && i == len(result) -1 {
		result[i] = matrix[idxRow][idxCol]
		return
	}

	if idxCol < colSize-1 && idxRow == rowStart {
		result[i] = matrix[idxRow][idxCol]
		// traverseRight
		traverseMatrix(idxRow, idxCol+1, i+1, rowStart, colStart, rowSize, colSize, result, matrix)
	}

	if idxCol == colSize-1 && idxRow < rowSize-1 {
		result[i] = matrix[idxRow][idxCol]
		// traverseBelow
		traverseMatrix(idxRow+1, idxCol, i+1, rowStart, colStart, rowSize, colSize, result, matrix)
	}

	if 	idxCol > colStart && idxRow == rowSize-1 {
		result[i] = matrix[idxRow][idxCol]
		if rowSize != 1 {
			// traverseLeft
			traverseMatrix(idxRow, idxCol-1, i+1, rowStart, colStart, rowSize, colSize, result, matrix)
		}
	}
	if idxRow > rowStart && idxCol == colStart {
		result[i] = matrix[idxRow][idxCol]
		if idxRow-1 == rowStart {
			// traverseInside
			traverseMatrix(idxRow, idxCol+1, i+1, rowStart+1, colStart+1, rowSize-1, colSize-1, result, matrix)
			return
		}
		if colSize != 1 {
			//traverseAbove
			traverseMatrix(idxRow-1, idxCol, i+1, rowStart, colStart, rowSize, colSize, result, matrix)
		}
	}
}
