package non_constructible_change

import (
	"sort"
)

// My correct Solution
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

// not corrected
func nonConstructibleChange1(coins []int) int {
	sort.Ints(coins)

	max := coins[len(coins) - 1]

	var notConstructible int
	for _, c := range coins[:len(coins) - 1] {
		notConstructible += c
	}

	notConstructible++
	if  max <= notConstructible {
		notConstructible = max + 1
	}

	return notConstructible
}
