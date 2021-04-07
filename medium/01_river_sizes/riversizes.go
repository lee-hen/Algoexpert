package river_sizes

// 0   2
// 0 1 2   4
//   1 2   4
// 0 1   3 4
// 0   2 3

//0 [0, 1 ,3, 4]
//1 [1, 2, 3]
//2 [0, 1, 2, 4]
//3 [3, 4]
//4 [1, 2, 3]


//{1, 0, 0, 1, 0},
//{1, 0, 1, 0, 0},
//{0, 0, 1, 0, 1},
//{1, 0, 1, 0, 1},
//{1, 0, 1, 1, 0},

// 1 0 1
// 1 0 1 0
// 1 1 1 1

//{1, 0, 0, 0, 0, 0, 1, 0},
//{0, 0, 0, 0, 0, 0, 1, 0},
//{1, 0, 1, 0, 1, 1, 1, 1},
//{1, 1, 0, 0, 0, 0, 1, 0},
//{1, 1, 1, 1, 1, 1, 1, 0},


//0 0
//1 0
//2 0
//3 0
//4 0

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

	var idxRow int
	var idxCol int
	return traverseMatrix(matrix, riverCoordinates, idxRow, idxCol)
}

func traverseMatrix(matrix [][]int, riverCoordinates []map[int]struct{}, idxRow int, idxCol int) (riverSizes []int) {

	for ; idxRow < len(matrix); idxRow++  {
		matrixRow := matrix[idxRow]
		for ; idxCol < len(matrixRow); idxCol++ {
			matrixCol := matrix[idxRow][idxCol]
			var above, below, right int
			if idxRow > 0 {
				above = matrix[idxRow - 1][idxCol]
			}
			if idxRow < len(matrix) - 1 {
				below = matrix[idxRow + 1][idxCol]
			}

			if idxCol < len(matrixRow) - 1 {
				right = matrixRow[idxCol + 1]
			}

			if matrixCol == 1 {
				if len(riverSizes) == 0 {
					riverSizes = append(riverSizes, 0)
				}
				riverNo := len(riverSizes) - 1
				riverSizes[riverNo]++
				delete(riverCoordinates[idxRow], idxCol)

				if below == 0 && right == 0 {
					riverSizes = append(riverSizes, 0)
				}

				if above == 1  {
					riverSizes[riverNo] += traverseBackRowMatrix(matrix, riverCoordinates, idxRow-1, idxCol)
				}

			}

			if right == 0 {
				idxCol = 0
				break
			}
		}

		//fmt.Println("^-^^^^^^^^^^^^^^-0-000---------------------------")
		//fmt.Println(riverSizes)
		//fmt.Println("^-^^^^^^^^^^^^^^-0-000---------------------------")
		//fmt.Println(riverSizes)
		//fmt.Println(idxRow)
		//fmt.Println(idxCol)

		//if idxRow == len(matrix) - 1  {
		//	idxRow = 0
		//	idxCol++
		//
		//	if len(riverCoordinates[idxRow]) == 0 {
		//		idxRow++
		//	}
		//
		//	if ok {
		//		riverSizes = traverseMatrix(matrix, riverCoordinates, idxRow, idxCol)
		//		riverSizes = append(riverSizes, riverSizes...)
		//	}
		//
		//
		//}

	}


	return

}

func traverseBackRowMatrix(matrix [][]int, riverCoordinates []map[int]struct{}, idxRow int, idxCol int) (riverSize int) {
	if _, ok := riverCoordinates[idxRow][idxCol]; !ok || idxRow < 0 || idxCol > len(matrix[idxRow])  {
		return
	}

	if matrix[idxRow][idxCol] == 1  {
		riverSize++
		delete(riverCoordinates[idxRow], idxCol)
	}


	above := matrix[idxRow-1][idxCol]
	left := matrix[idxRow][idxCol-1]
	right := matrix[idxRow][idxCol+1]
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
		riverSize += traverseBackRowMatrix(matrix, riverCoordinates, idxRow, idxCol)
	}

	return
}

