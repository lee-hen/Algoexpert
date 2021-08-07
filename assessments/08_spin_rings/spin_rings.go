package spin_rings

// prev 2
// curr 3

// [1, 2, 3, 4],
// [5, 6, 7, 8],
// [9, 10, 11, 12],
// [13, 14, 15, 16]

// [5, 1, 2, 3],
// [9, 10, 6, 4],
// [13, 11, 7, 8],
// [14, 15, 16, 12]

func SpinRings(array [][]int) {
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	for startRow <= endRow && startCol <= endCol {
		if startCol == endCol && startRow == endRow {
			break
		}

		var prev = array[startRow+1][startCol]

		for col := startCol; col <= endCol; col++ {
			curr := array[startRow][col]
			array[startRow][col] = prev
			prev = curr
		}

		for row := startRow + 1; row <= endRow; row++ {
			curr := array[row][endCol]
			array[row][endCol] = prev
			prev = curr

		}
		for col := endCol - 1; col >= startCol; col-- {
			if startRow == endRow {
				break
			}
			curr := array[endRow][col]
			array[endRow][col] = prev
			prev = curr
		}
		for row := endRow - 1; row > startRow; row-- {
			if startCol == endCol {
				break
			}

			curr := array[row][startCol]
			array[row][startCol] = prev
			prev = curr
		}
		startRow++
		endRow--
		startCol++
		endCol--
	}
}
