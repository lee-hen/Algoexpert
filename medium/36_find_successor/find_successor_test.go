package find_successor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Left.Parent = root
	root.Right = &BinaryTree{Value: 3}
	root.Right.Parent = root
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Left.Parent = root.Left
	root.Left.Right = &BinaryTree{Value: 5}
	root.Left.Right.Parent = root.Left
	root.Left.Left.Left = &BinaryTree{Value: 6}
	root.Left.Left.Left.Parent = root.Left.Left
	node := root.Left.Right
	expected := root
	actual := FindSuccessor(root, node)
	require.Equal(t, expected, actual)
}
