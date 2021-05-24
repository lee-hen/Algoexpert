package max_path_sum

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	test := NewBinaryTree(1).insertAll([]int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11})
	require.Equal(t,18, MaxPathSum(test))
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{Value: value}
}

func (tree *BinaryTree) insertAll(values []int) *BinaryTree {
	for _, value := range values {
		tree.insert(value)
	}
	return tree
}

func (tree *BinaryTree) insert(value int) {
	queue := []*BinaryTree{tree}
	var current *BinaryTree
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Left)
		if current.Right == nil {
			current.Right = NewBinaryTree(value)
			break
		}
		queue = append(queue, current.Right)
	}
}


//          1
//     2                 3
//   4     5        6       7
//  8   -9
//    100
//   1
func TestCase2(t *testing.T) {
	root := NewBinaryTree(1)
	root.Left = &BinaryTree{Value: 2}
	root.Right = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right.Left = &BinaryTree{Value: 6}
	root.Right.Right = &BinaryTree{Value: 7}

	root.Left.Left.Left = &BinaryTree{Value: 8}
	root.Left.Right.Right = &BinaryTree{Value: -9}

	root.Left.Right.Right.Left = &BinaryTree{Value: 100}
	root.Left.Right.Right.Left.Left = &BinaryTree{Value: 1}

	require.Equal(t,111, MaxPathSum(root))
}

//    -2
// -1
func TestCase3(t *testing.T) {
	root := NewBinaryTree(-2)
	root.Left = &BinaryTree{Value: -1}
	root.Right = &BinaryTree{Value: -3}

	require.Equal(t,-1, MaxPathSum(root))
}
