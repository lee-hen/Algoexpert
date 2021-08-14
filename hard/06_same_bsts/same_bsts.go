package same_bsts
// important question

import (
	"math"
)

// SameBSTs
// nodes in each array, respectively, and d is the depth
// of the BST that they represent
func SameBSTs(arrayOne, arrayTwo []int) bool {
	return areSameBSTs(arrayOne, arrayTwo, 0, 0, math.MinInt32, math.MaxInt32)
}

func areSameBSTs(arrayOne, arrayTwo []int, rootIdxOne, rootIdxTwo int, minVal, maxVal int) bool {
	if rootIdxOne == -1 || rootIdxTwo == -1 {
		return rootIdxOne == rootIdxTwo
	}

	if arrayOne[rootIdxOne] != arrayTwo[rootIdxTwo] {
		return false
	}

	leftRootIdxOne := getIdxOfFirstSmaller(arrayOne, rootIdxOne, minVal)
	leftRootIdxTwo := getIdxOfFirstSmaller(arrayTwo, rootIdxTwo, minVal)
	rightRootIdxOne := getIdxOfFirstBiggerOrEqual(arrayOne, rootIdxOne, maxVal)
	rightRootIdxTwo := getIdxOfFirstBiggerOrEqual(arrayTwo, rootIdxTwo, maxVal)

	currentValue := arrayOne[rootIdxOne]
	leftAreSame := areSameBSTs(arrayOne, arrayTwo, leftRootIdxOne, leftRootIdxTwo, minVal, currentValue)
	rightAreSame := areSameBSTs(arrayOne, arrayTwo, rightRootIdxOne, rightRootIdxTwo, currentValue, maxVal)
	return leftAreSame && rightAreSame
}

func getIdxOfFirstSmaller(array []int, startingIdx, minVal int) int {
	// Find the index of the first smaller value after the startingIdx.
	// Make sure that this value is greater than or equal to the minVal,
	// which is the value of the previous parent node in the BST. If it
	// isn't, then that value is located in the left subtree of the
	// previous parent node.
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] < array[startingIdx] && array[i] >= minVal {
			return i
		}
	}
	return -1
}

func getIdxOfFirstBiggerOrEqual(array []int, startingIdx, maxVal int) int {
	// Find the index of the first bigger/equal value after the startingIdx.
	// Make sure that this value is smaller than maxVal, which is the value
	// of the previous parent node in the BST. If it isn't, then that value
	// is located in the right subtree of the previous parent node.
	for i := startingIdx + 1; i < len(array); i++ {
		if array[i] >= array[startingIdx] && array[i] < maxVal {
			return i
		}
	}
	return -1
}


// sliceOne  10, 15, 8, 12, 94, 81, 5,  2, 11
// sliceTwo  10, 8,  5, 15, 2,  12, 11, 94, 81

// 10(right)
// 15 12 94 81 11
// 15 12 11 94 81

//15(left)
// 12 11
// 12 11
//15(right)
// 94 81
// 84 81

// 10(left)
// 8 5 2
// 8 5 2

func sameBSTs(sliceOne, sliceTwo []int) bool {
	return sameBSTValidator(sliceOne, sliceTwo)
}

func sameBSTValidator(sliceOne, sliceTwo []int) bool {
	if len(sliceOne) == 0 && len(sliceTwo) == 0 {
		return true
	}

	if len(sliceOne) != len(sliceTwo) || sliceOne[0] != sliceTwo[0] {
		return false
	}

	left1 := make([]int, 0)
	right1 := make([]int, 0)
	for i := 1; i < len(sliceOne); i++ {
		if sliceOne[i] < sliceOne[0] {
			left1 = append(left1, sliceOne[i])
		} else {
			right1 = append(right1, sliceOne[i])
		}
	}

	left2 := make([]int, 0)
	right2 := make([]int, 0)
	for i := 1; i < len(sliceTwo); i++ {
		if sliceTwo[i] < sliceTwo[0] {
			left2 = append(left2, sliceTwo[i])
		} else {
			right2 = append(right2, sliceTwo[i])
		}
	}

	return sameBSTValidator(left1, left2) && sameBSTValidator(right1, right2)
}
