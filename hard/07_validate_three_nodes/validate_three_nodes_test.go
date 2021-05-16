package validate_three_nodes

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	root := &BST{Value: 5}
	root.Left = &BST{Value: 2}
	root.Right = &BST{Value: 7}
	root.Left.Left = &BST{Value: 1}
	root.Left.Right = &BST{Value: 4}
	root.Right.Left = &BST{Value: 6}
	root.Right.Right = &BST{Value: 8}
	root.Left.Left.Left = &BST{Value: 0}
	root.Left.Right.Left = &BST{Value: 3}

	nodeOne := root
	nodeTwo := root.Left
	nodeThree := root.Left.Right.Left
	expected := true
	actual := ValidateThreeNodes(nodeOne, nodeTwo, nodeThree)
	require.Equal(t, expected, actual)
}
