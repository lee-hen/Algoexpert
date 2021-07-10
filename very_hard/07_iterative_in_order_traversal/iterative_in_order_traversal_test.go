package iterative_in_order_traversal

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewBinaryTree(root int, parent *BinaryTree) *BinaryTree {
	return &BinaryTree{Value: root, Parent: parent}
}

func TestCase1(t *testing.T) {
	root := NewBinaryTree(1, nil)
	root.Left = NewBinaryTree(2, root)
	root.Left.Left = NewBinaryTree(4, root.Left)

	//root.Left.Right = NewBinaryTree(5, root.Left)
	//root.Left.Right.Left = NewBinaryTree(10, root.Left.Right)
	//root.Left.Right.Right = NewBinaryTree(11, root.Left.Right)

	root.Left.Left.Right = NewBinaryTree(9, root.Left.Left)
	root.Right = NewBinaryTree(3, root)
	root.Right.Left = NewBinaryTree(6, root.Right)
	root.Right.Right = NewBinaryTree(7, root.Right)

	output := make([]int, 0)
	root.IterativeInOrderTraversal(func(i int) {
		output = append(output, i)
	})
	expected := []int{4, 9, 2, 1, 6, 3, 7}
	require.Equal(t, expected, output)
}
