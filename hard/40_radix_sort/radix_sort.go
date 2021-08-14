package radix_sort
// important question

import (
	"strconv"
)

// RadixSort
// O(d * (n + b)) time | O(n + b) space - where n is the length of the input array,
// d is the max number of digits, and b is the base of the numbering system used where here b is 10
func RadixSort(array []int) []int {

	if len(array) == 0 {
		return array
	}

	maxNumber := max(array)

	digit := 0
	for maxNumber/pow(10, digit) > 0 {
		countingSort(array, digit)
		digit += 1
	}

	return array
}

func countingSort(array []int, digit int) {
	sortedArray := make([]int, len(array))
	countArray := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	digitColumn := pow(10, digit)
	for _, num := range array {
		countIndex := (num / digitColumn) % 10
		countArray[countIndex] += 1
	}

	for idx := 1; idx < 10; idx++ {
		countArray[idx] += countArray[idx-1]
	}

	for idx := len(array) - 1; idx >= 0; idx-- {
		countIndex := (array[idx] / digitColumn) % 10
		countArray[countIndex] -= 1
		sortedIndex := countArray[countIndex]
		sortedArray[sortedIndex] = array[idx]
	}

	for idx := range array {
		array[idx] = sortedArray[idx]
	}
}


// radixSort
// my solution
// 8762, 654, 3008, 345, 87, 65, 234, 12, 2
func radixSort(array []int) []int {
	maxVal := max(array)
	maxStr :=  strconv.Itoa(maxVal)

	for k := 1; k <= len(maxStr); k++ {
		array = countSortWithKthDigit(k, array, func(number, k int) int {
			var digitNumber int
			for i := 0; i < k; i++ {
				digitNumber = number%10
				number = number/10
			}
			return digitNumber
		})
	}

	return array
}

func countSortWithKthDigit(k int, array []int,  digitNumber func(number, k int) int) []int {
	counts := make([]int, 10, 10)
	for _, el := range array {
		counts[digitNumber(el, k)]++
	}

	//for i := 0; i < len(counts)-1; i++ {
	//	counts[i+1] += counts[i]
	//}

	//
	//sorted := make([]int, len(array), len(array))
	//for i := len(array)-1; i >= 0; i-- {
	//	counts[digitNumber(array[i], k)]--
	//	position := counts[digitNumber(array[i], k)]
	//	sorted[position] = array[i]
	//}

	l := len(array)
	for i := len(counts)-1; i >= 0; i-- {
		l -= counts[i]
		counts[i] = l
	}

	sorted := make([]int, len(array), len(array))
	for i := 0; i < len(array); i++{
		position := counts[digitNumber(array[i], k)]
		sorted[position] = array[i]
		counts[digitNumber(array[i], k)]++
	}

	return sorted
}

func max(array []int) int {
	currentMax := array[0]
	for _, element := range array {
		if currentMax < element {
			currentMax = element
		}
	}
	return currentMax
}

func pow(a int, power int) int {
	var result = 1
	for i := 0; i < power; i++ {
		result *= a
	}
	return result
}
