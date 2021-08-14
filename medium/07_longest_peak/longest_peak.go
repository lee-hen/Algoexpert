package longest_peak
// important question

func LongestPeak(array []int) int {
	longestPeakLength := 0
	i := 1
	for i < len(array)-1 {
		isPeak := array[i-1] < array[i] && array[i] > array[i+1]
		if !isPeak {
			i += 1
			continue
		}
		leftIdx := i - 2
		for leftIdx >= 0 && array[leftIdx] < array[leftIdx+1] {
			leftIdx -= 1
		}
		rightIdx := i + 2
		for rightIdx < len(array) && array[rightIdx] < array[rightIdx-1] {
			rightIdx += 1
		}
		currentPeakLength := rightIdx - leftIdx - 1
		if currentPeakLength > longestPeakLength {
			longestPeakLength = currentPeakLength
		}
		i = rightIdx
	}
	return longestPeakLength
}

func longestPeak(array []int) int {
	if len(array) < 3 {
		return 0
	}

	var longestPeak int
	longestPeaks := make([]int, 0)
	for i :=0; i < len(array)-2; {
		if array[i] < array[i+1] {
			longestPeaks = append(longestPeaks, array[i])
			downPeaks, idx := traverseDownFromPeek(i, array)
			if len(downPeaks) > 0 {
				longestPeaks = append(longestPeaks, downPeaks...)
				if longestPeak < len(longestPeaks) {
					longestPeak = len(longestPeaks)
				}
				longestPeaks = make([]int, 0)
				i = idx
			} else {
				i++
			}
		} else {
			if longestPeak < len(longestPeaks) {
				longestPeak = len(longestPeaks)
			}
			longestPeaks = make([]int, 0)
			i++
		}
	}
	if longestPeak > 2 {
		return longestPeak
	}
	return 0
}

func traverseDownFromPeek(i int , array []int) ([]int, int) {
	longestPeaks := make([]int, 0)
	var idx int
	if array[i+1] > array[i+2] {
		longestPeaks = append(longestPeaks, []int{array[i+1], array[i+2]}...)
		idx = i+2
		for j := i+3; j < len(array); j++ {
			if array[j-1] <= array[j] {
				break
			}
			longestPeaks = append(longestPeaks, array[j])
			idx = j+1
		}
	}
	return longestPeaks, idx
}