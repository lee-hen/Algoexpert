package minimum_passes_of_matrix

type IntPair struct {
	First, Second int
}

func MinimumPassesOfMatrix(matrix [][]int) int {
	passes := ConvertNegatives(matrix)
	if !containsNegative(matrix) {
		return passes - 1
	} else {
		return -1
	}
}

func ConvertNegatives(matrix [][]int) int {
	queue := getAllPositivePositions(matrix)
	var passes = 0

	for len(queue) > 0 {
		var currentSize = len(queue)

		for currentSize > 0 {
			// In Go, removing elements from the start of a list is an O(n)-time operation.
			// To make this an O(1)-time operation, we could use a more legitimate queue structure.
			// For our time complexity analysis, we'll assume this runs in O(1) time.
			nextElement := queue[0]
			queue = queue[1:]
			currentRow, currentCol := nextElement.First, nextElement.Second
			adjacentPositions := getAdjacentPositions(currentRow, currentCol, matrix)

			for _, position := range adjacentPositions {
				row, col := position.First, position.Second
				value := matrix[row][col]

				if value < 0 {
					matrix[row][col] *= -1
					queue = append(queue, IntPair{row, col})
				}
			}

			currentSize -= 1
		}

		passes += 1
	}
	return passes
}

func MinimumPassesOfMatrix1(matrix [][]int) int {
	passes := ConvertNegatives1(matrix)
	if !containsNegative(matrix) {
		return passes - 1
	} else {
		return -1
	}
}

func ConvertNegatives1(matrix [][]int) int {
	nextPassQueue := getAllPositivePositions(matrix)

	var passes = 0
	for len(nextPassQueue) > 0 {
		currentPassQueue := nextPassQueue
		nextPassQueue = make([]IntPair, 0)

		for len(currentPassQueue) > 0 {
			// In Go, removing elements from the start of a list is an O(n)-time operation.
			// To make this an O(1)-time operation, we could use a more legitimate queue structure.
			// For our time complexity analysis, we'll assume this runs in O(1) time.
			// Also, for this particular solution (Solution #1), we could actually
			// just turn this queue into a stack and replace the removal of the first element
			// with the removal of the last element.
			firstElement := currentPassQueue[0]
			currentPassQueue = currentPassQueue[1:]
			currentRow, currentCol := firstElement.First, firstElement.Second
			adjacentPositions := getAdjacentPositions(currentRow, currentCol, matrix)

			for _, position := range adjacentPositions {
				row, col := position.First, position.Second
				value := matrix[row][col]

				if value < 0 {
					matrix[row][col] *= -1
					nextPassQueue = append(nextPassQueue, IntPair{row, col})
				}
			}
		}
		passes += 1
	}

	return passes
}

func getAllPositivePositions(matrix [][]int) []IntPair {
	positivePositions := make([]IntPair, 0)

	for row := range matrix {
		for col := range matrix[row] {
			value := matrix[row][col]

			if value > 0 {
				positivePositions = append(positivePositions, IntPair{row, col})
			}
		}
	}

	return positivePositions
}

func getAdjacentPositions(row int, col int, matrix [][]int) []IntPair {
	adjacentPositions := make([]IntPair, 0)

	if row > 0 {
		adjacentPositions = append(adjacentPositions, IntPair{row - 1, col})
	}

	if row < len(matrix)-1 {
		adjacentPositions = append(adjacentPositions, IntPair{row + 1, col})
	}

	if col > 0 {
		adjacentPositions = append(adjacentPositions, IntPair{row, col - 1})
	}

	if col < len(matrix[0])-1 {
		adjacentPositions = append(adjacentPositions, IntPair{row, col + 1})
	}

	return adjacentPositions
}

func containsNegative(matrix [][]int) bool {
	for _, row := range matrix {
		for _, value := range row {
			if value < 0 {
				return true
			}
		}
	}
	return false
}

// first
//[0, -1, 3, 2, 0],
//[1, 2, -5, 1, -3],
//[3, 0, 0, -4, -1],

// second
//[0, 1, 3, 2, 0],
//[1, 2, 5, 1, 3],
//[3, 0, 0, 4, -1],

// third
//[0, 1, 3, 2, 0],
//[1, 2, 5, 1, 3],
//[3, 0, 0, 4, 1],

//[0, -1, -3, 2, 0],
//[1, -2, -5, -1, -3],
//[3, 0, 0, -4, -1],

//  2 -> -3 -1
//  1 -> -2

// -3 -> -1 -5
// -1 -> -3 -4
// -2 -> []
// my solution
func minimumPassesOfMatrix(matrix [][]int) int {
	negativesToConvert := make([][]int, 0)

	// if you don't ues queue implementation matrix[i][j] > 0 will be elevated see TestCase3
	// either check negativesToConvert has the positive value or just colored the matrix
	matrixColor := make([][]bool, len(matrix))
	for i := range matrix {
		matrixColor[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] > 0 {
				negativesToConvert = append(negativesToConvert, negativeNeighbors(i, j, matrix, matrixColor)...)
			}
		}
	}

	var numberToPass int
	convertNegatives(&numberToPass, negativesToConvert, matrix, matrixColor)

	if containsNegative(matrix) {
		return  - 1
	}

	return numberToPass
}

func convertNegatives(numberToPass *int, negativesToConvert, matrix[][]int, matrixColor [][]bool) {
	if len(negativesToConvert) > 0 {
		*numberToPass++
	} else {
		return
	}

	nextNegativesToConvert := make([][]int, 0)
	for _, negative := range negativesToConvert {
		if matrix[negative[0]][negative[1]] < 0 {
			matrix[negative[0]][negative[1]] *= -1
			nextNegativesToConvert = append(nextNegativesToConvert, negativeNeighbors(negative[0], negative[1], matrix, matrixColor)...)
		}
	}

	convertNegatives(numberToPass, nextNegativesToConvert, matrix, matrixColor)
}

func negativeNeighbors(i, j int, matrix [][]int, matrixColor [][]bool) [][]int{
	negativesToConvert := make([][]int, 0)
	neighbors := getNeighbors(i, j, matrix)
	for _, neighbor := range neighbors {
		if matrix[neighbor[0]][neighbor[1]] < 0 && !matrixColor[neighbor[0]][neighbor[1]]{
			matrixColor[neighbor[0]][neighbor[1]] = true
			negativesToConvert = append(negativesToConvert, []int{neighbor[0], neighbor[1]})
		}
	}

	return negativesToConvert
}


func getNeighbors(i, j int, matrix [][]int) [][]int {
	neighbors := make([][]int, 0)

	if i > 0 {
		neighbors = append(neighbors, []int{i - 1, j})
	}

	if i < len(matrix)-1 {
		neighbors = append(neighbors, []int{i + 1, j})
	}

	if j > 0 {
		neighbors = append(neighbors, []int{i, j - 1})
	}

	if j < len(matrix[0])-1 {
		neighbors = append(neighbors, []int{i, j + 1})
	}

	return neighbors
}
