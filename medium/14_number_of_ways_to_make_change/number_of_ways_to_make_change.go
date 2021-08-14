package number_of_ways_to_make_change
// important question

// [1, 2, 3, 4] 6

// 0 1 2 3 4 5 6 amount
// ------------  /coins
// 1 0 0 0 0 0 0  []
// 1 1 1 1 1 1 1  [1]
// 1 1 2 2 3 3 4  [1,2]
// 1 1 2 3 4 5 7  [1,2,3]
// 1 1 2 3 5 6 9  [1,2,3,4]

// 1 1 1 1 1 1
// 2 1 1 1
// 2 2 2
// 2 2 1 1
// 3 3
// 3 1 1 1
// 3 2 1
// 4 2
// 4 1 1

// NumberOfWaysToMakeChange
// ?? why?
func NumberOfWaysToMakeChange(n int, coins []int) int {
	ways := make([]int, n+1, n+1)
	ways[0] = 1

	for _, coin := range coins {
		for amount := 1; amount <= n; amount++ {
			if coin <= amount {
				ways[amount] += ways[amount-coin]
			}
		}
	}
	return ways[n]
}
