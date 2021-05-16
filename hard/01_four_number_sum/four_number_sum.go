package four_number_sum

import (
	"sort"
)

func FourNumberSum(array []int, target int) [][]int {
	allPairSums := map[int][][]int{}
	quadruplets := make([][]int, 0)

	for i := 1; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			difference := target - currentSum

			if pairs, found := allPairSums[difference]; found {
				for _, pair := range pairs {
					newQuad := append(pair, array[i], array[j])
					quadruplets = append(quadruplets, newQuad)
				}
			}
		}

		for k := 0; k < i; k++ {
			currentSum := array[i] + array[k]
			if _, found := allPairSums[currentSum]; !found {
				allPairSums[currentSum] = [][]int{{array[k], array[i]}}
			} else {
				allPairSums[currentSum] = append(allPairSums[currentSum], []int{array[k], array[i]})
			}

		}
	}
	return quadruplets
}

// my solution
func fourNumberSum(slice []int, target int) [][]int {
	sort.Ints(slice)

	indices := make([][]int, 0)
	for i := 0; i < len(slice); i++ {
		for j := len(slice)-1; j > i+2; j-- {
			indices = append(indices, []int{i, j})
		}
	}

	result := make([][]int, 0)
	for _, index := range indices {
		i, j := index[0], index[1]
		left, right := i+1, j-1

		targetSum := target - slice[i] - slice[j]
		for left < right {
			if targetSum == slice[left] + slice[right] {
				result = append(result, []int{slice[i], slice[left], slice[right], slice[j]})
				left++
				right--
			} else if slice[left] + slice[right] < targetSum {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
