package river_sizes

// my corrected solution
func RiverSizes(matrix [][]int) []int {
	riverCoordinates := make([]map[int]struct{}, len(matrix))
	for idxRow, matrixRow := range matrix {
		var riverCoordinate = make(map[int]struct{})
		for idxCol, matrixCol := range matrixRow {
			if matrixCol == 1 {
				riverCoordinate[idxCol] = struct{}{}
				riverCoordinates[idxRow]= riverCoordinate
			}
		}
	}
	riverSizes := traverseMatrix(matrix, riverCoordinates)

	if riverSizes == nil {
		return []int{}
	}
	return riverSizes
}

func traverseMatrix(matrix [][]int, riverCoordinates []map[int]struct{}) (riverSizes []int) {
	for idxR, riverCoordinate := range riverCoordinates {
		for idxC := range  riverCoordinate {
			var riverSize int
			riverSize += traverseRiverSize(matrix, riverCoordinates, idxR, idxC)
			riverSizes = append(riverSizes, riverSize)
		}
	}

	return
}

func traverseRiverSize(matrix [][]int, riverCoordinates []map[int]struct{}, idxRow int, idxCol int) (riverSize int) {

	var above, below, left, right int

	if idxRow > 0 {
		above = matrix[idxRow-1][idxCol]
	}

	if idxRow < len(matrix)-1 {
		below = matrix[idxRow+1][idxCol]
	}

	if idxCol > 0 {
		left = matrix[idxRow][idxCol-1]
	}

	if idxCol < len(matrix[idxRow])-1 {
		right = matrix[idxRow][idxCol+1]
	}

	if matrix[idxRow][idxCol] == 1  {
		if !isVisited(riverCoordinates, idxRow, idxCol) {
			riverSize++
			delete(riverCoordinates[idxRow], idxCol)
		} else {
			return 0
		}
	}

	if below == 1 {
		riverSize += traverseRiverSize(matrix, riverCoordinates, idxRow+1, idxCol)
	}

	if right == 1 {
		riverSize += traverseRiverSize(matrix, riverCoordinates, idxRow, idxCol+1)
	}

	if above == 1 {
		riverSize += traverseRiverSize(matrix, riverCoordinates, idxRow-1, idxCol)
	}

	if  left == 1  {
		riverSize += traverseRiverSize(matrix, riverCoordinates, idxRow, idxCol-1)
	}

	return
}

func isVisited(riverCoordinates []map[int]struct{}, idxRow int, idxCol int) bool {
	if _, ok := riverCoordinates[idxRow][idxCol]; ok  {
		return false
	}
	return true
}


// my uncorrected solution
func riverSizes(matrix [][]int) []int {
	riverCoordinates := make([]map[int]struct{}, len(matrix))
	for idxRow, matrixRow := range matrix {
		var riverCoordinate = make(map[int]struct{})
		for idxCol, matrixCol := range matrixRow {
			if matrixCol == 1 {
				riverCoordinate[idxCol] = struct{}{}
				riverCoordinates[idxRow]= riverCoordinate
			}
		}
	}
	returnValue := traverseMatrix1(matrix, riverCoordinates, 0, 0)
	riverSizes := make([]int, 0)
	for _, riverSize := range returnValue {
		if riverSize != 0 {
			riverSizes = append(riverSizes, riverSize)
		}

	}
	return returnValue
}

func traverseMatrix1(matrix [][]int, riverCoordinates []map[int]struct{}, idxRow int, initIdxCol int) (riverSizes []int) {

	out:
	for ; idxRow < len(matrix); idxRow++  {
		matrixRow := matrix[idxRow]
		for idxCol := initIdxCol; idxCol < len(matrixRow); idxCol++ {

			matrixCol := matrix[idxRow][idxCol]
			var above, below, left, right int
			if idxRow > 0 {
				above = matrix[idxRow - 1][idxCol]
			}
			if idxRow < len(matrix) - 1 {
				below = matrix[idxRow + 1][idxCol]
			}

			if idxCol < len(matrixRow) - 1 {
				right = matrixRow[idxCol + 1]
			}

			if idxCol > 0 {
				left = matrix[idxRow][idxCol-1]
			}

			if matrixCol == 1 {
				if len(riverSizes) == 0 {
					riverSizes = append(riverSizes, 0)
				}

				if _, ok := riverCoordinates[idxRow][idxCol]; !ok  {
					return
				}

				riverNo := len(riverSizes) - 1
				riverSizes[riverNo]++
				delete(riverCoordinates[idxRow], idxCol)

				if below == 0 && right == 0  {
					riverSizes = append(riverSizes, 0)
						break out
				}


				if above == 1  {
					riverSizes[riverNo] += traverseBackMatrix1(matrix, riverCoordinates, idxRow-1, idxCol)
				}

				if left == 1 {
					riverSizes[riverNo] += traverseBackMatrix1(matrix, riverCoordinates, idxRow, idxCol-1)
				}
			}

			if right == 0 {
				idxCol = initIdxCol
				break
			}
		}
	}


	for idxR, riverCoordinate := range riverCoordinates {
		for idxC := range  riverCoordinate {
			x := traverseMatrix1(matrix, riverCoordinates, idxR, idxC)
			riverSizes = append(riverSizes, x...)
		}
	}

	return
}

func traverseBackMatrix1(matrix [][]int, riverCoordinates []map[int]struct{}, idxRow int, idxCol int) (riverSize int) {
	if _, ok := riverCoordinates[idxRow][idxCol]; !ok || idxRow < 0  {
		return
	}

	if matrix[idxRow][idxCol] == 1  {
		riverSize++
		delete(riverCoordinates[idxRow], idxCol)
	}

	var above, left, right int

	if idxRow > 1 {
		above = matrix[idxRow-1][idxCol]
	}

	if idxCol > 0 {
		left = matrix[idxRow][idxCol-1]
	}

	if idxCol < len(matrix[idxRow]) - 1 {
		right = matrix[idxRow][idxCol+1]
	}

	if above == 1 {
		idxRow--
	}

	if  left == 1  {
		idxCol--
	}

	if right == 1 {
		idxCol++
	}

	if above == 1 || left == 1 || right == 1 {
		riverSize += traverseBackMatrix1(matrix, riverCoordinates, idxRow, idxCol)
	}

	return
}

