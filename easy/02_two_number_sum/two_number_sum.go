package two_number_sum

import (
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right]}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{}
}

// My solution
func twoNumberSum1(array []int, target int) []int {
	for i, x := range array {
		for _, y := range array[i+1:] {
			if x+y == target {
				return []int{x, y}
			}
		}
	}
	return []int{}
}

func twoNumberSum2(array []int, target int) []int {
	memorize := make(map[int]struct{})

	for _, y := range array {
		x := target - y

		if _, ok := memorize[x]; ok {
			return []int{x, y}
		} else {
			memorize[y] = struct{}{}
		}
	}
	return []int{}
}
