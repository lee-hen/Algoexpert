package build_failures

// BuildFailures
// O(nlog(m)) time | O(1) space - where n is the number of
// build runs and m is the length of the longest build run
func BuildFailures(buildRuns [][]bool) int {
	longestLength := 1
	currentLongestLength := 1
	previousGreenPercentage := calculateGreenPercentage(buildRuns[0])

	for i := 1; i < len(buildRuns); i++ {
		currentGreenPercentage := calculateGreenPercentage(buildRuns[i])
		if currentGreenPercentage < previousGreenPercentage {
			currentLongestLength++
			if longestLength < currentLongestLength {
				longestLength = currentLongestLength
			}
		} else {
			currentLongestLength = 1
		}
		previousGreenPercentage = currentGreenPercentage
	}

	if longestLength > 1 {
		return longestLength
	}
	return -1
}

func calculateGreenPercentage(buildRun []bool) float64 {
	firstFalseIdx := binarySearchForFirstFalse(buildRun)
	return float64(firstFalseIdx) / float64(len(buildRun))
}


func getLongestDecreasingSubarrayLength(array []float64) int {
	longestLength := 1
	currentLongestLength := 1
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			currentLongestLength++
			if currentLongestLength > longestLength {
				longestLength = currentLongestLength
			}
		} else {
			currentLongestLength = 1
		}
	}

	if longestLength > 1 {
		return longestLength
	}
	return -1
}


//buildRuns = [
//	[true, true, true, false, false],  60%
//	[true, true, true, true, false],   80%
//	[true, true, true, true, true, true, false, false, false], -66.66%
//	[true, false, false, false, false, false], -16.66%
//	[true, true, true, true, true, true, true, true, true, true, true, true, false], 92.30%
//	[true, false], -50%
//	[true, true, true, true, false, false],  66.66%
//]

// buildFailures
// my solution
func buildFailures(buildRuns [][]bool) int {
	greensPercentage := make([]float64, len(buildRuns), len(buildRuns))


	for idx, buildRun := range buildRuns {
		var greens int

		// You can use binary search to optimally find the first false in each build run and to optimally calculate each build run's green percentage.
		greens = binarySearchForFirstFalse(buildRun)

		//for _, result := range buildRun {
		//	if result {
		//		greens++
		//	}
		//}

		greensPercentage[idx] = float64(greens) / float64(len(buildRun)) * 100
	}

	var numberOfDecreasing = 1
	var greatest int
	for i := 0; i < len(greensPercentage)-1; i++ {
		if greensPercentage[i] > greensPercentage[i+1] {
			numberOfDecreasing++
		} else {
			greatest = max(greatest, numberOfDecreasing)
			numberOfDecreasing = 1
		}
	}

	greatest = max(greatest, numberOfDecreasing)

	if greatest == 0 || greatest == 1{
		return -1
	}


	return greatest
}

func max(a int, others ...int) int {
	m := a
	for _, other := range others {
		if other > m {
			m = other
		}
	}
	return m
}

// Iterative Binary Search.
func binarySearchForFirstFalse(array []bool) int {
	leftIdx := 0
	rightIdx := len(array) - 1

	for leftIdx <= rightIdx {
		middleIdx := leftIdx + ((rightIdx - leftIdx) / 2)
		isFalse := !array[middleIdx]
		if !isFalse {
			leftIdx = middleIdx + 1
			continue
		}
		isFirstFalse := middleIdx == 0 || array[middleIdx-1]
		if isFirstFalse {
			return middleIdx
		}
		rightIdx = middleIdx - 1
	}

	return -1
}
