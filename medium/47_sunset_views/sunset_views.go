package sunset_views

import (
	"math"
)

func SunsetViews(buildings []int, direction string) []int {
	candidateBuildings := make([]int, 0)
	startIdx := len(buildings) - 1
	step := -1
	if direction == "EAST" {
		startIdx = 0
		step = 1
	}

	var idx = startIdx
	for idx >= 0 && idx < len(buildings) {
		buildingHeight := buildings[idx]

		for len(candidateBuildings) > 0 &&
			buildings[candidateBuildings[len(candidateBuildings)-1]] <= buildingHeight {
			candidateBuildings = candidateBuildings[:len(candidateBuildings)-1]
		}

		candidateBuildings = append(candidateBuildings, idx)
		idx += step
	}

	if direction == "WEST" {
		ReverseSlice(candidateBuildings)
	}

	return candidateBuildings
}


func SunsetViews2(buildings []int, direction string) []int {
	buildingsWithSunsetViews := make([]int, 0)
	startIdx := len(buildings) - 1
	step := -1
	if direction == "WEST" {
		startIdx = 0
		step = 1
	}

	idx := startIdx
	runningMaxHeight := 0
	for idx >= 0 && idx < len(buildings) {
		buildingHeight := buildings[idx]
		if buildingHeight > runningMaxHeight {
			buildingsWithSunsetViews = append(buildingsWithSunsetViews, idx)
		}
		runningMaxHeight = max(runningMaxHeight, buildingHeight)
		idx += step
	}

	if direction == "EAST" {
		ReverseSlice(buildingsWithSunsetViews)
	}
	return buildingsWithSunsetViews
}


// 5 4 3 2

// 0, 1, 2, 3, 4, 5, 6, 7
// 3, 5, 4, 4, 3, 1, 3, 2
// EAST(right)
// 1, 3, 6, 7

//  3, 5, 4, 4, 3, 1, 3, 2
// WEST(left)
// 0, 1

func sunsetViews1(buildings []int, direction string) (sunsetBuildings []int) {
	if len(buildings) == 0 {
		return []int{}
	}

	if direction == "WEST" {
		ReverseSlice(buildings)
	}

	current := math.MinInt32
	for idx := len(buildings)-1; idx >= 0; idx-- {
		if buildings[idx] > current {
			if direction == "WEST" {
				sunsetBuildings = append(sunsetBuildings, len(buildings)-idx-1 )
			} else {
				sunsetBuildings = append(sunsetBuildings, idx)
			}
		}
		current = max(current, buildings[idx])
	}

	if direction == "EAST" {
		ReverseSlice(sunsetBuildings)
	} else {
		ReverseSlice(buildings)
	}

	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sunsetViews2(buildings []int, direction string) []int {
	if len(buildings) == 0 {
		return []int{}
	}

	if direction == "EAST" {
		return eastSunsetViewsHelper(0, buildings)
	}

	sunsetBuildings := westSunsetViewsHelper(len(buildings)-1, buildings)
	ReverseSlice(sunsetBuildings)
	return sunsetBuildings
}

func eastSunsetViewsHelper(idx int, buildings []int) (sunsetBuildings []int) {
	if idx > len(buildings)-1 {
		return
	}

	max := math.MinInt32
	var peek = -1
	for idx < len(buildings) {
		if buildings[idx] >= max {
			max = buildings[idx]
			peek = idx
		}
		idx++
	}

	sunsetBuildings = append(sunsetBuildings, peek)
	sunsetBuildings = append(sunsetBuildings, eastSunsetViewsHelper(peek+1, buildings)...)
	return
}

func westSunsetViewsHelper(idx int, buildings []int) (sunsetBuildings []int) {
	if idx < 0 {
		return
	}

	max := math.MinInt32
	var peek = -1
	for idx >= 0 {
		if buildings[idx] >= max {
			max = buildings[idx]
			peek = idx
		}
		idx--
	}

	sunsetBuildings = append(sunsetBuildings, peek)
	sunsetBuildings = append(sunsetBuildings, westSunsetViewsHelper(peek-1, buildings)...)
	return
}


func ReverseSlice(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
