package non_constructible_change
// important question

import (
	"sort"
)

// NonConstructibleChange
func NonConstructibleChange(coins []int) int {
	if len(coins) == 0 {
		return 1
	}
	sort.Ints(coins)

	var total int
	for i, c := range coins {
		if i > 0 {
			total += coins[i - 1]
		}

		if c > total + 1 {
			return total + 1
		}
	}

	return total + coins[len(coins)-1] + 1
}
