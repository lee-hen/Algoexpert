package water_area
// important question

import (
	"math"
)

func WaterArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	leftIdx := 0
	rightIdx := len(heights) - 1

	leftMax := heights[leftIdx]
	rightMax := heights[rightIdx]

	surfaceArea := 0
	for leftIdx < rightIdx {
		if heights[leftIdx] < heights[rightIdx] {
			leftIdx++
			leftMax = max(leftMax, heights[leftIdx])
			surfaceArea += leftMax - heights[leftIdx]
		} else {
			rightIdx--
			rightMax = max(rightMax, heights[rightIdx])
			surfaceArea += rightMax - heights[rightIdx]
		}
	}

	return surfaceArea
}

func WaterArea2(heights []int) int {
	maxes := make([]int, len(heights))

	leftMax := 0
	for i, height := range heights {
		maxes[i] = leftMax
		leftMax = max(leftMax, height)
	}

	rightMax := 0
	for i := range heights {
		j := len(heights) - i - 1

		height := heights[j]
		minHeight := min(rightMax, maxes[j])

		if height < minHeight {
			maxes[j] = minHeight - height
		} else {
			maxes[j] = 0
		}

		rightMax = max(rightMax, height)
	}

	var waterArea int
	for _, height := range maxes {
		waterArea += height
	}
	return waterArea
}

//                      |
//                      |
//    |  .  .  .  .  .  |
//    |  .  .  .  .  .  |
//    |  .  .  .  .  .  |
//    |  .  .  |  .  .  |
//    |  .  .  |  .  .  |
//    |  .  .  |  .  .  |  .  .  .  .  .  |
//    |  .  .  |  .  .  |  .  .  .  .  .  |
// _  |  .  .  |  .  .  |  .  .  |  |  .  |  .  |
// 0, 8, 0, 0, 5, 0, 0,10, 0, 0, 1, 1, 0, 3  0  1
// 0  0  8  8  3  8  8  0  3  3  2  2  3  0  1  0
// 49

// 0, 100, 0, 0, 10, 1, 1, 10, 1, 0, 1, 1, 0, 0
// 0   0  10  10  0  9  9   0  0  1  0  0  0  0


// 0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3  0  1
// 0  0  8  8  3  8  8  0

// 0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3  0  1
//                       0  3  3  2  2  3  0  1  0

// my solution 2
func waterArea2(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	surfaceAreas := make([]int, len(heights))

	maxLeftHeight, maxRightHeight := heights[0], heights[len(heights)-1]
	left, right := 1, len(heights)-2

	for left <= right {
		maxRightHeight = max(heights[right], maxRightHeight)
		maxLeftHeight = max(heights[left], maxLeftHeight)

		if maxRightHeight < maxLeftHeight {
			if heights[right] <= maxRightHeight {
				surfaceAreas[right] = maxRightHeight - heights[right]
			}
			right--
		} else {
			if heights[left] < maxLeftHeight {
				surfaceAreas[left] = maxLeftHeight - heights[left]
			}
			left++
		}
	}

	var surfaceArea int
	for _, area := range surfaceAreas {
		surfaceArea += area
	}
	return surfaceArea
}

// my solution 1
func waterArea1(heights []int) int {
	waterHeights := make([]int, len(heights))
	left, maxHeight := -1, math.MinInt32

	for i := 0; i < len(heights); i++ {
		if heights[i] == 0 {
			continue
		}

		for j := left+1; j < i; j++ {
			if heights[i] > heights[j] && heights[i] > waterHeights[j] {
				waterHeights[j] = heights[i] - heights[j]
			}

			maxHeight = max(heights[j], maxHeight)
			if maxHeight <= heights[i] {
				waterHeights[j] = maxHeight-heights[j]
				left = j
			}

			if j == 0 || j == len(heights)-1 {
				waterHeights[j] = 0
			}
		}
	}

	var waterArea int
	for _, height := range waterHeights {
		waterArea += height
	}

	return waterArea
}

func max(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
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
