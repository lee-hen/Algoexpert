package reconstruct_bst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func getDfsOrder(node *BST, values []int) []int {
	if node == nil {
		return nil
	}
	values = append(values, node.Value)
	getDfsOrder(node.Left, values)
	getDfsOrder(node.Right, values)
	return values
}

func TestCase1(t *testing.T) {
	preOrderTraversalValues := []int{10, 4, 2, 1, 5, 17, 19, 18}
	tree := &BST{Value: 10}
	tree.Left = &BST{Value: 4}
	tree.Left.Left = &BST{Value: 2}
	tree.Left.Left.Left = &BST{Value: 1}
	tree.Left.Right = &BST{Value: 5}
	tree.Right = &BST{Value: 17}
	tree.Right.Right = &BST{Value: 19}
	tree.Right.Right.Left = &BST{Value: 18}
	expected := getDfsOrder(tree, nil)
	actual := ReconstructBst(preOrderTraversalValues)
	actualOrder := getDfsOrder(actual, nil)
	require.Equal(t, expected, actualOrder)
}
