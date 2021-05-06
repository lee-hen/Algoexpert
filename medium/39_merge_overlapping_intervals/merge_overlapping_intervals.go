package merge_overlapping_intervals

import (
	"sort"
)

func MergeOverlappingIntervals(intervals [][]int) [][]int {
	sortedIntervals := make([][]int, len(intervals))
	copy(sortedIntervals, intervals)

	sort.Slice(sortedIntervals, func(i, j int) bool {
		return sortedIntervals[i][0] < sortedIntervals[j][0]
	})

	mergedIntervals := make([][]int, 0)
	currentInterval := sortedIntervals[0]
	mergedIntervals = append(mergedIntervals, currentInterval)

	for _, nextInterval := range sortedIntervals {
		currentIntervalEnd := currentInterval[1]
		nextIntervalStart, nextIntervalEnd := nextInterval[0], nextInterval[1]

		if currentIntervalEnd >= nextIntervalStart {
			currentInterval[1] = max(currentIntervalEnd, nextIntervalEnd)
		} else {
			currentInterval = nextInterval
			mergedIntervals = append(mergedIntervals, currentInterval)
		}
	}
	return mergedIntervals
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// {1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10}
// [[0 5] [1 2] [4 7] [6 8] [9 10]]

func mergeOverlappingIntervals(intervals [][]int) (overlappingIntervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for idx := 0; idx < len(intervals); idx++ {
		startInterval := intervals[idx][0]

		j := idx+1
		for j < len(intervals) {
			if intervals[idx][1] <= intervals[j][1] && intervals[idx][1] >= intervals[j][0] {
				idx = j
				j = idx+1
			} else if intervals[idx][1] > intervals[j][1] {
				j++
			} else if intervals[idx][1] < intervals[j][0]   {
				break
			}
		}

		endInterval := intervals[idx][1]
		overlappingIntervals = append(overlappingIntervals, []int{startInterval, endInterval})

		if j >= len(intervals) {
			break
		}
	}
	return
}
