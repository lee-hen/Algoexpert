package dijkstras_algorithm
// important question

import (
	"math"
)

type Item struct {
	Vertex   int
	Distance int
}

type MinHeap struct {
	array     []Item
	vertexMap map[int]int
}

func NewMinHeap(array []Item) *MinHeap {
	vertexMap := map[int]int{}

	// Holds the position in the heap that each vertex is at
	// There is no new item to insert into the heap we keep remove the vertex that has already visited
	// and instead of insert new item into the heap we use vertexMap to update for the current node's destination(item)
	// and then sift the item up
	for _, item := range array {
		vertexMap[item.Vertex] = item.Vertex
	}

	heap := &MinHeap{array: array, vertexMap: vertexMap}
	heap.buildHeap()
	return heap
}

// DijkstrasAlgorithm
// O((v + e) * log(v)) time | O(v) space - where v is the number
// of vertices and e is the number of edges in the input graph
func DijkstrasAlgorithm(start int, edges [][][]int) []int {
	minDistances := make([]int, len(edges), len(edges))
	minDistancesPairs := make([]Item, len(edges), len(edges))

	for i := range edges {
		minDistances[i] = math.MaxInt32
	}

	for i := range edges {
		minDistancesPairs[i] = Item{i, math.MaxInt32}
	}

	minDistances[start] = 0
	dijkstrasUsingMinHeap(start, minDistancesPairs, minDistances, edges)

	return minDistances
}

func dijkstrasUsingMinHeap(start int, minDistancesPairs []Item, minDistances []int, edges [][][]int) {
	minDistancesHeap := NewMinHeap(minDistancesPairs)
	minDistancesHeap.Update(start, 0)

	for !minDistancesHeap.IsEmpty() {
		vertex, currentMinDistance := minDistancesHeap.Remove()
		if currentMinDistance == math.MaxInt32 {
			break
		}

		for _, edge := range edges[vertex] {
			destination, distanceToDestination := edge[0], edge[1]

			newPathDistance := currentMinDistance + distanceToDestination
			if newPathDistance < minDistances[destination] {
				minDistancesHeap.Update(destination, newPathDistance)
			}

			minDistances[destination] = min(minDistances[destination], newPathDistance)
		}
	}

	for i, distance := range minDistances {
		if distance == math.MaxInt32 {
			minDistances[i] = -1
		}
	}
}

func (h *MinHeap) IsEmpty() bool { return h.length() == 0 }

// Remove
// O(log(n)) time | O(1) space
func (h *MinHeap) Remove() (int, int) {
	l := h.length()
	h.swap(0, l-1)
	peeked := h.array[l-1]
	h.array = h.array[0 : l-1]
	delete(h.vertexMap, peeked.Vertex)
	h.siftDown(0, l-2)
	return peeked.Vertex, peeked.Distance
}

// Update
// O(log(n)) time | O(1) space
func (h *MinHeap) Update(vertex int, value int) {
	h.array[h.vertexMap[vertex]] = Item{vertex, value}
	h.siftUp(h.vertexMap[vertex])
}

func (h MinHeap) swap(i, j int) {
	h.vertexMap[h.array[i].Vertex] = j
	h.vertexMap[h.array[j].Vertex] = i
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func (h MinHeap) length() int { return len(h.array) }

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
		if childTwoIdx > -1 && h.array[childTwoIdx].Distance < h.array[childOneIdx].Distance {
			indexToSwap = childTwoIdx
		}
		if h.array[indexToSwap].Distance < h.array[currentIdx].Distance {
			h.swap(currentIdx, indexToSwap)
			currentIdx = indexToSwap
			childOneIdx = currentIdx*2 + 1
		} else {
			return
		}
	}
}

func (h *MinHeap) siftUp(currentIdx int) {
	parentIdx := (currentIdx - 1) / 2
	for currentIdx > 0 && h.array[currentIdx].Distance < h.array[parentIdx].Distance {
		h.swap(currentIdx, parentIdx)
		currentIdx = parentIdx
		parentIdx = (currentIdx - 1) / 2
	}
}
