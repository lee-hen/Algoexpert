package largest_rectangle_under_skyline

// 1, 3, 3, 2, 4, 1, 5, 3, 2

// 2, 2, 4, 4, 4, 2, 2, 1



func LargestRectangleUnderSkyline(buildings []int) int {
	minValue := min(buildings)
	largestArea := minValue * len(buildings)
	//
	//for _, hi := range buildings {
	//
	//}
	//
	//
	//
	//
	//fmt.Println(minValue)
	//

	return largestArea
}


func min(array []int) int {
	currentMin := array[0]
	for _, el := range array {
		if el == 1 {
			return el
		}

		if currentMin > el {
			currentMin = el
		}
	}
	return currentMin
}

