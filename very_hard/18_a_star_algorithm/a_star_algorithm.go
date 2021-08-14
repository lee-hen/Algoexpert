package a_star_algorithm
// important question

import "math"

type NodeInfo struct {
	h int // heuristic value
	g int // current shortest distance from the  start node to the give node
	f int // g + h
	visited bool
	from []int
}

func aStarAlgorithm(startRow int, startCol int, endRow int, endCol int, graph [][]int) [][]int {
	graphInfo := make([][]*NodeInfo, len(graph), len(graph))

	for i := range graph {
		graphInfo[i] = make([]*NodeInfo, len(graph[i]), len(graph[i]))
		for j := range graph[i] {
			graphInfo[i][j] = &NodeInfo{
				h: math.MaxInt32,
				g: math.MaxInt32,
				f: math.MaxInt32,
			}
		}
	}

	minHeap := make([][]int, 0)
	minHeap = append(minHeap, []int{startRow, startCol})
	startNodeInfo := graphInfo[startRow][startCol]
	startNodeInfo.g = 0
	startNodeInfo.h = endRow-startRow + endCol-startCol
	startNodeInfo.f = startNodeInfo.g + startNodeInfo.h

	for len(minHeap) != 0 {
		node := pop(&minHeap, graphInfo)
		if node[0] == endRow && node[1] == endCol {
			break
		}

		nodeInfo := graphInfo[node[0]][node[1]]
		if nodeInfo.visited {
			continue
		}

		neighbors := getNeighbors(node[0], node[1], graph)

		for _, neighbor := range neighbors {
			if graph[neighbor[0]][neighbor[1]] == 1 {
				continue
			}

			neighborInfo := graphInfo[neighbor[0]][neighbor[1]]
			if neighborInfo.visited {
				continue
			}

			neighborInfo.from = node

			neighborInfo.g = min(neighborInfo.g, nodeInfo.g+1)
			neighborInfo.h = min(neighborInfo.h, abs(endRow-neighbor[0]) + abs(endCol-neighbor[1]))
			neighborInfo.f = neighborInfo.g + neighborInfo.h

			minHeap = append(minHeap, neighbor)
		}

		graphInfo[node[0]][node[1]].visited = true
	}

	return reverseSlice(buildShortestPath(endRow, endCol, graphInfo))
}

func buildShortestPath(endRow, endCol int, graphInfo [][]*NodeInfo) [][]int {
	shortestPath := make([][]int, 0)
	nodeFrom := graphInfo[endRow][endCol].from
	if len(nodeFrom) != 0 {
		shortestPath = append(shortestPath, []int{endRow, endCol})
	}
	for len(nodeFrom) != 0 {
		shortestPath = append(shortestPath, nodeFrom)
		nodeFrom = graphInfo[nodeFrom[0]][nodeFrom[1]].from
	}

	return shortestPath
}

func pop(minHeap *[][]int, graphInfo [][]*NodeInfo) []int {
	g, f := math.MaxInt32, math.MaxInt32
	var node []int
	var idx int
	for i, position := range *minHeap {
		nodeInfo := graphInfo[position[0]][position[1]]
		if nodeInfo.g < g {
			g = nodeInfo.g
			idx = i
			node = position
		} else if nodeInfo.g == g {
			if nodeInfo.f < f {
				f = nodeInfo.f
				idx = i
				node = position
			}
		}
	}

	*minHeap = append((*minHeap)[:idx], (*minHeap)[idx+1:]...)
	return node
}

func getNeighbors(i, j int, graph [][]int) [][]int {
	neighbors := make([][]int, 0)

	if j < len(graph[i])-1 {
		neighbors = append(neighbors, []int{i, j+1})
	}

	if j > 0 {
		neighbors = append(neighbors, []int{i, j-1})
	}

	if i > 0 {
		neighbors = append(neighbors, []int{i-1, j})
	}

	if i < len(graph)-1 {
		neighbors = append(neighbors, []int{i+1, j})
	}
	return neighbors
}

func min(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num < curr {
			curr = num
		}
	}
	return curr
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func reverseSlice(slice [][]int) [][]int {
	start := 0
	end := len(slice) -1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start += 1
		end -= 1
	}
	return slice
}
