package smallest_difference
// important question

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)
	idxOne, idxTwo := 0, 0
	smallest, current := math.MaxInt32, math.MaxInt32
	smallestPair := []int{}
	for idxOne < len(array1) && idxTwo < len(array2) {
		first, second := array1[idxOne], array2[idxTwo]
		if first < second {
			current = second - first
			idxOne += 1
		} else if second < first {
			current = first - second
			idxTwo += 1
		} else {
			return []int{first, second}
		}
		if smallest > current {
			smallest = current
			smallestPair = []int{first, second}
		}
	}
	return smallestPair
}

func smallestDifference(array1, array2 []int) []int {

	redArray := make(map[int]struct{})
	blackArray := make(map[int]struct{})

	for _, arr := range array1 {
		redArray[arr] = struct{}{}
	}

	for _, arr := range array2 {
		blackArray[arr] = struct{}{}
	}

	array := append(array1, array2...)
	sort.Ints(array)

	smallestDifference := make([]int, 2, 2)

	smallest := array2[0] - array1[0]
	for i, arr := range array {
		if i > len(array) - 2 {
			break
		}
		nextArr := array[i+1]
		_, ok := redArray[arr]
		if ok {
			if _, ok = blackArray[nextArr]; !ok {
				continue
			}

			if abs(nextArr - arr) <= abs(smallest) {
				smallest = nextArr - arr
				smallestDifference[0] = arr
				smallestDifference[1] = nextArr
			}
		}

		_, ok = blackArray[arr]
		if ok {
			if _, ok = redArray[nextArr]; !ok {
				continue
			}

			if abs(nextArr - arr) <= abs(smallest) {
				smallest = nextArr - arr
				smallestDifference[0] = nextArr
				smallestDifference[1] = arr

			}
		}
	}

	return smallestDifference

}

func abs(val int) int {
	if val < 0 {
		return -val
	}

	return val
}



