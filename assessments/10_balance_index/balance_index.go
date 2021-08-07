package balance_index

// array = [0, 9, -8, 2, 7, 1, 11, -2, 1]
// 5        0 + 9 + -8 + 2 + 7 == 11 + -2 + 1

// 0    1   2   3   4   5
// [0,  9,  1,  3,  10, 11, 22, 20, 21]
// [30  30, 21, 20, 18, 11, 10, -1, 1]

// BalanceIndex
// O(n) time | O(1) space - where n is the length of the input array
func BalanceIndex(array []int) int {
	arraySum := 0
	for _, n := range array {
		arraySum += n
	}

	leftSideSum, rightSideSum := 0, arraySum
	for i := 0; i < len(array); i++ {
		rightSideSum -= array[i]
		if leftSideSum == rightSideSum {
			return i
		}
		leftSideSum += array[i]
	}
	return -1
}

// my solution

func balanceIndex(array []int) int {
	sum := make([]int, len(array), len(array))

	for i, num := range array {
		if i == 0 {
			sum[0] = num
			continue
		}

		sum[i] = sum[i-1] + num
	}

	if array[len(array)-1] == sum[len(array)-1]  {
		return len(array)-1
	}

	prevSum := array[len(array)-1]
	for j := len(array)-2; j >= 0; j-- {
		currSum := prevSum + array[j]
		if currSum == sum[j] {
			return j
		}
		prevSum = currSum
	}
	return -1
}
