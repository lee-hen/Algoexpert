package largest_bst_size
// important question
// similar question github.com/lee-hen/Algoexpert/medium/09_validate_bst

import (
	"math"
)

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// LargestBSTSize
// Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree
func LargestBSTSize(tree *BinaryTree) int {
	return getBinaryTreeInfo(tree).runningLargestBstSize
}

type treeInfo struct {
	isBst                 bool
	maxValue              int
	minValue              int
	runningLargestBstSize int
	treeSize              int
}

func getBinaryTreeInfo(tree *BinaryTree) treeInfo {
	if tree == nil {
		return treeInfo{
			isBst:                 true,
			maxValue:              math.MinInt32,
			minValue:              math.MaxInt32,
			runningLargestBstSize: 0,
			treeSize:              0,
		}
	}

	leftTreeInfo := getBinaryTreeInfo(tree.Left)
	rightTreeInfo := getBinaryTreeInfo(tree.Right)

	treeSize := 1 + leftTreeInfo.treeSize + rightTreeInfo.treeSize

	satisfiesBstProp := tree.Value > leftTreeInfo.maxValue && tree.Value <= rightTreeInfo.minValue
	isBst := satisfiesBstProp && leftTreeInfo.isBst && rightTreeInfo.isBst

	maxValue := max(tree.Value, leftTreeInfo.maxValue, rightTreeInfo.maxValue)
	minValue := min(tree.Value, leftTreeInfo.minValue, rightTreeInfo.minValue)

	runningLargestBstSize := 0
	if isBst {
		runningLargestBstSize = treeSize
	} else {
		runningLargestBstSize = max(leftTreeInfo.runningLargestBstSize, rightTreeInfo.runningLargestBstSize)
	}

	return treeInfo{
		isBst:                 isBst,
		maxValue:              maxValue,
		minValue:              minValue,
		runningLargestBstSize: runningLargestBstSize,
		treeSize:              treeSize,
	}
}

func min(a int, others ...int) int {
	m := a
	for _, other := range others {
		if other < m {
			m = other
		}
	}
	return m
}

func max(a int, others ...int) int {
	m := a
	for _, other := range others {
		if other > m {
			m = other
		}
	}
	return m
}


//
//　　　　　　　　20
//　　　　/    　  　        \
//　　　7                 　 10
//　　/  \              　/      \
//　 0    8  　  　      4     　 15
//　　　 /   \   　    /   \   　/   \
//　　　7　    9  　　2      5 　13   22
//　　　　　　　　　　/      /     \
//　　　　　　　　　1      50       14
//                     /
//                    60

// my solution easy understand.
func largestBSTSize(tree *BinaryTree) int {
	var validateSize int
 	getBinaryTreeSize(tree, &validateSize)

	return validateSize
}

// bool(isBst), int(Tree Size)
func getBinaryTreeSize(tree *BinaryTree, validateSize *int) (bool, int) {
	if tree == nil {
		return true, 0
	}

	validateLeft, leftSize := getBinaryTreeSize(tree.Left, validateSize)
	validateRight, rightSize := getBinaryTreeSize(tree.Right, validateSize)

	if validateLeft {
		*validateSize = max(*validateSize, leftSize)
	}

	if validateRight {
		*validateSize = max(*validateSize, rightSize)
	}

	currentSize := leftSize + rightSize + 1

	var validate bool
	if tree.Left != nil && tree.Right != nil {
		if tree.Value > tree.Left.Value && tree.Value <= tree.Right.Value {
			validate = true
		}
	} else if tree.Left != nil {
		if tree.Value > tree.Left.Value {
			validate = true
		}
	} else if tree.Right != nil {
		if tree.Value <= tree.Right.Value {
			validate = true
		}
	} else {
		validate = true
	}

	validateCurrent := validate && validateLeft && validateRight

	if validateCurrent {
		*validateSize = max(*validateSize, currentSize)
	}

	return validateCurrent, currentSize
}
