package compare_leaf_traversal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	tree1 := &BinaryTree{Value: 1}
	tree1.Left = &BinaryTree{Value: 2}
	tree1.Right = &BinaryTree{Value: 3}
	tree1.Left.Left = &BinaryTree{Value: 4}
	tree1.Left.Right = &BinaryTree{Value: 5}
	tree1.Right.Right = &BinaryTree{Value: 6}
	tree1.Left.Right.Left = &BinaryTree{Value: 7}
	tree1.Left.Right.Right = &BinaryTree{Value: 8}

	tree2 := &BinaryTree{Value: 1}
	tree2.Left = &BinaryTree{Value: 2}
	tree2.Right = &BinaryTree{Value: 3}
	tree2.Left.Left = &BinaryTree{Value: 4}
	tree2.Left.Right = &BinaryTree{Value: 7}
	tree2.Right.Right = &BinaryTree{Value: 5}
	tree2.Right.Right.Left = &BinaryTree{Value: 8}
	tree2.Right.Right.Right = &BinaryTree{Value: 6}

	expected := true
	actual := CompareLeafTraversal(tree1, tree2)
	require.Equal(t, expected, actual)
}
