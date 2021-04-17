package validate_bst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func  TestCase1(t *testing.T) {
	root := NewBST(7)

	root.Right = NewBST(10)
	root.Right.Left = NewBST(7)
	root.Right.Left.Left = NewBST(6)
	root.Right.Right = NewBST(12)

	output := root.ValidateBst()
	require.False(t, output)
}
