package three_number_sort
// important question

func ThreeNumberSort(array []int, order []int) []int {
	firstValue, secondValue := order[0], order[1]
	firstIdx, secondIdx, thirdIdx := 0, 0, len(array)-1

	for secondIdx <= thirdIdx {
		value := array[secondIdx]
		if value == firstValue {
			array[firstIdx], array[secondIdx] = array[secondIdx], array[firstIdx]
			firstIdx += 1
			secondIdx += 1
		} else if value == secondValue {
			secondIdx += 1
		} else {
			array[secondIdx], array[thirdIdx] = array[thirdIdx], array[secondIdx]
			thirdIdx -= 1
		}
	}
	return array
}

func ThreeNumberSort2(array []int, order []int) []int {
	firstValue, thirdValue := order[0], order[2]

	var firstIdx int
	for idx := range array {
		if array[idx] == firstValue {
			array[firstIdx], array[idx] = array[idx], array[firstIdx]
			firstIdx += 1
		}
	}

	thirdIdx := len(array) - 1
	for idx := len(array) - 1; idx >= 0; idx-- {
		if array[idx] == thirdValue {
			array[thirdIdx], array[idx] = array[idx], array[thirdIdx]
			thirdIdx -= 1
		}
	}
	return array
}

func threeNumberSort3(slice []int, order []int) []int {
	g := make([]int, 3)

	for i := 0; i < len(slice); i++ {
		if order[0] == slice[i] {
			g[0]++
		}
		if order[1] == slice[i] {
			g[1]++
		}
		if order[2] == slice[i] {
			g[2]++
		}
	}

	var k int
	for i := 0; i < len(g); i++ {
		var j int
		for j < g[i] {
			slice[k] = order[i]
			j++
			k++
		}
	}

	return slice
}

func threeNumberSort2(slice []int, order []int) []int {
	g := make([]int, 4)

	for i := 0; i < len(slice); i++ {
		if order[0] == slice[i] {
			g[1]++
		}
		if order[1] == slice[i] {
			g[2]++
		}
		if order[2] == slice[i] {
			g[3]++
		}
	}

	for i := 1; i < len(g); i++ {
		g[i] += g[i-1]
	}

	for k := 1; k <= 3; k++ {
		for i, j := g[k], g[k]-1; i < len(slice) && j >= g[k-1]; i++ {
			if order[k-1] == slice[i] {
				for slice[i] == slice[j] {
					j--
				}

				slice[j], slice[i] = slice[i], slice[j]
				j--
			}
		}
	}

	return slice
}

func threeNumberSort1(slice []int, order []int) []int {
	var g, k int
	for i := 0; i < len(order); i++ {
		for j := 0; j < len(slice); j++ {
			if order[i] == slice[j] {
				g++
			}
		}
		for k < g {
			for j := g; j < len(slice); j++ {
				if order[i] == slice[j] {
					slice[k], slice[j] = slice[j], slice[k]
				}
			}
			k++
		}
	}

	return slice
}
