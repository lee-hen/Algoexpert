package next_greater_element
// important question

import (
	"math"
)

// NextGreaterElement
// better solution
// O(n) time | O(n) space - where n is the length of the array
func NextGreaterElement(array []int) []int {
	result := make([]int, 0)
	for range array {
		result = append(result, -1)
	}

	stack := make([]int, 0)

	for idx := 2*len(array) - 1; idx >= 0; idx-- {
		circularIdx := idx % len(array)

		for len(stack) > 0 {
			if stack[len(stack)-1] <= array[circularIdx] {
				stack = stack[:len(stack)-1]
			} else {
				result[circularIdx] = stack[len(stack)-1]
				break
			}
		}

		stack = append(stack, array[circularIdx])
	}

	return result
}

// NextGreaterElement1
// better solution
// O(n) time | O(n) space - where n is the length of the array
func NextGreaterElement1(array []int) []int {
	result := make([]int, 0)
	for range array {
		result = append(result, -1)
	}

	stack := make([]int, 0)
	for idx := 0; idx < 2*len(array); idx++ {
		circularIdx := idx % len(array)

		for len(stack) > 0 && array[stack[len(stack)-1]] < array[circularIdx] {
			var top int
			top, stack = stack[len(stack)-1], stack[:len(stack)-1]
			result[top] = array[circularIdx]
		}

		stack = append(stack, circularIdx)
	}

	return result
}

// [1, 2]
// [2, -1]

// [2, 5, -3, -4, 6, 7, 2]
//  0  1   2   3  4  5  6
// [5, 6, 6, 6, 7, -1, 5]

// 2, 5, -3, -4, 6, 7, 2

// 5, -3, -4, 6, 7, 2, 2
// 5  6  6  6   7   -1  5

// NextGreaterElement
// my solution worst time complexity is O(n^2)
func nextGreaterElement(array []int) []int {
	result := make([]int, len(array))

	maxVal := math.MinInt32
	for idx := 0; idx < len(array); idx++ {
		if array[idx] > maxVal {
			maxVal = array[idx]
		}

		result[idx] = math.MaxInt32
	}

	nextElements := make([]int, len(array), len(array))
	copy(nextElements, array)
	reverseSlice(nextElements)

	nextElements = fillNextGreaterElement(array, nextElements, result)
	reverseSlice(nextElements)

	for idx := len(array)-1; idx >=0 ; idx-- {
		if result[idx] != math.MaxInt32 {
			continue
		}

		if array[idx] == maxVal {
			result[idx] = -1
		} else {
			for len(nextElements) > 1 && array[idx] >= nextElements[len(nextElements)-1] {
				nextElements = nextElements[:len(nextElements)-1]
			}

			if array[idx] < nextElements[len(nextElements)-1] {
				result[idx] = nextElements[len(nextElements)-1]
			}
		}
	}


	return result
}

func fillNextGreaterElement(array, nextElements, result []int) []int{
	smaller := make([]int, 0)

	for idx := 0; idx < len(array); idx++ {
		for len(nextElements) > 1 && array[idx] >= nextElements[len(nextElements)-1] {
			smaller = append(smaller, nextElements[len(nextElements)-1])
			nextElements = nextElements[:len(nextElements)-1]
		}

		if array[idx] < nextElements[len(nextElements)-1] {
			result[idx] = nextElements[len(nextElements)-1]
		}

		for len(smaller)-1 > idx {
			nextElements = append(nextElements, smaller[len(smaller)-1])
			smaller	= smaller[:len(smaller)-1]
		}
	}

	return smaller
}

func reverseSlice(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
