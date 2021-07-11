package right_sibling_tree

import "fmt"

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

//
//               1
//          /           \
//         2             3
//       /   \         /   \
//      4     5       6     7
//    /   \    \    /     /   \
//   8     9   10  11    12   13
//                /
//               14

//         1
//      /
//     2-----------3
//    /           /
//   4-----5-----6-----7
//  /           /     /
// 8---9    10-11    12-13
//            /
//          14

func RightSiblingTree(root *BinaryTree) *BinaryTree {
	arr := make([]int, 0)
	fmt.Println(root.PreOrderTraverse(arr))
	return nil
}
//
//func (tree *BinaryTree) PreOrderTraverse() {
//	if tree == nil {
//		return
//	}
//
//	fmt.Println(tree.Value)
//	tree.Left.PreOrderTraverse()
//	tree.Right.PreOrderTraverse()
//}

func (tree *BinaryTree) PreOrderTraverse(arr []int) []int {
	current := tree
	if current == nil {
		return arr
	}

	arr = append(arr, current.Value)
	arr = current.Left.PreOrderTraverse(arr)
	arr = current.Right.PreOrderTraverse(arr)
	return arr
}