package minimum_waiting_time

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	minWaitTime := make([]int, len(queries))
	for i := range queries {
		if i == 1 {
			minWaitTime[i] += queries[i - 1]
		}

		if i > 1 {
			minWaitTime[i] = minWaitTime[i - 1] + queries[i - 1]
		}
	}

	var min int
	for _, v := range minWaitTime {
		min += v
	}

	return min
}

