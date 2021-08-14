package flatten_binary_tree
// important question

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// FlattenBinaryTree
// O(n) time | O(d) space - where n is the number of nodes in the Binary Tree
// and d is the depth (height) of the Binary Tree
func FlattenBinaryTree(root *BinaryTree) *BinaryTree {
	leftMost, _ := flattenTree(root)
	return leftMost
}

func flattenTree(node *BinaryTree) (leftMost, rightMost *BinaryTree) {
	leftMost = node
	if node.Left != nil {
		leftSubtreeLeftMost, leftSubtreeRightMost := flattenTree(node.Left)
		connectNodes(leftSubtreeRightMost, node)
		leftMost = leftSubtreeLeftMost
	}

	rightMost = node
	if node.Right != nil {
		rightSubtreeLeftMost, rightSubtreeRightMost := flattenTree(node.Right)
		connectNodes(node, rightSubtreeLeftMost)
		rightMost = rightSubtreeRightMost
	}
	return leftMost, rightMost
}

func connectNodes(left, right *BinaryTree) {
	left.Right = right
	right.Left = left
}

//           1
//        /     \
//       2       3
//     /   \    /
//    4     5  6
//        /   \
//       7     8
// 4 <-> 2 <-> 7 <-> 5 <-> 8 <-> 1 <-> 6 <-> 3

//                1
//           /           \
//          2               3
//        /    \           /  \
//       4       5        6    7
//     /   \    /  \     /
//    8     9  10   11  12
//  /   \
// 13    14
// 13  8, 14,  4, 9, 2, 10, 5, 11, 1, 12, 6, 3, 7
// 11 1 12 consider root node left right most right left most

// my solution
// flattenBinaryTree
func flattenBinaryTree(root *BinaryTree) *BinaryTree {
	root.dfs()
	left := root

	for left.Left != nil {
		left = left.Left
	}

	return left
}

func (tree *BinaryTree) dfs() {
	if tree == nil {
		return
	}

	tree.Left.dfs()
	if tree.Left != nil {
		rightMost := getRightmostChild(tree.Left)
		rightMost.Right = tree
		tree.Left = rightMost
	}

	tree.Right.dfs()
	if tree.Right != nil {
		leftMost := getLeftmostChild(tree.Right)
		tree.Right = leftMost
		leftMost.Left = tree
	}
}

func getRightmostChild(node *BinaryTree) *BinaryTree {
	var currentNode = node
	for currentNode.Right != nil  {
		currentNode = currentNode.Right
	}
	return currentNode
}

func getLeftmostChild(node *BinaryTree) *BinaryTree {
	var currentNode = node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}
