package powerset
// important question

// PowerSet
// ex. [1 2 3]
// []
// [1]
// [2]
// [2 1]
// [3]
// [3 1]
// [3 2]
// [3 2 1]
func PowerSet(array []int) [][]int {
	subsets := [][]int{{}}
	for _, ele := range array {
		length := len(subsets)
		for i := 0; i < length; i++ {
			currentSubset := subsets[i]
			newSubSet := append([]int{}, currentSubset...)
			newSubSet = append(newSubSet, ele)
			subsets = append(subsets, newSubSet)
		}
	}
	return subsets
}

// PowerSet1
// recursive
func PowerSet1(array []int) [][]int {
	return powerSet1(array, len(array)-1)
}

func powerSet1(array []int, index int) [][]int {
	if index < 0 {
		return [][]int{{}}
	}
	subsets := powerSet1(array, index-1)
	ele := array[index]
	length := len(subsets)
	for i := 0; i < length; i++ {
		currentSubset := subsets[i]
		newSubset := append([]int{}, currentSubset...)
		newSubset = append(newSubset, ele)
		subsets = append(subsets, newSubset)
	}
	return subsets
}

// [1 2 3 4]
// []
// [1]
// [2]
// [1 2]
// [3]
// [1 3]
// [2 3]
// [1 2 3]
// [4]
// [1 4]
// [2 4]
// [1 2 4]
// [3 4]
// [1 3 4]
// [1 2 3 4]
func powerset(array []int) [][]int {
	return powerSetHelper(array, []int{})
}

func powerSetHelper(array, power []int) (powerSt [][]int) {
	if len(power) == 0 {
		powerSt = append(powerSt, []int{})
	} else {
		powerSt = append(powerSt, power)
	}

	for i := range array {
		newPower := make([]int, 0)
		newPower = append(newPower, array[i])
		newPower = append(newPower, power...)

		powerSt = append(powerSt, powerSetHelper(array[:i], newPower)...)
	}
	return
}
