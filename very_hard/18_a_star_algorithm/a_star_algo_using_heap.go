package a_star_algorithm

import (
	"fmt"
	"math"
)

// AStarAlgorithm
// O(w * h * log(w * h)) time | O(w * h) space - where
// w is the width of the graph and h is the height
func AStarAlgorithm(startRow int, startCol int, endRow int, endCol int, graph [][]int) [][]int {
	nodes := initializeNodes(graph)

	startNode := nodes[startRow][startCol]
	endNode := nodes[endRow][endCol]

	startNode.distanceFromStart = 0
	startNode.estimatedDistanceToEnd = calculateManhattanDistance(startNode, endNode)

	nodesToVisit := newMinHeap([]*Node{startNode})

	for !nodesToVisit.IsEmpty() {
		currentMinDistanceNode := nodesToVisit.Remove()
		if currentMinDistanceNode == endNode {
			break
		}

		neighbors := getNeighboringNodes(currentMinDistanceNode, nodes)
		for _, neighbor := range neighbors {
			if neighbor.value == 1 {
				continue
			}

			tentativeDistanceToNeighbor := currentMinDistanceNode.distanceFromStart + 1
			if tentativeDistanceToNeighbor >= neighbor.distanceFromStart {
				continue
			}

			neighbor.cameFrom = currentMinDistanceNode
			neighbor.distanceFromStart = tentativeDistanceToNeighbor
			neighbor.estimatedDistanceToEnd = tentativeDistanceToNeighbor +
			calculateManhattanDistance(neighbor, endNode)

			if !nodesToVisit.containsNode(neighbor) {
				nodesToVisit.Insert(neighbor)
			} else {
				nodesToVisit.Update(neighbor)
			}
		}
	}

	return reconstructPath(endNode)
}

type Node struct {
	id                     string
	row                    int
	col                    int
	value                  int
	distanceFromStart      int
	estimatedDistanceToEnd int
	cameFrom               *Node
}

func newNode(row, col, value int) *Node {
	return &Node{
		id:                     fmt.Sprintf("%d-%d", row, col),
		row:                    row,
		col:                    col,
		value:                  value,
		distanceFromStart:      math.MaxInt32,
		estimatedDistanceToEnd: math.MaxInt32,
		cameFrom:               nil,
	}
}

func initializeNodes(graph [][]int) [][]*Node {
	nodes := make([][]*Node, 0)
	for row := range graph {
		newRow := make([]*Node, 0)
		for col := range graph[0] {
			value := graph[row][col]
			newRow = append(newRow, newNode(row, col, value))
		}
		nodes = append(nodes, newRow)
	}
	return nodes
}

func calculateManhattanDistance(currentNode *Node, endNode *Node) int {
	return abs(currentNode.col-endNode.col) + abs(currentNode.row-endNode.row)
}

func getNeighboringNodes(node *Node, nodes [][]*Node) []*Node {
	neighbors := make([]*Node, 0)
	numRows := len(nodes)
	numCols := len(nodes[0])

	row := node.row
	col := node.col

	if row < numRows-1 {
		neighbors = append(neighbors, nodes[row+1][col])
	} // DOWN
	if row > 0 {
		neighbors = append(neighbors, nodes[row-1][col])
	} // UP
	if col < numCols-1 {
		neighbors = append(neighbors, nodes[row][col+1])
	} // RIGHT
	if col > 0 {
		neighbors = append(neighbors, nodes[row][col-1])
	} // LEFT

	return neighbors
}

func reconstructPath(endNode *Node) [][]int {
	if endNode.cameFrom == nil {
		return [][]int{}
	}

	currentNode := endNode
	path := make([][]int, 0)
	for currentNode != nil {
		path = append(path, []int{currentNode.row, currentNode.col})
		currentNode = currentNode.cameFrom
	}
	return reversePath(path)
}

func reversePath(path [][]int) [][]int {
	newPath := make([][]int, len(path))
	for i := range path {
		j := len(path) - i - 1
		newPath[i] = path[j]
	}
	return newPath
}

type MinHeap struct {
	array               []*Node
	nodePositionsInHeap map[string]int
}

func newMinHeap(array []*Node) *MinHeap {
	nodePositionsInHeap := map[string]int{}
	for i, node := range array {
		nodePositionsInHeap[node.id] = i
	}
	heap := &MinHeap{array: array, nodePositionsInHeap: nodePositionsInHeap}
	heap.buildHeap()
	return heap
}

func (h *MinHeap) IsEmpty() bool { return len(h.array) == 0 }

// Remove
// O(log(n)) time | O(1) space
func (h *MinHeap) Remove() *Node {
	if h.IsEmpty() {
		return nil
	}

	h.swap(0, len(h.array)-1)

	peeked := h.array[len(h.array)-1]
	h.array = h.array[0 : len(h.array)-1]

	delete(h.nodePositionsInHeap, peeked.id)
	h.siftDown(0, len(h.array)-1)
	return peeked
}

// Update
// O(log(n)) time | O(1) space
func (h *MinHeap) Update(node *Node) {
	h.siftUp(h.nodePositionsInHeap[node.id])
}

// Insert
// O(log(n)) time | O(1) space
func (h *MinHeap) Insert(node *Node) {
	h.array = append(h.array, node)
	h.nodePositionsInHeap[node.id] = len(h.array) - 1
	h.siftUp(len(h.array) - 1)
}

func (h *MinHeap) containsNode(node *Node) bool {
	_, found := h.nodePositionsInHeap[node.id]
	return found
}

func (h MinHeap) swap(i, j int) {
	h.nodePositionsInHeap[h.array[i].id] = j
	h.nodePositionsInHeap[h.array[j].id] = i
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

// O(n) time | O(1) space
func (h *MinHeap) buildHeap() {
	first := (len(h.array) - 2) / 2
	for currentIdx := first + 1; currentIdx >= 0; currentIdx-- {
		h.siftDown(currentIdx, len(h.array)-1)
	}
}

// O(log(n)) time | O(1) space
func (h *MinHeap) siftDown(currentIdx, endIdx int) {
	childOneIdx := currentIdx*2 + 1
	for childOneIdx <= endIdx {
		childTwoIdx := -1
		if currentIdx*2+2 <= endIdx {
			childTwoIdx = currentIdx*2 + 2
		}
		indexToSwap := childOneIdx
		c1Distance := h.array[childOneIdx].estimatedDistanceToEnd
		if childTwoIdx > -1 && h.array[childTwoIdx].estimatedDistanceToEnd < c1Distance {
			indexToSwap = childTwoIdx
		}

		if h.array[indexToSwap].estimatedDistanceToEnd < h.array[currentIdx].estimatedDistanceToEnd {
			h.swap(currentIdx, indexToSwap)
			currentIdx = indexToSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

// O(log(n)) time | O(1) space
func (h *MinHeap) siftUp(currentIdx int) {
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && h.array[currentIdx].estimatedDistanceToEnd < h.array[parentIdx].estimatedDistanceToEnd {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}
