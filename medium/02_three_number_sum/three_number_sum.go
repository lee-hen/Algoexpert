package three_number_sum

import (
	"sort"
)

// ThreeNumberSum
// Better solution
func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	triplets := make([][]int, 0)
	for i := 0; i < len(array)-2; i++ {
		left, right := i+1, len(array)-1
		for left < right {
			currentSum := array[i] + array[left] + array[right]
			if currentSum == target {
				triplets = append(triplets, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if currentSum < target {
				left += 1
			} else if currentSum > target {
				right -= 1
			}
		}
	}
	return triplets
}

func threeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	result := make([][]int, 0)
	if array[0] > target {
		return result
	}

	for i, number := range array {
		if number > target || i > len(array)-2 {
			break
		}

		nextIdx := i + 1
		otherIdx := len(array) - 1

		for i+1 < otherIdx {
			nextNumber := array[nextIdx]
			if target > number+nextNumber+array[otherIdx] {
				nextIdx++
				otherIdx = len(array) - 1

				if nextIdx > otherIdx {
					break
				}
				continue
			}

			if nextNumber >= array[otherIdx] {
				break
			}

			if target == number+nextNumber+array[otherIdx] {
				result = append(result, []int{number, nextNumber, array[otherIdx]})
			}

			otherIdx--
		}
	}

	return result
}
