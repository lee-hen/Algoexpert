package cycle_in_graph
// important question

type Color int

const (
	White Color = 0 // unvisited
	Grey  Color = 1 // grey
	Black Color = 2 // finished
)

// CycleInGraph
// O(v + e) time | O(v) space - where v is the number of
// vertices and e is the number of edges in the graph
func CycleInGraph(edges [][]int) bool {
	numberOfNodes := len(edges)
	colors := make([]Color, len(edges))
	for node := 0; node < numberOfNodes; node++ {
		if colors[node] != White {
			continue
		}
		containsCycle := traverseAndColorNodes(node, edges, colors)
		if containsCycle {
			return true
		}
	}
	return false
}

func traverseAndColorNodes(node int, edges [][]int, colors []Color) bool {
	colors[node] = Grey
	neighbors := edges[node]
	for _, neighbor := range neighbors {
		neighborColor := colors[neighbor]
		if neighborColor == Grey {
			return true
		}
		if neighborColor == Black {
			continue
		}
		containsCycle := traverseAndColorNodes(neighbor, edges, colors)
		if containsCycle {
			return true
		}
	}
	colors[node] = Black
	return false
}

// CycleInGraph1
// O(v + e) time | O(v) space - where v is the number of
// vertices and e is the number of edges in the graph
// Will the recursive calls of our algorithm affect its space complexity?
// There will never be more than v concurrent recursive calls on the call stack, since there are v vertices in the graph and since our algorithm naturally stops at cycles.
// Thus, the recursive calls won't change the O(v) space complexity of our algorithm.
func CycleInGraph1(edges [][]int) bool {
	numberOfNodes := len(edges)
	visited := make([]bool, len(edges))
	currentlyInStack := make([]bool, len(edges))
	for node := 0; node < numberOfNodes; node++ {
		if visited[node] {
			continue
		}
		containsCycle := isNodeInCycle(node, edges, visited, currentlyInStack)
		if containsCycle {
			return true
		}
	}
	return false
}

func isNodeInCycle(node int, edges [][]int, visited []bool, currentlyInStack []bool) bool {
	visited[node] = true
	currentlyInStack[node] = true
	neighbors := edges[node]
	for _, neighbor := range neighbors {
		if currentlyInStack[neighbor] {
			return true
		}

		if !visited[neighbor] {
			containsCycle := isNodeInCycle(neighbor, edges, visited, currentlyInStack)
			if containsCycle {
				return true
			}
		}
	}
	currentlyInStack[node] = false
	return false
}

// CycleInGraph
// my solution brute force
// ---------------
// {1, 3},
// {2, 3, 4},
// {0},
// {},
// {2, 5},
// {},
// ---------------
// 0 -> 1 3
// 1 -> 2 3 4
// 2 -> 0
// 3 -> nil
// 4 -> 2 5
// 5 -> nil
// ---------------
// 0 -> 1
// 1 -> 2
// 2 -> 4
// 3 -> 4
// 4 -> 2
// 5 -> nil
// ---------------
// 0 -> 1  1 -> 2  2 -> 4 4 -> 2
func cycleInGraph(edges [][]int) bool {
	if len(edges) < 1 {
		return false
	}

	traversedEdges := make([]int, 0)
	visited := make([]bool, len(edges))
	for i := 0; i < len(edges); i++ {
		if visited[i] {
			continue
		}
		if traverseEdges(i, visited, &traversedEdges, edges) {
			return true
		}
	}
	return false
}

func traverseEdges(current int, visited []bool, traversedEdges *[]int, edges [][]int) bool {
	vertices := edges[current]
	if len(edges[current]) == 0 {
		return false
	}
	*traversedEdges = append(*traversedEdges, current)
	visited[current] = true

	for _, v := range vertices {
		for _, vertex := range *traversedEdges {
			if v == vertex {
				return true
			}
		}
		if !visited[v] {
			if traverseEdges(v, visited, traversedEdges, edges) {
				return true
			}
		}
	}

	// if do not delete the underline situation is buggy
	// 0 -> [1] 1 -> [2, 4] 4 -> [3] 2 -> [4] 3 -> []
	*traversedEdges = (*traversedEdges)[:len(*traversedEdges)-1]
	return false
}
