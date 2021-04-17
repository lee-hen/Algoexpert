package min_height_bst

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	array := []int{1, 2, 5, 7, 10, 13, 14, 15, 22}
	tree := MinHeightBST(array)

	require.True(t, validateBST(tree, math.MinInt32, math.MaxInt32))
	require.Equal(t, 4, getTreeHeight(tree))

	inOrder := inOrderTraverse(tree, []int{})
	require.Equal(t, []int{1, 2, 5, 7, 10, 13, 14, 15, 22}, inOrder)
}

func validateBST(tree *BST, min, max int) bool {
	if tree.Value < min || tree.Value >= max {
		return false
	}
	if tree.Left != nil && !validateBST(tree.Left, min, tree.Value) {
		return false
	}
	if tree.Right != nil && !validateBST(tree.Right, tree.Value, max) {
		return false
	}
	return true
}

func inOrderTraverse(tree *BST, array []int) []int {
	if tree.Left != nil {
		array = inOrderTraverse(tree.Left, array)
	}
	array = append(array, tree.Value)
	if tree.Right != nil {
		array = inOrderTraverse(tree.Right, array)
	}
	return array
}

func getTreeHeight(tree *BST) int {
	if tree == nil {
		return 0
	}
	left := getTreeHeight(tree.Left)
	right := getTreeHeight(tree.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

