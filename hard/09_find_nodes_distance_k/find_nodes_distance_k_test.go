package find_nodes_distance_k

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

//                  10
//             11
//          1
//    2           3
//  4   5             6
//                  7   8
func TestCase1(t *testing.T) {
	root1 := &BinaryTree{Value: 10}
	root1.Left = &BinaryTree{Value: 11}
	root := &BinaryTree{Value: 1}
	root1.Left.Left = root
	root.Left = &BinaryTree{Value: 2}
	root.Right = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right.Right = &BinaryTree{Value: 6}
	root.Right.Right.Left = &BinaryTree{Value: 7}
	root.Right.Right.Right = &BinaryTree{Value: 8}
	target := 3
	k := 2
	expected := []int{2, 7, 8, 11}
	actual := FindNodesDistanceK(root1, target, k)
	sort.Ints(actual)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Right = &BinaryTree{Value: 3}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Left.Right = &BinaryTree{Value: 5}
	root.Right.Right = &BinaryTree{Value: 6}
	root.Right.Right.Left = &BinaryTree{Value: 7}
	root.Right.Right.Right = &BinaryTree{Value: 8}
	target := 8
	k := 5
	expected := []int{4, 5}
	actual := FindNodesDistanceK(root, target, k)
	sort.Ints(actual)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	root := &BinaryTree{Value: 1}
	root.Left = &BinaryTree{Value: 2}
	root.Left.Left = &BinaryTree{Value: 4}
	root.Right = &BinaryTree{Value: 3}
	root.Right.Left = &BinaryTree{Value: 5}
	root.Right.Right = &BinaryTree{Value: 6}
	root.Right.Right.Right = &BinaryTree{Value: 7}
	root.Right.Right.Right.Right = &BinaryTree{Value: 8}

	target := 3
	k := 1
	expected := []int{1, 5, 6}
	actual := FindNodesDistanceK(root, target, k)
	sort.Ints(actual)
	require.Equal(t, expected, actual)
}
