package min_number_of_coins_for_change
// important question

import (
	"math"
)

// [1, 2, 3, 4] 6

//  0  1  2  3  4  5  6  amount
//  -------------------  /coins
//  0      math.Max      []
//  0  1  2  3  4  5  6  [1]
//  0  1  1  2  2  3  3  [2]
//  0  1  1  1  2  2  2  [3]
//  0  1  1  1  1  2  2  [4]

// MinNumberOfCoinsForChange
// my solution
func MinNumberOfCoinsForChange(n int, coins []int) int {
	if n == 0 {
		return 0
	}

	coinNumbers := make([]int, n+1, n+1)
	for i := 1; i <=n; i++ {
		coinNumbers[i] = math.MaxUint32
	}

	for _, coin := range coins {
		for amount := 1; amount <= n; amount++ {
			if coin == amount {
				coinNumbers[amount] = 1
			} else if coin < amount {
				coinNumbers[amount] = min(coinNumbers[amount-coin] + 1, coinNumbers[amount])
			}
		}
	}
	if coinNumbers[n] == math.MaxUint32 {
		return -1
	}
	return coinNumbers[n]
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

//
//func min(x, y int) int {
//	if x < y {
//		return x
//	}
//	return y
//}
