package river_sizes
// important question

// RiverSizes
// Better solution but not understandable
// O(wh) time | O(wh) space
func RiverSizes(matrix [][]int) []int {
	sizes := make([]int, 0)
	visited := make([][]bool, len(matrix))

	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := range matrix {
		for j := range matrix[i] {
			if visited[i][j] {
				continue
			}

			sizes = traverseNode(i, j, matrix, visited, sizes)
		}
	}
	return sizes
}

func traverseNode(i, j int, matrix [][]int, visited [][]bool, sizes []int) []int {
	currentRiverSize := 0
	nodesToExplore := [][]int{{i, j}}

	for len(nodesToExplore) > 0 {
		currentNode := nodesToExplore[0]
		nodesToExplore = nodesToExplore[1:]
		i, j := currentNode[0], currentNode[1]

		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}

		currentRiverSize += 1
		unvisitedNeighbors := getUnvisitedNeighbors(i, j, matrix, visited)

		for _, neighbor := range unvisitedNeighbors {
			nodesToExplore = append(nodesToExplore, neighbor)
		}
	}

	if currentRiverSize > 0 {
		sizes = append(sizes, currentRiverSize)
	}
	return sizes
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

// my corrected solution
func riverSizes(matrix [][]int) []int {
	riverPositions := make([]map[int]struct{}, len(matrix))
	for idxRow, matrixRow := range matrix {
		var riverCoordinate = make(map[int]struct{})
		for idxCol, matrixCol := range matrixRow {
			if matrixCol == 1 {
				riverCoordinate[idxCol] = struct{}{}
				riverPositions[idxRow] = riverCoordinate
			}
		}
	}
	riverSizes := traverseMatrix(matrix, riverPositions)

	if riverSizes == nil {
		return []int{}
	}
	return riverSizes
}

func traverseMatrix(matrix [][]int, riverPositions []map[int]struct{}) (riverSizes []int) {
	for idxR, riverCoordinate := range riverPositions {
		for idxC := range  riverCoordinate {
			var riverSize int
			riverSize += traverseRiverSize(idxR, idxC, riverPositions, matrix)
			riverSizes = append(riverSizes, riverSize)
		}
	}

	return
}

// remove_islands is more clear
func traverseRiverSize(idxRow, idxCol int, riverPositions []map[int]struct{}, matrix [][]int) (riverSize int) {
	if counted(riverPositions, idxRow, idxCol) {
		return
	}
	delete(riverPositions[idxRow], idxCol)

	if matrix[idxRow][idxCol] == 0 {
		return
	}
	riverSize++

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

	if below == 1 {
		riverSize += traverseRiverSize(idxRow+1, idxCol, riverPositions, matrix)
	}

	if right == 1 {
		riverSize += traverseRiverSize(idxRow, idxCol+1, riverPositions, matrix)
	}

	if above == 1 {
		riverSize += traverseRiverSize(idxRow-1, idxCol, riverPositions,  matrix)
	}

	if  left == 1  {
		riverSize += traverseRiverSize(idxRow, idxCol-1, riverPositions, matrix)
	}

	return
}

func counted(riverPositions []map[int]struct{}, idxRow int, idxCol int) bool {
	if _, ok := riverPositions[idxRow][idxCol]; ok  {
		return false
	}
	return true
}
