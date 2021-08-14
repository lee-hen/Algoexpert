package max_path_sum
// important question

import (
	"math"
)

type BinaryTree struct {
	Value, MaxSumValue, maxSumValue int
	Left, Right *BinaryTree
}

func MaxPathSum(tree *BinaryTree) int {
	_, maxSum := findMaxSum(tree)
	return maxSum
}

func findMaxSum(tree *BinaryTree) (int, int) {
	if tree == nil {
		return 0, math.MinInt32
	}

	// branch 就是直线
	// root 就是三角
	leftMaxSumAsBranch, leftMaxPathSum := findMaxSum(tree.Left)
	rightMaxSumAsBranch, rightMaxPathSum := findMaxSum(tree.Right)
	maxChildSumAsBranch := max(leftMaxSumAsBranch, rightMaxSumAsBranch)

	value := tree.Value
	maxSumAsBranch := max(maxChildSumAsBranch+value, value)
	maxSumAsRootNode := max(leftMaxSumAsBranch+value+rightMaxSumAsBranch, maxSumAsBranch)
	maxPathSum := max(leftMaxPathSum, rightMaxPathSum, maxSumAsRootNode)

	return maxSumAsBranch, maxPathSum
}

// MaxPathSum1
// my solution
func MaxPathSum1(tree *BinaryTree) int {
	if tree == nil {
		return math.MinInt32
	}

	var maxSum int
	if tree.Left != nil && tree.Right != nil {
		maxLeft := max(maxPathSum(tree.Left), maxPathSum(tree.Left) + tree.Value)
		maxRight := max(maxPathSum(tree.Right), maxPathSum(tree.Right) + tree.Value)
		maxSum = max(max(maxLeft, maxRight), maxPathSum(tree.Left) + tree.Value + maxPathSum(tree.Right))
	} else if tree.Left != nil {
		maxSum = max(maxPathSum(tree.Left), maxPathSum(tree.Left) + tree.Value)
	} else if tree.Right != nil {
		maxSum = max(maxPathSum(tree.Right), maxPathSum(tree.Right) + tree.Value)
	} else {
		return tree.Value
	}

	if tree.MaxSumValue != 0 {
		return tree.MaxSumValue
	}

	tree.MaxSumValue = max(max(MaxPathSum(tree.Left), MaxPathSum(tree.Right)), maxSum)
	return tree.MaxSumValue
}

func maxPathSum(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}

	if tree.maxSumValue != 0 {
		return tree.maxSumValue
	}

	tree.maxSumValue = max(maxPathSum(tree.Left), maxPathSum(tree.Right)) + tree.Value
	return tree.maxSumValue
}

func max(first int, values ...int) int {
	for _, val := range values {
		if val > first {
			first = val
		}
	}
	return first
}
