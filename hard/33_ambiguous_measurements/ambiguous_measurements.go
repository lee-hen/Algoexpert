package ambiguous_measurements
// important question

import (
	"fmt"
	"sort"
)

// AmbiguousMeasurements
// O(low * high * n) time | O(low * high) space - where n is the number of measuring cups
func AmbiguousMeasurements(measuringCups [][]int, low int, high int) bool {
	memoization := map[string]bool{}
	return canMeasureInRange(measuringCups, low, high, memoization)
}

func canMeasureInRange(measuringCups [][]int, low, high int, memoization map[string]bool) bool {
	memoizeKey := createHashtableKey(low, high)
	if val, found := memoization[memoizeKey]; found {
		return val
	}

	if low <= 0 && high <= 0 {
		return false
	}

	canMeasure := false

	for _, cup := range measuringCups {
		if low <= cup[0] && cup[1] <= high {
			canMeasure = true
			break
		}

		newLow := max(0, low-cup[0])
		newHigh := max(0, high-cup[1])
		canMeasure = canMeasureInRange(measuringCups, newLow, newHigh, memoization)
		if canMeasure {
			break
		}
	}

	memoization[memoizeKey] = canMeasure
	return canMeasure
}

func createHashtableKey(low, high int) string {
	return fmt.Sprintf("%d:%d", low, high)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ambiguousMeasurements
// my solution
func ambiguousMeasurements(measuringCups [][]int, low int, high int) bool {
	sort.Slice(measuringCups, func(i, j int) bool {
		return measuringCups[i][0] < measuringCups[j][0]
	})

	return ambiguousMeasure(len(measuringCups)-1, 0, 0, low, high, measuringCups)
}

func ambiguousMeasure(cupIdx, currLow, currHigh, low, high int, measuringCups [][]int) bool {
	if cupIdx < 0 {
		return false
	}

	for belowSpecifiedRange(currLow, low) {
		currLow += measuringCups[cupIdx][0]
		currHigh += measuringCups[cupIdx][1]
	}

	if measurable(currLow, currHigh, low, high) {
		return true
	}

	currLow -= measuringCups[cupIdx][0]
	currHigh -= measuringCups[cupIdx][1]

	for belowSpecifiedRange(currLow, low) {
		currLow += measuringCups[0][0]
		currHigh += measuringCups[0][1]
	}

	if measurable(currLow, currHigh, low, high) {
	 	return true
	 }

	return ambiguousMeasure(cupIdx-1, 0, 0, low, high, measuringCups)
}

func belowSpecifiedRange(currLow, low int) bool {
	return currLow < low
}

func measurable(l, h, low, high int) bool {
	return l <= h && l >= low && l <= high && h <= high
}
