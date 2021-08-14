package max_profit_with_k_transactions
// important question

import (
	"math"
	"sort"
)

// MaxProfitWithKTransactions
// O(nk) time | O(nk) space
func MaxProfitWithKTransactions(prices []int, k int) int {
	if len(prices) == 0 {
		return 0
	}
	profits := make([][]int, k+1)
	for i := range profits {
		profits[i] = make([]int, len(prices))
	}
	for t := 1; t < k+1; t++ {
		maxThusFar := math.MinInt32
		for d := 1; d < len(prices); d++ {
			maxThusFar = max(maxThusFar, profits[t-1][d-1]-prices[d-1])
			profits[t][d] = max(profits[t][d-1], maxThusFar+prices[d])
		}
	}
	return profits[k][len(prices)-1]
}

// MaxProfitWithKTransactionsOptimize
// # O(nk) time | O(n) space
func MaxProfitWithKTransactionsOptimize(prices []int, k int) int {
	if len(prices) == 0 {
		return 0
	}
	evenProfits := make([]int, len(prices))
	oddProfits := make([]int, len(prices))
	var currentProfits, previousProfits []int
	for t := 1; t < k+1; t++ {
		maxThusFar := math.MinInt32
		if t%2 == 1 {
			currentProfits, previousProfits = oddProfits, evenProfits
		} else {
			currentProfits, previousProfits = evenProfits, oddProfits
		}
		for d := 1; d < len(prices); d++ {
			maxThusFar = max(maxThusFar, previousProfits[d-1]-prices[d-1])
			currentProfits[d] = max(currentProfits[d-1], maxThusFar+prices[d])
		}
	}
	if k%2 == 0 {
		return evenProfits[len(prices)-1]
	}
	return oddProfits[len(prices)-1]
}

// k=2

//
//          0   1   2  3   4   5
// prices = [5, 11, 3, 50, 60, 90]
// 93
// Buy: 5, Sell: 11; Buy: 3, Sell: 90

// idx 1  5

//     0      1        2        3       4      5
//     5      11       3        50      60     90
// 	   0      0        0        0       0      0
//     0      6        6        47      57     87
//     0      6        6        53      63     93


// 0 4 6 8 10

//      0  1   2   3   4   5   6   7   8   9   10
//      1  25  24  23  12  36  14  40  31  41  5
//  0   0  0   0   0   0   0   0   0   0   0   0
//  1   0  24  24  24  24  35  35  39  39  40  40 (minIdx -1 minVal MaxInt32)
//  2   0  24  24  24  24  48  48  61  61  62  62 (minIdx 0  minVal 0)
//  3   0  24  24  24  24  48  48  74  74  75  75 (minIdx 6  minVal 14)
//  4   0  24  24  24  24  48  48  74  74  84  84 (minIdx 8  minVal 31)

// maxProfitWithKTransactions
// O(kn^2) time | O(nk)
func maxProfitWithKTransactions(prices []int, k int) int {

	minIndices := make([]int, 0)
	for i := 0; i < len(prices); i++ {
		if i == 0 && prices[i+1] > prices[i] {
			minIndices = append(minIndices, i)
		} else if i == len(prices)-1 && prices[i-1] > prices[i]{
			minIndices = append(minIndices, i)
		} else if i > 0 && prices[i-1] > prices[i] && prices[i+1] > prices[i] {
			minIndices = append(minIndices, i)
		}
	}

	sort.Ints(minIndices)

	profits := make([][]int, k+1, k+1)
	for i := range profits {
		profits[i] = make([]int, len(prices))
	}

	for j := 1; j <= k; j++ {
		for i, price := range prices {
			var maxProfit int
			if price > prices[minIndices[0]] {
				for _, minIdx := range minIndices {
					if i > minIdx {
						currProfit := price - prices[minIdx] + profits[j-1][minIdx]
						if currProfit > maxProfit {
							maxProfit = currProfit
						}
					} else {
						break
					}
				}
			}

			if i > 0 {
				profits[j][i] = max(maxProfit, profits[j][i-1])
			}
		}
	}


	//for j := 1; j <= k; j++ {
	//	for i, price := range prices {
	//		var maxProfit = math.MinInt32
	//
	//      maxProfit = max(maxProfit, profits[j-1][i-1]-prices[i-1])
	//		for minIdx := 0; minIdx < i; minIdx++ {
	//			currProfit := profits[j-1][minIdx]-prices[minIdx]
	//			if currProfit > maxProfit {
	//				maxProfit = currProfit
	//			}
	//		}
	//
	//		if i > 0 {
	//			profits[j][i] = max(maxProfit+price, profits[j][i-1])
	//		}
	//	}
	//}

	return profits[k][len(prices)-1]
}

func max(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}
