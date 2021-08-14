package find_kth_largest_value_in_bst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	root := &BST{Value: 15}
	root.Left = &BST{Value: 5}
	root.Left.Left = &BST{Value: 2}
	root.Left.Left.Left = &BST{Value: 1}
	root.Left.Left.Right = &BST{Value: 3}
	root.Left.Right = &BST{Value: 5}
	root.Right = &BST{Value: 20}
	root.Right.Left = &BST{Value: 17}
	root.Right.Right = &BST{Value: 22}
	k := 3
	expected := 17
	actual := FindKthLargestValueInBst(root, k)
	require.Equal(t, expected, actual)
}
