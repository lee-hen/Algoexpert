package min_number_of_jumps
// important question

import "math"

func MinNumberOfJumps(array []int) int {
	jumps := make([]int, len(array))
	for i := range jumps {
		jumps[i] = math.MaxInt32
	}
	jumps[0] = 0

	// check all the index before i, find minimum jumps to get the current i
	// array[j]+j >= i means the j index can jump to the current i
	for i := 1; i < len(array); i++ {
		for j := 0; j < i; j++ {
			if array[j]+j >= i {
				jumps[i] = min(jumps[j]+1, jumps[i])
			}
		}
	}
	return jumps[len(array)-1]
}

func MinNumberOfJumps1(array []int) int {
	if len(array) == 1 {
		return 0
	}

	jumps := 0
	maxReach := array[0]
	steps := array[0]

	for i := 1; i < len(array)-1; i++ {
		if i+array[i] > maxReach {
			maxReach = i + array[i]
		}

		steps -= 1

		if steps == 0 {
			jumps += 1
			steps = maxReach - i
		}
	}

	return jumps + 1
}

// 3->4
// 4->2 4->3
// 2->7 3->7
// 7->3

// 3  5  5  5  6  8 10 10 10 10 10
// 3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3
//    _           _  _           _

// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11
//    3, 4, 2, 1, 2, 3, 7, 1, 1, 1, 3
//    3  5  5  5  6  8  13

// 0  1  2  3  4  5  6
//    2, 1, 2, 2, 1, 1, 1
//    2  2  4  5  5  6  6
// my solution
func minNumberOfJumps(array []int) int {
	if len(array) == 1 {
		return 0
	}

	var numberOfJumps, previousJump, currentJump int
	for idx := 1; idx <= len(array); idx++ {
		currentJump = idx-1 + array[idx-1]
		if currentJump > previousJump {
			numberOfJumps++
		}

		if currentJump >= len(array)-1 {
			break
		}

		if idx > 2 && currentJump == previousJump+1 && array[idx-1] > array[idx-2] {
			numberOfJumps--
		}

		previousJump = max(currentJump, previousJump)
	}

	return numberOfJumps
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
