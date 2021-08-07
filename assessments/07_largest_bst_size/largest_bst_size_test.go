package largest_bst_size

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{Value: value}
}

func TestCase1(t *testing.T) {
	root := NewBinaryTree(20)
	root.Left = NewBinaryTree(7)
	root.Left.Left = NewBinaryTree(0)
	root.Left.Right = NewBinaryTree(8)
	root.Left.Right.Left = NewBinaryTree(7)
	root.Left.Right.Right = NewBinaryTree(9)
	root.Right = NewBinaryTree(10)
	root.Right.Left = NewBinaryTree(4)
	root.Right.Left.Left = NewBinaryTree(2)
	root.Right.Left.Left.Left = NewBinaryTree(1)
	root.Right.Left.Right = NewBinaryTree(5)
	root.Right.Left.Right.Left = NewBinaryTree(50)
	root.Right.Left.Right.Left.Left = NewBinaryTree(60)
	root.Right.Right = NewBinaryTree(15)
	root.Right.Right.Left = NewBinaryTree(13)
	root.Right.Right.Left.Right = NewBinaryTree(14)
	root.Right.Right.Right = NewBinaryTree(22)

	actual := LargestBSTSize(root)
	require.Equal(t, 5, actual)
}
