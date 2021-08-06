package max_subsequence_dot_product

import (
	"math"
)

// MaxSubsequenceDotProduct
// O(n * m) time | O(n * m) space - where n is the length of the
// first input array and m is the length of the second input array
func MaxSubsequenceDotProduct(arrayOne, arrayTwo []int) int {
	maxDotProducts := initializeDotProducts(arrayOne, arrayTwo)
	for i := 1; i <= len(arrayOne); i++ {
		for j := 1; j <= len(arrayTwo); j++ {
			currentProduct := arrayOne[i-1] * arrayTwo[j-1]
			maxDotProducts[i][j] = max(
				currentProduct,
				maxDotProducts[i-1][j-1]+currentProduct,
				maxDotProducts[i-1][j-1], // なくてもいけます。
				maxDotProducts[i-1][j],
				maxDotProducts[i][j-1],
			)
		}
	}
	return maxDotProducts[len(arrayOne)][len(arrayTwo)]
}

func initializeDotProducts(arrayOne, arrayTwo []int) [][]int {
	dotProducts := make([][]int, 0)
	for i := 0; i <= len(arrayOne); i++ {
		row := make([]int, 0)
		for j := 0; j <= len(arrayTwo); j++ {
			row = append(row, math.MinInt32)
		}
		dotProducts = append(dotProducts, row)
	}
	return dotProducts
}

//arrayOne = [4, 7, 9, -6, 6]
//arrayTwo = [5, 1, -1, -3, -2, -10]
// 105  From the subsequences [9, -6] and [5, -10].

// index2
// 5 0

// index1
// 3 2

//  0   1    2    3    4     5   index2/
//  5   1   -1   -3   -2   -10      index1
//  -   -   -     -   -     -
//- 20   20  20   20   20    20    4  0
//- 35   35  35   35   35    35    7  1
//- 45   45  45   45   45    45    9  2
//- 45   45  51   63   63    105  -6  3
//- 45   51  51   63   63    105   6  4


// [3, -2],
// [2, -2, 3]

//    2  -2    3
//-   -   -    -
//-   6   6    9   3
//-  -4  10   10  -2

// my solution
func maxSubsequenceDotProduct(arrayOne, arrayTwo []int) int {
	dotProduct := make([][]int, len(arrayOne)+1, len(arrayOne)+1)
	for i := range dotProduct {
		dotProduct[i] = make([]int, len(arrayTwo)+1, len(arrayTwo)+1)
		dotProduct[i][0] = math.MinInt32
	}

	for i := range dotProduct[0] {
		dotProduct[0][i] = math.MinInt32
	}

	for i, num1 := range arrayOne {
		for j, num2 := range arrayTwo {
			dotProduct[i+1][j+1] = max(max(max(dotProduct[i+1][j], num1 * num2), dotProduct[i][j+1]), dotProduct[i][j] + num1 * num2)
		}
	}

	return dotProduct[len(arrayOne)][len(arrayTwo)]
}

func max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}
