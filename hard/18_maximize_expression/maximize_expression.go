package maximize_expression
// important question

import (
	"math"
)

func MaximizeExpression(array []int) int {
	if len(array) < 4 {
		return 0
	}

	maxOfA := []int{array[0]}
	maxOfAMinusB := []int{math.MinInt32}
	maxOfAMinusBPlusC := []int{math.MinInt32, math.MinInt32}
	maxOfAMinusBPlusCMinusD := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	for idx := 1; idx < len(array); idx++ {
		currentMax := max(maxOfA[idx-1], array[idx])
		maxOfA = append(maxOfA, currentMax)
	}

	for idx := 1; idx < len(array); idx++ {
		currentMax := max(maxOfAMinusB[idx-1], maxOfA[idx-1]-array[idx])
		maxOfAMinusB = append(maxOfAMinusB, currentMax)
	}

	for idx := 2; idx < len(array); idx++ {
		currentMax := max(maxOfAMinusBPlusC[idx-1], maxOfAMinusB[idx-1]+array[idx])
		maxOfAMinusBPlusC = append(maxOfAMinusBPlusC, currentMax)
	}

	for idx := 3; idx < len(array); idx++ {
		currentMax := max(maxOfAMinusBPlusCMinusD[idx-1], maxOfAMinusBPlusC[idx-1]-array[idx])
		maxOfAMinusBPlusCMinusD = append(maxOfAMinusBPlusCMinusD, currentMax)
	}

	return maxOfAMinusBPlusCMinusD[len(maxOfAMinusBPlusCMinusD)-1]
}

// naive approach
func MaximizeExpression1(array []int) int {
	if len(array) < 4 {
		return 0
	}

	var maximumValueFound = math.MinInt32

	for a := range array {
		aValue := array[a]
		for b := a + 1; b < len(array); b++ {
			bValue := array[b]
			for c := b + 1; c < len(array); c++ {
				cValue := array[c]
				for d := c + 1; d < len(array); d++ {
					dValue := array[d]
					expressionValue := evaluateExpression(aValue, bValue, cValue, dValue)
					maximumValueFound = max(expressionValue, maximumValueFound)
				}
			}
		}
	}
	return maximumValueFound
}

func evaluateExpression(a, b, c, d int) int {
	return a - b + c - d
}

// 3, 6, 1, -3, 2, 7
// 6-(-3)+2-7
// 4

// 3 6 | 1 -3 2 7
// 3 6 1 | -3 2 7
// 3 6 1 -3 | 2 7

// 3, -1, 1, -1, -2, 4, 5, -4

//  0  4  2  4  5 -1 -1 9
// -7 -3 -5 -3 -2 -8 -9 0

// MaximizeExpression well i think my solution is better
// array[a]-array[b]+array[c]-array[d]
// array[a]-array[b]-(array[d]-array[c])
// a<b<c<d
func maximizeExpression(array []int) int {
	if len(array) < 4 {
		return 0
	}

	if len(array) == 4 {
		return array[0] - array[1] + array[2] - array[3]
	}

	maxSoFar := make([]int, len(array), len(array))
	minSoFar := make([]int, len(array), len(array))

	maxVal := math.MinInt32
	minVal := math.MaxInt32

	minValSoFar := math.MaxInt32

	for i, j := 1, len(array)-2; i < len(array) && j >= 0; i, j = i+1, j-1 {
		maxVal = max(maxVal, array[i-1])
		maxSoFar[i] = maxVal - array[i]

		minVal = min(minVal, array[j+1])
		minValSoFar = min(minValSoFar, minVal-array[j])
		minSoFar[j] = minValSoFar
	}

	maxVal = math.MinInt32
	for i := 1; i < len(array)-1; i++ {
		maxVal = max(maxVal, maxSoFar[i]-minSoFar[i+1])
	}
	return maxVal
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
