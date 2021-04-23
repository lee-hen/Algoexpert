package kadanes_algorithm

import "math"

//KadanesAlgorithm
// input [3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]
// output 19 [1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1]
// better solution
func KadanesAlgorithm(array []int) int {
	maxEndingHere := array[0]
	maxSoFar := array[0]
	for i := 1; i < len(array); i++ {
		num := array[i]
		maxEndingHere = max(num, maxEndingHere+num)
		maxSoFar = max(maxSoFar, maxEndingHere)
	}
	return maxSoFar
}

//KadanesAlgorithm
// input [3, 5, -9, 1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1, -5, 4]
// output 19 [1, 3, -2, 3, 4, 7, 2, -9, 6, 3, 1]
// my solution
func kadanesAlgorithm(array []int) int {
	var current = math.MinInt32
	var maxValue = math.MinInt32
	for _, arr := range array {
		current += arr
		if current < arr {
			current = arr
		}
		maxValue = max(maxValue, current)
	}
	return maxValue
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
