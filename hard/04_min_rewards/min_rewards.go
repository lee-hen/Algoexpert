package min_rewards

func MinRewards(scores []int) int {
	rewards := make([]int, len(scores))
	fill(rewards, 1)
	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			rewards[i] = rewards[i-1] + 1
		}
	}
	for i := len(scores) - 2; i >= 0; i-- {
		if scores[i] > scores[i+1] {
			rewards[i] = max(rewards[i], rewards[i+1]+1)
		}
	}
	return sum(rewards)
}

func MinRewards1(scores []int) int {
	rewards := make([]int, len(scores))
	fill(rewards, 1)
	localMinIndices := getLocalMinIndices(scores)
	for _, localMinIdx := range localMinIndices {
		expandFromLocalMinIdx(localMinIdx, scores, rewards)
	}
	return sum(rewards)
}

func getLocalMinIndices(arr []int) []int {
	localMinIndices := make([]int, 0)
	if len(arr) == 1 {
		localMinIndices = append(localMinIndices, 0)
		return localMinIndices
	}

	for i := 0; i < len(arr); i++ {
		if i == 0 && arr[i] < arr[i+1] {
			localMinIndices = append(localMinIndices, i)
		}
		if i == len(arr)-1 && arr[i] < arr[i-1] {
			localMinIndices = append(localMinIndices, i)
		}
		if i == 0 || i == len(arr)-1 {
			continue
		}
		if arr[i] < arr[i+1] && arr[i] < arr[i-1] {
			localMinIndices = append(localMinIndices, i)
		}
	}
	return localMinIndices
}

func expandFromLocalMinIdx(localMinIdx int, scores, rewards []int) {
	leftIdx := localMinIdx - 1
	for leftIdx >= 0 && scores[leftIdx] > scores[leftIdx+1] {
		rewards[leftIdx] = max(rewards[leftIdx], rewards[leftIdx+1]+1)
		leftIdx--
	}

	rightIdx := localMinIdx + 1
	for rightIdx < len(scores) && scores[rightIdx] > scores[rightIdx-1] {
		rewards[rightIdx] = rewards[rightIdx-1] + 1
		rightIdx++
	}
}

func fill(arr []int, val int) {
	for i := range arr {
		arr[i] = val
	}
}

func sum(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// MinRewards
// all students must receive one reward
// reward are all unique score is different
// scores:  8, 4, 2, 1, 3, 6, 7, 9, 5
// rewards: 4  3  2  1  2  3  4  5  1
// 25
func minRewards(scores []int) int {
	if len(scores) == 1 {
		return 1
	}

	peeks := make(map[int][]int)
	var i int
	for i < len(scores)-1 {
		if scores[i] < scores[i+1] {
			left, peek := i, i
			for peek < len(scores)-1 && scores[peek] < scores[peek+1] {
				peek++
			}

			right := peek
			for right < len(scores)-1 && scores[right] > scores[right+1] {
				right++
			}
			i = right

			peeks[peek] = append([]int{}, []int{left, right}...)
		} else {
			peek, right := i, i
			for right < len(scores)-1 && scores[right] > scores[right+1] {
				right++
			}
			i = right
			peeks[peek] = append([]int{}, []int{peek, right}...)
		}
	}

	ranges := make(map[int]int)
	var minRewards int
	for k, val := range peeks {
		left := val[0]
		peek := k
		right := val[1]

		ranges[left]++
		ranges[right]++

		if right-peek > peek-left {
			peekReward := right - peek + 1
			rightReward := 1
			minRewards += (peekReward + rightReward) * peekReward / 2

			maxLeftReward := peek - left
			leftReward := 1
			minRewards += (leftReward + maxLeftReward) * maxLeftReward / 2
		} else {
			peekReward := peek - left + 1
			leftReward := 1
			minRewards += (peekReward + leftReward) * peekReward / 2

			maxRightReward := right - peek
			rightReward := 1
			minRewards += (rightReward + maxRightReward) * maxRightReward / 2
		}
	}

	for _, v := range ranges {
		if v > 1 {
			minRewards -= 1
		}
	}

	return minRewards
}
