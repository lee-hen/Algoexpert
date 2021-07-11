package flatten_binary_tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{Value: value}
}

func (tree *BinaryTree) insert(value int) *BinaryTree {
	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		queue, current = queue[1:], queue[0]
		if current.Left == nil {
			current.Left = NewBinaryTree(value)
			return tree
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = NewBinaryTree(value)
			return tree
		}
		queue = append(queue, current.Right)
	}
	return tree
}

func (tree *BinaryTree) insertAll(values ...int) *BinaryTree {
	for _, value := range values {
		tree.insert(value)
	}
	return tree
}

func (tree *BinaryTree) leftToRightToLeft() []int {
	if tree == nil {
		return nil
	}

	subResult := append(tree.Right.leftToRightToLeft(), tree.Value)
	return append([]int{tree.Value}, subResult...)
}

func TestCase1(t *testing.T) {
	root := NewBinaryTree(1).insertAll(2, 3, 4, 5, 6)
	root.Left.Right.Left = NewBinaryTree(7)
	root.Left.Right.Right = NewBinaryTree(8)
	leftMostNode := FlattenBinaryTree(root)
	actual := leftMostNode.leftToRightToLeft()
	expected := []int{4, 2, 7, 5, 8, 1, 6, 3, 3, 6, 1, 8, 5, 7, 2, 4}
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	root := NewBinaryTree(1).insertAll(2, 3)
	root.Right.Left = NewBinaryTree(6)
	root.Right.Right = NewBinaryTree(7)

	root.Right.Left.Left = NewBinaryTree(12)

	root.Left.Left = NewBinaryTree(4)
	root.Left.Right = NewBinaryTree(5)

	root.Left.Right.Left =  NewBinaryTree(10)
	root.Left.Right.Right = NewBinaryTree(11)

	root.Left.Left.Left = NewBinaryTree(8)
	root.Left.Left.Left.Left = NewBinaryTree(13)
	root.Left.Left.Left.Right = NewBinaryTree(14)

	root.Left.Left.Right = NewBinaryTree(9)

	leftMostNode := FlattenBinaryTree(root)
	actual := leftMostNode.leftToRightToLeft()
	expected := []int{13, 8, 14, 4, 9, 2, 10, 5, 11, 1, 12, 6, 3, 7, 7, 3, 6, 12, 1, 11, 5, 10, 2, 9, 4, 14, 8, 13}
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	root := NewBinaryTree(1)
	leftMostNode := FlattenBinaryTree(root)
	actual := leftMostNode.leftToRightToLeft()
	expected := []int{1, 1}
	require.Equal(t, expected, actual)
}
