package subtrees_within_range

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func newBST(val int) *BST {
	return &BST{Value: val}
}

func TestCase1(t *testing.T) {
	root := newBST(10)
	root.Left = newBST(5)
	root.Left.Left = newBST(2)
	root.Left.Left.Left = newBST(1)
	root.Left.Right = newBST(8)
	root.Left.Right.Left = newBST(5)
	root.Left.Right.Right = newBST(9)
	root.Right = newBST(15)
	root.Right.Left = newBST(13)
	root.Right.Left.Right = newBST(14)
	root.Right.Right = newBST(22)
	// root.Right.Right.Left = newBST(15)
	targetRange := []int{5, 15}

	output := SubtreesWithinRange(root, targetRange)
	require.Equal(t, 5, output)
}

func TestCase2(t *testing.T) {
	root := newBST(10)
	root.Left = newBST(5)
	root.Left.Right = newBST(8)
	root.Left.Right.Left = newBST(5)
	root.Left.Right.Right = newBST(9)
	root.Left.Left = newBST(2)
	root.Left.Left.Left = newBST(1)
	root.Right = newBST(15)
	root.Right.Left = newBST(13)
	root.Right.Left.Right = newBST(14)
	root.Right.Right = newBST(22)
	targetRange := []int{1, 21}

	output := SubtreesWithinRange(root, targetRange)
	require.Equal(t, 8, output)
}

func TestCase3(t *testing.T) {
	root := newBST(10)
	root.Left = newBST(5)
	root.Left.Left = newBST(2)
	root.Left.Left.Left = newBST(1)
	root.Left.Right = newBST(8)
	root.Left.Right.Left = newBST(5)
	root.Left.Right.Right = newBST(9)
	root.Right = newBST(15)
	root.Right.Left = newBST(13)
	root.Right.Left.Right = newBST(14)
	root.Right.Right = newBST(22)
	targetRange := []int{10, 12}

	output := SubtreesWithinRange(root, targetRange)
	require.Equal(t, 0, output)
}


func TestCase4(t *testing.T) {
	root := newBST(10)
	root.Left = newBST(5)
	root.Left.Left = newBST(2)
	root.Left.Left.Left = newBST(1)
	root.Left.Right = newBST(8)
	root.Left.Right.Left = newBST(5)
	root.Left.Right.Right = newBST(9)
	root.Right = newBST(15)
	root.Right.Left = newBST(13)
	root.Right.Left.Right = newBST(14)
	root.Right.Right = newBST(22)
	targetRange := []int{2, 8}

	output := SubtreesWithinRange(root, targetRange)
	require.Equal(t, 1, output)
}


func TestCase5(t *testing.T) {
	root := newBST(10)
	root.Left = newBST(5)
	root.Left.Left = newBST(2)
	root.Left.Left.Left = newBST(1)
	root.Left.Right = newBST(8)
	root.Left.Right.Left = newBST(5)
	root.Left.Right.Right = newBST(9)
	root.Right = newBST(15)
	root.Right.Left = newBST(13)
	root.Right.Left.Right = newBST(14)
	root.Right.Right = newBST(22)
	targetRange := []int{2, 22}

	output := SubtreesWithinRange(root, targetRange)
	require.Equal(t, 7, output)
}

func TestCase6(t *testing.T) {
	root := newBST(10)
	root.Left = newBST(5)
	root.Left.Left = newBST(2)
	root.Left.Left.Left = newBST(1)
	root.Left.Right = newBST(8)
	root.Left.Right.Left = newBST(5)
	root.Left.Right.Right = newBST(9)
	root.Right = newBST(15)
	root.Right.Left = newBST(13)
	root.Right.Left.Right = newBST(14)
	root.Right.Right = newBST(22)
	targetRange := []int{1, 13}

	output := SubtreesWithinRange(root, targetRange)
	require.Equal(t, 6, output)
}
