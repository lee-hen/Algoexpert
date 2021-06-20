package sorted_squared_array

import "sort"

// SortedSquaredArray
// Better solution
func SortedSquaredArray(array []int) []int {
	sortedSquares := make([]int, len(array))

	smallerValueIdx := 0
	largerValueIdx := len(array) - 1

	for idx := len(array) - 1; idx >= 0; idx-- {
		smallerValue := array[smallerValueIdx]
		largerValue := array[largerValueIdx]

		if abs(smallerValue) > abs(largerValue) {
			sortedSquares[idx] = smallerValue * smallerValue
			smallerValueIdx += 1
		} else {
			sortedSquares[idx] = largerValue * largerValue
			largerValueIdx -= 1
		}
	}

	return sortedSquares
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}


// My Solution 2
func sortedSquaredArray(array []int) []int {
	squaredArray := make([]int, 0, len(array))

	stack := &Stack{}
	for _, arr := range array {
		if arr < 0 {
			stack.push(arr)
		} else {
			for ; stack.size() != 0 && stack.peek() * (-1) < arr; {
				s := stack.pop()
				squaredArray = append(squaredArray, s * s)
			}

			squaredArray = append(squaredArray, arr * arr)

		}
	}

	for ; stack.size() != 0; {
		arr := stack.pop()
		squaredArray = append(squaredArray, arr * arr)
	}

	return squaredArray
}

type Stack []int

func (stack *Stack) push(arr int) {
	*stack = append(*stack, arr)
}

func (stack *Stack) pop() int {
	s := *stack
	last := s[len(s)-1]
	*stack = s[:len(s)-1]
	return last
}

func (stack *Stack) peek() int {
	s := *stack
	return s[len(s)-1]
}

func (stack *Stack) size() int {
	return len(*stack)
}

// My Solution 1
func sortedSquaredArray1(array []int) []int {
	squaredArray := make([]int, 0, len(array))
	needSort := false
	for _, arr := range array {
		if arr < 0 {
			needSort = true
		}
		squaredArray = append(squaredArray, arr * arr)
	}
	if needSort {
		sort.Ints(squaredArray)
	}

	return squaredArray
}
