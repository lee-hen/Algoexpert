package repeated_matrix_values

// RepeatedMatrixValues
// O(n) time | O(min(w, h)) space - where n is the total number of elements in
// the matrix and w and h are the width and height of the matrix, respectively
func RepeatedMatrixValues(matrix [][]int) []int {
	valueCounts := initializeCountsOfPotentialValues(matrix)

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0]); col++ {
			value := matrix[row][col]
			correctCountSoFar := row
			checkAndIncrementValueCount(value, valueCounts, correctCountSoFar)
		}
	}

	for col := 0; col < len(matrix[0]); col++ {
		for row := 0; row < len(matrix); row++ {
			value := matrix[row][col]
			correctCountSoFar := len(matrix) + col
			checkAndIncrementValueCount(value, valueCounts, correctCountSoFar)
		}
	}

	finalValues := make([]int, 0)
	for value := range valueCounts {
		if valueCounts[value] == len(matrix)+len(matrix[0]) {
			finalValues = append(finalValues, value)
		}
	}
	return finalValues
}

func initializeCountsOfPotentialValues(matrix [][]int) map[int]int {
	valueCounts := map[int]int{}

	smallerSide := matrix[0]
	if len(matrix) < len(matrix[0]) {
		smallerSide = []int{}
		for _, arr := range matrix {
			smallerSide = append(smallerSide, arr[0])
		}
	}

	for _, value := range smallerSide {
		valueCounts[value] = 0
	}
	return valueCounts
}

func checkAndIncrementValueCount(value int, valueCounts map[int]int, correctCountSoFar int) {
	if _, found := valueCounts[value]; !found {
		return
	}
	if valueCounts[value] != correctCountSoFar {
		return
	}
	valueCounts[value]++
}

// my solution
func repeatedMatrixValues(matrix [][]int) []int {
	valueInRows := make(map[int]int)
	for i := range matrix {
		for j := range matrix[i] {
			if prevIdx, ok := valueInRows[matrix[i][j]]; ok && prevIdx == i-1 || i == 0 {
				valueInRows[matrix[i][j]] = i
			}
		}
	}

	valueInCols := make(map[int]int)
	var i, j int
	for i < len(matrix) && j < len(matrix[0]) {
		for i <= len(matrix)-1 {
			if prevIdx, ok := valueInCols[matrix[i][j]]; ok && prevIdx == j-1 || j == 0 {
				valueInCols[matrix[i][j]] = j
			}
			i++
		}
		i = 0
		j++
	}

	result := make([]int, 0)
	for num, rowIdx := range valueInRows {
		if colIdx, ok := valueInCols[num]; rowIdx == len(matrix)-1 && ok && colIdx == len(matrix[rowIdx])-1  {
			result = append(result, num)
		}
	}

	return result
}
