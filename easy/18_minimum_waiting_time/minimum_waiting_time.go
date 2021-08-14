package minimum_waiting_time
// important question

import "sort"

func MinimumWaitingTime(queries []int) int {
	sort.Ints(queries)

	var total int
	for idx, duration := range queries {
		queriesLeft := len(queries) - (idx + 1)
		total += queriesLeft * duration
	}
	return total
}

// My Solution
func minimumWaitingTime1(queries []int) int {
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
