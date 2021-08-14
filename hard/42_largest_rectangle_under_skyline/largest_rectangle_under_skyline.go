package largest_rectangle_under_skyline
// important question

func LargestRectangleUnderSkyline(buildings []int) int {
	var currentArea int

	stack := make([]int, 0)
	buildings = append(buildings, 0)
	for idx, height := range buildings {
		// well the current height's idx is right bound of rectangle
		for len(stack) > 0 && buildings[stack[len(stack)-1]] >= height {
			topIdx := stack[len(stack)-1] // the target building to create the rectangle
			stack = stack[:len(stack)-1]

			// find left bound of the topIdx
			var unit int
			if len(stack) != 0 {
				rightBound := idx
				// the most left bound is buildings[stack[len(stack)-1]] < height
				// but every left bound to the topIdx might be used to the max area of the rectangle
				leftBound := stack[len(stack)-1]
				unit = rightBound-leftBound-1
			} else {
				unit = idx
			}

			// well create the area
			currentArea = max(currentArea, unit * buildings[topIdx])
		}
		stack = append(stack, idx)
	}

	return currentArea
}

// my solution
func largestRectangleUnderSkyline(buildings []int) int {
	var currentArea int

	for i := 0; i < len(buildings); i++ {
		var unit = 1

		for j := i+1; j < len(buildings); j++ {
			if buildings[i] > buildings[j] {
				break
			}
			unit++
		}

		for k := i - 1; k >= 0; k-- {
			if buildings[i] > buildings[k] {
				break
			}
			unit++
		}

		currentArea = max(currentArea, unit * buildings[i])
	}

	return currentArea
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
