package apartment_hunting
// important question

import (
	"math"
)

type Block map[string]bool

// ApartmentHunting
// O(br) time | O(br) space - where b is the number of blocks
// and r is the number of requirements.
func ApartmentHunting(blocks []Block, reqs []string) int {
	minDistancesFromBlocks := make([][]int, 0)
	for _, req := range reqs {
		minDistancesFromBlocks = append(minDistancesFromBlocks, getMinDistances(blocks, req))
	}
	maxDistancesAtBlocks := getMaxDistancesAtBlocks(blocks, minDistancesFromBlocks)

	var optimalBlockIdx int
	smallestMaxDistance := math.MaxInt32
	for i, currentDistance := range maxDistancesAtBlocks {
		if currentDistance < smallestMaxDistance {
			smallestMaxDistance = currentDistance
			optimalBlockIdx = i
		}
	}
	return optimalBlockIdx
}

func getMinDistances(blocks []Block, req string) []int {
	minDistances := make([]int, len(blocks))
	closestReq := math.MaxInt32
	for i := range blocks {
		if val, found := blocks[i][req]; found && val {
			closestReq = i
		}
		minDistances[i] = distanceBetween(i, closestReq)
	}

	for i := len(blocks) - 1; i >= 0; i-- {
		if val, found := blocks[i][req]; found && val {
			closestReq = i
		}
		minDistances[i] = min(minDistances[i], distanceBetween(i, closestReq))
	}
	return minDistances
}

func getMaxDistancesAtBlocks(blocks []Block, minDistancesFromBlocks [][]int) []int {
	maxDistancesAtBlocks := make([]int, len(blocks))
	for i := range blocks {
		minDistancesAtBlock := make([]int, 0)
		for _, distances := range minDistancesFromBlocks {
			minDistancesAtBlock = append(minDistancesAtBlock, distances[i])
		}
		maxDistancesAtBlocks[i] = maxSlice(minDistancesAtBlock)
	}
	return maxDistancesAtBlocks
}

func maxSlice(array []int) int {
	if len(array) == 0 {
		return 0
	}

	max := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

// ApartmentHunting1
// O(b^2*r) time | O(b) space - where b is the number of blocks
// and r is the number of requirements.
func ApartmentHunting1(blocks []Block, reqs []string) int {
	maxDistancesAtBlocks := make([]int, len(blocks))
	for i := range blocks {
		maxDistancesAtBlocks[i] = -1
		for _, req := range reqs {
			closestReqDistance := math.MaxInt32
			for j := range blocks {
				if blocks[j][req] {
					closestReqDistance = min(closestReqDistance, distanceBetween(i, j))
				}
			}
			maxDistancesAtBlocks[i] = max(maxDistancesAtBlocks[i], closestReqDistance)
		}
	}

	var optimalBlockIdx int
	smallestMaxDistance := math.MaxInt32
	for i, currentDistance := range maxDistancesAtBlocks {
		if currentDistance < smallestMaxDistance {
			smallestMaxDistance = currentDistance
			optimalBlockIdx = i
		}
	}
	return optimalBlockIdx
}

func distanceBetween(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//[
// {
//    "gym": false,
//    "school": true,
//    "store": false,
// },
// {
//    "gym": true,
//    "school": false,
//    "store": false,
// },
// {
//    "gym": true,
//    "school": true,
//    "store": false,
// },
// {
//    "gym": false,
//    "school": true,
//    "store": false,
// },
// {
//    "gym": false,
//    "school": true,
//    "store": true,
// },
//]

// reqs = ["gym", "school", "store"]
// 3 at index 3, the farthest you'd have to walk to reach a gym, a school, or a store is 1 block; at any other index, you'd have to walk farther

func apartmentHunting(blocks []Block, reqs []string) int {
	var maxStep = math.MaxInt32
	var targetIdx int

	for idx, block := range blocks {
		unSatisfiedReqs := make(map[string]struct{})
		for _, req := range reqs {
			if ok, found := block[req]; found && !ok {
				unSatisfiedReqs[req] = struct{}{}
			}
		}

		var step, leftStep, rightStep int
		if len(unSatisfiedReqs) == 0 {
			if step < maxStep {
				maxStep = step
				targetIdx = idx
			}
			break
 		} else {
			// Scan left / right
			left, right := idx-1, idx+1
 			for len(unSatisfiedReqs) != 0 && (left >= 0 || right < len(blocks)) {
 				// Scan left
 				if left >= 0 {
					for req := range unSatisfiedReqs {
						if ok, found := blocks[left][req]; found && ok {
							delete(unSatisfiedReqs, req)
						}
					}
				}

				// Scan right
				if right < len(blocks) {
					for req := range unSatisfiedReqs {
						if ok, found := blocks[right][req]; found && ok {
							delete(unSatisfiedReqs, req)
						}
					}
				}

				left--
				right++
				leftStep++
				rightStep++
			}

			if leftStep > rightStep {
				step = leftStep
			} else {
				step = rightStep
			}

			if step < maxStep {
				maxStep = step
				targetIdx = idx
			}
		}

	}

	return targetIdx
}

