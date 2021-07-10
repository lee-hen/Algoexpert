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
