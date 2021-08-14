package number_of_ways_to_traverse_graph
// important question

// NumberOfWaysToTraverseGraph
// dynamic programming
// very interesting
func NumberOfWaysToTraverseGraph(width int, height int) int {
	numberOfWays := make([][]int, height+1)
	for i := range numberOfWays {
		numberOfWays[i] = make([]int, width+1)
	}
	for widthIdx := 1; widthIdx < width+1; widthIdx++ {
		for heightIdx := 1; heightIdx < height+1; heightIdx++ {
			if widthIdx == 1 || heightIdx == 1 {
				numberOfWays[heightIdx][widthIdx] = 1
			} else {
				waysLeft := numberOfWays[heightIdx][widthIdx-1]
				waysUp := numberOfWays[heightIdx-1][widthIdx]
				numberOfWays[heightIdx][widthIdx] = waysLeft + waysUp
			}
		}
	}
	return numberOfWays[height][width]
}

func NumberOfWaysToTraverseGraph2(width int, height int) int {
	xDistanceToCorner := width - 1
	yDistanceToCorner := height - 1

	numerator := factorial(xDistanceToCorner + yDistanceToCorner)
	denominator := factorial(xDistanceToCorner) * factorial(yDistanceToCorner)
	return numerator / denominator
}

func factorial(num int) int {
	result := 1
	for n := 2; n < num+1; n++ {
		result *= n
	}
	return result
}

func NumberOfWaysToTraverseGraph1(width int, height int) int {
	if width == 1 || height == 1 {
		return 1
	}
	return NumberOfWaysToTraverseGraph(width-1, height) + NumberOfWaysToTraverseGraph(width, height-1)
}

func numberOfWaysToTraverseGraph(width int, height int) int {
	numberOfNodes := (width+1) * (height+1)
	unavailableRight := make(map[int]struct{}, 0)
	var current int
	for current < numberOfNodes {
		current = current+width+1
		unavailableRight[current] = struct{}{}
	}

	var edges [][]int
	for i := 0; i < numberOfNodes; i++ {

		var edge []int
		right := i+1
		if right >= numberOfNodes {
			break
		}

		if _, found := unavailableRight[right]; !found && right != width+2 {
			edge = append(edge, right)
		}

		down := i+width+1
		if down < numberOfNodes && down != width+2 {
			edge = append(edge, down)
		}

		edges = append(edges, edge)
	}

	targetNode1 := numberOfNodes-2
	targetNode2 := numberOfNodes-width-2


	return numberOfWaysToReachTargetNode(1, targetNode2, edges) +
		numberOfWaysToReachTargetNode(width+1, targetNode1, edges)
}

func numberOfWaysToReachTargetNode(node, target int, edges [][]int) (numberOfWays int) {
	if node >= len(edges)-1 {
		return
	}
	neighbors := edges[node]
	for _, neighbor := range neighbors {

		if neighbor != target {
			numberOfWays += numberOfWaysToReachTargetNode(neighbor, target, edges)
		} else {
			numberOfWays=1
		}
	}
	return
}
