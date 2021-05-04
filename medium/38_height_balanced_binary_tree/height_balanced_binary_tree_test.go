package height_balanced_binary_tree

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Right = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right.Right = &BinaryTree{Value: 6}
	root.Left.Right.Left = &BinaryTree{Value: 7}
	root.Left.Right.Right = &BinaryTree{Value: 8}
	expected := true
	actual := HeightBalancedBinaryTree(root)
	require.Equal(t, expected, actual)
}
