package binary_tree_diameter

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 7}
	root.Left.Left.Left = &BinaryTree{Value: 8}
	root.Left.Left.Left.Left = &BinaryTree{Value: 9}
	root.Left.Right = &BinaryTree{Value: 4}
	root.Left.Right.Right = &BinaryTree{Value: 5}
	root.Left.Right.Right.Right = &BinaryTree{Value: 6}
	root.Right = &BinaryTree{Value: 2}
	expected := 6
	actual := BinaryTreeDiameter(root)
	require.Equal(t, expected, actual)
}
