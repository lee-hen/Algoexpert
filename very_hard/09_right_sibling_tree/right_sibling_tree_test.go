package right_sibling_tree

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

func (tree *BinaryTree) getDfsOrder() []int {
	if tree == nil {
		return []int{}
	}
	values := []int{tree.Value}
	if tree.Left != nil {
		values = append(values, tree.Left.getDfsOrder()...)
	}
	if tree.Right != nil {
		values = append(values, tree.Right.getDfsOrder()...)
	}
	return values
}

func TestCase1(t *testing.T) {
	root := NewBinaryTree(1).insertAll(2, 3, 4, 5, 6, 7, 8, 9)
	root.Left.Right.Right = NewBinaryTree(10)
	root.Right.Left.Left = NewBinaryTree(11)
	root.Right.Right.Left = NewBinaryTree(12)
	root.Right.Right.Right = NewBinaryTree(13)
	root.Right.Left.Left.Left = NewBinaryTree(14)
	mutatedRoot := RightSiblingTree(root)
	actual := mutatedRoot.getDfsOrder()
	expected := []int{1, 2, 4, 8, 9, 5, 6, 11, 14, 7, 12, 13, 3, 6, 11, 14, 7, 12, 13}
	require.Equal(t, expected, actual)
}
