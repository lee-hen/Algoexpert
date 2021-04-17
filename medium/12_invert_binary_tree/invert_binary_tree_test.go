package invert_binary_tree


import (
	"github.com/stretchr/testify/require"
	"testing"
)

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

func (tree *BinaryTree) Insert(value int) *BinaryTree {
	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Left == nil {
			current.Left = NewBinaryTree(value)
			break
		} else if current.Right == nil {
			current.Right = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Left, current.Right)
	}
	return tree
}

func (tree *BinaryTree) InvertedInsert(value int) *BinaryTree {
	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.Right == nil {
			current.Right = NewBinaryTree(value)
			break
		} else if current.Left == nil {
			current.Left = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Right, current.Left)
	}
	return tree
}

func (tree *BinaryTree) Equals(other *BinaryTree) bool {
	if other == nil || tree.Value != other.Value {
		return false
	}
	if tree.Left != nil && !tree.Left.Equals(other.Left) {
		return false
	}
	if tree.Right != nil && !tree.Right.Equals(other.Right) {
		return false
	}
	return true
}

func (tree *BinaryTree) InsertAll(values ...int) *BinaryTree {
	for _, value := range values {
		tree.Insert(value)
	}
	return tree
}

func (tree *BinaryTree) InvertedInsertAll(values ...int) *BinaryTree {
	for _, value := range values {
		tree.InvertedInsert(value)
	}
	return tree
}

func TestCase1(t *testing.T) {
	output := NewBinaryTree(1).InsertAll(2, 3, 4, 5, 6, 7, 8, 9)
	output.InvertBinaryTree()
	expected := NewBinaryTree(1).InvertedInsertAll(2, 3, 4, 5, 6, 7, 8, 9)
	require.True(t, output.Equals(expected))
}

