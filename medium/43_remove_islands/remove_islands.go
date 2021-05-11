package remove_islands

func RemoveIslands(matrix [][]int) [][]int {
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			rowIsBorder := row == 0 || row == len(matrix)-1
			colIsBorder := col == 0 || col == len(matrix[row])-1
			isBorder := rowIsBorder || colIsBorder

			if !isBorder {
				continue
			}

			if matrix[row][col] != 1 {
				continue
			}

			changeOnesConnectedToBorderToTwos(matrix, row, col)
		}
	}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			color := matrix[row][col]

			if color == 1 {
				matrix[row][col] = 0
			} else if color == 2 {
				matrix[row][col] = 1
			}
		}
	}

	return matrix
}

func changeOnesConnectedToBorderToTwos(matrix [][]int, startRow, startCol int) {
	stack := [][]int{{startRow, startCol}}

	var currentPosition []int
	for len(stack) > 0 {
		currentPosition, stack = stack[len(stack)-1], stack[:len(stack)-1]
		currentRow, currentCol := currentPosition[0], currentPosition[1]
		matrix[currentRow][currentCol] = 2

		neighbors := getNeighbors(matrix, currentRow, currentCol)
		for _, neighbor := range neighbors {
			row, col := neighbor[0], neighbor[1]

			if matrix[row][col] != 1 {
				continue
			}
			stack = append(stack, neighbor)
		}
	}
}

func getNeighbors(matrix [][]int, row, col int) [][]int {
	neighbors := make([][]int, 0)
	numRows := len(matrix)
	numCols := len(matrix[row])

	if row-1 >= 0 {
		neighbors = append(neighbors, []int{row - 1, col}) // UP
	}
	if row+1 < numRows {
		neighbors = append(neighbors, []int{row + 1, col}) // DOWN
	}
	if col-1 >= 0 {
		neighbors = append(neighbors, []int{row, col - 1}) // LEFT
	}
	if col+1 < numCols {
		neighbors = append(neighbors, []int{row, col + 1}) // RIGHT
	}
	return neighbors
}

func RemoveIslands1(matrix [][]int) [][]int {
	nodesToRemove := make([][]int, 0)
	visited := make([][]bool, len(matrix))

	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || i == len(matrix)-1 || j == 0 || j == len(matrix[i])-1 {
				continue
			}

			if visited[i][j] {
				continue
			}

			nodesToRemove = append(nodesToRemove, findNodesToRemove(i, j,  matrix, visited)...)
		}
	}

	for _, node := range nodesToRemove {
		i, j := node[0], node[1]
		matrix[i][j] = 0
	}

	return matrix
}

func findNodesToRemove(i, j int, matrix [][]int, visited [][]bool)  (nodesToRemove [][]int)  {
	nodesToExplore := [][]int{{i, j}}
	touchedBorder := false
	currentNodes := make([][]int, 0)

	for len(nodesToExplore) > 0 {
		currentNode := nodesToExplore[0]
		nodesToExplore = nodesToExplore[1:]
		i, j := currentNode[0], currentNode[1]

		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}

		if i == 0 || i == len(matrix)-1 || j == 0 || j == len(matrix[i])-1 {
			touchedBorder = true
		}

		currentNodes = append(currentNodes, []int{i, j})

		unvisitedNeighbors := getUnvisitedNeighbors(i, j, matrix, visited)
		for _, neighbor := range unvisitedNeighbors {
			nodesToExplore = append(nodesToExplore, neighbor)
		}
	}

	if !touchedBorder && len(currentNodes) != 0 {
		nodesToRemove = append(nodesToRemove, currentNodes...)
	}

	return
}

func getUnvisitedNeighbors(i, j int, matrix [][]int, visited [][]bool) [][]int {
	unvisitedNeighbors := make([][]int, 0)

	if i > 0 && !visited[i-1][j] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i - 1, j})
	}
	if i < len(matrix)-1 && !visited[i+1][j] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i + 1, j})
	}
	if j > 0 && !visited[i][j-1] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i, j - 1})
	}
	if j < len(matrix[0])-1 && !visited[i][j+1] {
		unvisitedNeighbors = append(unvisitedNeighbors, []int{i, j + 1})
	}

	return unvisitedNeighbors
}

func removeIslands(matrix [][]int) [][]int {
	visited := make([][]bool, len(matrix))

	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	nodesToRemove := make([][]int, 0)
	for i := range matrix {
		for j := range matrix[i] {
			if i == 0 || i == len(matrix)-1 || j == 0 || j == len(matrix[i])-1 {
				continue
			}

			if visited[i][j] {
				continue
			}

			touchedBorder := false
			nodes := traverseNode(i, j, &touchedBorder, visited, matrix)
			if touchedBorder {
				continue
			}
			nodesToRemove = append(nodesToRemove, nodes...)
		}
	}

	for _, node := range nodesToRemove {
		i, j := node[0], node[1]
		matrix[i][j] = 0
	}

	return matrix
}

func traverseNode(i, j int, touchedBorder *bool, visited [][]bool, matrix [][]int) (nodesToRemove [][]int) {
	if visited[i][j] {
		return
	}
	visited[i][j] = true

	if matrix[i][j] == 0 {
		return
	}

	if i == 0 || i == len(matrix)-1 || j == 0 || j == len(matrix[i])-1 {
		*touchedBorder = true
	}
	nodesToRemove = append(nodesToRemove, []int{i, j})

	var above, below, left, right int
	if i > 0 {
		above = matrix[i-1][j]
	}

	if i < len(matrix)-1 {
		below = matrix[i+1][j]
	}

	if j > 0 {
		left = matrix[i][j-1]
	}

	if j < len(matrix[i])-1 {
		right = matrix[i][j+1]
	}

	if below == 1 {
		nodes := traverseNode(i+1, j, touchedBorder, visited, matrix)
		if !*touchedBorder {
			nodesToRemove = append(nodesToRemove, nodes...)
		}
	}

	if right == 1 {
		nodes := traverseNode(i, j+1, touchedBorder, visited, matrix)
		if !*touchedBorder {
			nodesToRemove = append(nodesToRemove, nodes...)
		}
	}

	if above == 1 {
		nodes := traverseNode(i-1, j, touchedBorder, visited,  matrix)
		if !*touchedBorder {
			nodesToRemove = append(nodesToRemove, nodes...)
		}
	}

	if  left == 1  {
		nodes := traverseNode(i, j-1, touchedBorder, visited, matrix)
		if !*touchedBorder {
			nodesToRemove = append(nodesToRemove, nodes...)
		}
	}
	return
}
