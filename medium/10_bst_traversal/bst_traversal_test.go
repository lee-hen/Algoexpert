package bst_traversal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func TestCase1(t *testing.T) {
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Right = NewBST(22)

	inOrder := []int{1, 2, 5, 5, 10, 15, 22}
	preOrder := []int{10, 5, 2, 1, 5, 15, 22}
	postOrder := []int{1, 2, 5, 5, 22, 15, 10}

	require.Equal(t, inOrder, root.InOrderTraverse([]int{}))
	require.Equal(t, preOrder, root.PreOrderTraverse([]int{}))
	require.Equal(t, postOrder, root.PostOrderTraverse([]int{}))
}
