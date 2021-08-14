package staircase_traversal
// important question

// maxSteps  2
// height    4

// currentHeight    0   1   2   3   4   5
// currentWays      0   1   2   3   5   8 (5          -          2  +          5)
// waysToTop        1   1   2   3   5   8 (waysToTop[4]-waysToTop[2]+waysToTop[4])

// StaircaseTraversal
// waysToTop[currentHeight] =
// waysToTop[currentHeight-1]-waysToTop[currentHeight-maxSteps-1]+waysToTop[currentHeight-1]
func StaircaseTraversal(height int, maxSteps int) int {
	currentNumberOfWays := 0
	waysToTop := []int{1}

	// ex currentHeight = 4 currentNumberOfWays = 3
	for currentHeight := 1; currentHeight < height+1; currentHeight++ {
		startOfWindow := currentHeight - maxSteps - 1 // 4 - 2 - 1  = 1
		endOfWindow := currentHeight - 1              // 3

		if startOfWindow >= 0 {
			currentNumberOfWays -= waysToTop[startOfWindow] // 3-1= 2
		}

		currentNumberOfWays += waysToTop[endOfWindow] // 2 + 3 = 5
		waysToTop = append(waysToTop, currentNumberOfWays)
	}
	return waysToTop[height]
}

func StaircaseTraversal3(height int, maxSteps int) int {
	waysToTop := make([]int, height+1)
	waysToTop[0] = 1
	waysToTop[1] = 1

	for currentHeight := 2; currentHeight < height+1; currentHeight++ {
		for step := 1; step <= maxSteps && step <= currentHeight; step++ {
			waysToTop[currentHeight] += waysToTop[currentHeight-step]
		}
	}
	return waysToTop[height]
}

func StaircaseTraversal2(height int, maxSteps int) int {
	return numberOfWaysToTop2(height, maxSteps, map[int]int{0: 1, 1: 1})
}

func numberOfWaysToTop2(height int, maxSteps int, memoize map[int]int) int {
	if ways, found := memoize[height]; found {
		return ways
	}

	var numberOfWays = 0
	for step := 1; step < min(maxSteps, height)+1; step++ {
		numberOfWays += numberOfWaysToTop2(height-step, maxSteps, memoize)
	}

	memoize[height] = numberOfWays
	return numberOfWays
}

func StaircaseTraversal1(height int, maxSteps int) int {
	return numberOfWaysToTop1(height, maxSteps)
}

func numberOfWaysToTop1(height int, maxSteps int) int {
	if height <= 1 {
		return 1
	}

	var numberOfWays = 0
	// min(maxSteps, height) used to reduce unnecessary loops
	for step := 1; step < min(maxSteps, height)+1; step++ {
		numberOfWays += numberOfWaysToTop1(height-step, maxSteps)
	}

	return numberOfWays
}

//func numberOfWaysToTop1(height int, maxSteps int) int {
//	if height == 1 || height == 0 {
//		return 1
//	}
//
//	if height < 0 {
//		return 0
//	}
//
//	var numberOfWays = 0
//	for step := 1; step <= maxSteps; step++ {
//		numberOfWays += numberOfWaysToTop1(height-step, maxSteps)
//	}
//
//	return numberOfWays
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// height 4
// maxSteps 2

// 1 1 1 1
// 1 1 2
// 1 2 1
// 2 1 1
// 2 2
func staircaseTraversal(height int, maxSteps int) int {
	return staircaseTraversalHelper(maxSteps, 0, height, map[int]int{height: 1})
}

func staircaseTraversalHelper(maxSteps, currentHeight, height int, memoize map[int]int) (ways int) {
	if currentHeight > height {
		return 0
	}

	if ways, found := memoize[currentHeight]; found {
		return ways
	}

	// why do not use min(maxSteps, height) because maxSteps is always less than height
	for step := 1; step <= maxSteps; step++ {
		ways += staircaseTraversalHelper(maxSteps, step+currentHeight, height, memoize)
	}

	memoize[currentHeight] = ways
	return ways
}
