package two_edge_connected_graph
// important question

// TwoEdgeConnectedGraph
// O(v + e) time | O(v) space - where v is the number of
// vertices and e is the number of edges in the graph
func TwoEdgeConnectedGraph(edges [][]int) bool {
	if len(edges) == 0 {
		return true
	}

	arrivalTimes := make([]int, len(edges))
	for i := range edges {
		arrivalTimes[i] = -1
	}
	startVertex := 0

	if getMinimumArrivalTimeOfAncestors(startVertex, -1, 0, arrivalTimes, edges) == -1 {
		return false
	}

	return areAllVerticesVisited(arrivalTimes)
}

func areAllVerticesVisited(arrivalTimes []int) bool {
	for _, time := range arrivalTimes {
		if time == -1 {
			return false
		}
	}
	return true
}

func getMinimumArrivalTimeOfAncestors(currentVertex, parent, currentTime int, arrivalTimes []int, edges [][]int) int {
	arrivalTimes[currentVertex] = currentTime

	var minimumArrivalTime = currentTime

	for _, destination := range edges[currentVertex] {
		if arrivalTimes[destination] == -1 {
			minimumArrivalTime = min(minimumArrivalTime,
				getMinimumArrivalTimeOfAncestors(destination, currentVertex, currentTime+1, arrivalTimes, edges))
		} else if destination != parent {
			minimumArrivalTime = min(minimumArrivalTime, arrivalTimes[destination])
		}
	}

	// A bridge was detected, which means the graph isn't two-edge-connected.
	if minimumArrivalTime == currentTime && parent != -1 {
		return -1
	}

	return minimumArrivalTime
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// twoEdgeConnectedGraph
// my solution
// not good
func twoEdgeConnectedGraph(edges [][]int) bool {
	if len(edges) == 1 && len(edges[0]) == 0 {
		return true
	}

	visited := make([]bool, len(edges))
	for i := range edges {
		if len(edges[i]) <= 1 {
			return false
		}

		if visited[i] {
			continue
		}

		if !dfs(i, i, i-1,i-1, edges, visited) {
			return false
		}
	}

	return true
}

func dfs(idx, start, parent, successor int, edges [][]int, visited []bool) bool {
	if visited[edges[start][len(edges[start])-1]] {
		return true
	}

	var comeBackThroughSuccessor bool

	for i := 0; i < len(edges[idx]); i++ {
		if edges[idx][i] == parent || idx == start && visited[edges[idx][i]]  {
			continue
		}

		if successor != parent && successor == idx && edges[idx][i] == start {
			return false
		}

		if edges[idx][i] == start || visited[edges[idx][i]] {
			visited[start] = true
			visited[successor] = true
			visited[idx] = true
			return true
		}

		if idx == start {
			successor = edges[idx][i]
		}

		comeBackThroughSuccessor = dfs(edges[idx][i], start, idx, successor, edges, visited)
	}
	return comeBackThroughSuccessor
}
