package knapsack_problem
// important question

// 11
// [1, 2], [4, 3], [5, 6], [6, 7]
// 11 [1, 2], [4, 3], [5, 6]


//0  1  2  3  4  5    6     7    8   9     10
//0  0  0  0  0  0    0     0    0   0     0    []
//0  0  0  0  0  0    0     6    6   6     6    [6, 7]
//0  0  1  1  1  1    1     6    6   6     6    [1, 2]
//0  0  1  4  4  5    5     6    6   6    10    [4, 3]
//0  0  1  4  4  5    5     6    6   9    10    [5, 6]


//0  1  2  3  4  5    6     7    8   9     10   11
//0  0  0  0  0  0    0     0    0   0     0    0   []      0
//0  0  1  1  1  1    1     1    1   1     1    1   [1, 2]  1
//0  0  1  4  4  5    5     5    5   5     5    5   [4, 3]  2
//0  0  1  4  4  5    5     5    6   9     9   10   [5, 6]  3
//0  0  1  4  4  5    5     6    6   9    10   10   [6, 7]  4

// KnapsackProblem
// my solution is the better solution
func KnapsackProblem(items [][]int, capacity int) []interface{} {
	knapsack := make([][]int, len(items)+1, len(items)+1)
	for i := range knapsack {
		knapsack[i] = make([]int, capacity+1, capacity+1)
	}

	for i := 1; i <= len(items); i++  {
		item := items[i-1]

		for cap := 1; cap <= capacity; cap++ {
			if item[1] > cap {
				knapsack[i][cap] = knapsack[i-1][cap]
			} else {
				knapsack[i][cap] = max(knapsack[i-1][cap], item[0] + knapsack[i-1][cap-item[1]])
			}
		}
	}

	return []interface{}{
		knapsack[len(knapsack)-1][capacity],          // total value
		buildKnapsackItemIndices(knapsack, items),    // item indices
	}
}

func buildKnapsackItemIndices(knapsack, items [][]int) []int {
	knapsackItemIndices := make([]int, 0)
	i := len(knapsack)-1
	j := len(knapsack[len(knapsack)-1])-1

	for i > 0 && j > 0 {
		if knapsack[i][j] == knapsack[i-1][j] {
			i--
		} else {
			knapsackItemIndices = append(knapsackItemIndices, i-1)

			item := items[i-1]
			if j >= item[1] {
				j = j - item[1]
				i--
			}
		}
	}

	ReverseSlice(knapsackItemIndices)
	return knapsackItemIndices
}

func ReverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
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
