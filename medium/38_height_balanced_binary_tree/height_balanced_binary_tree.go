package height_balanced_binary_tree

type BinaryTree struct {
	Value, Height int

	Left  *BinaryTree
	Right *BinaryTree
}

type treeInfo struct {
	isBalanced bool
	height     int
}

// O(n) time | O(h) space - where n is the number of nodes in the binary tree

func HeightBalancedBinaryTree(tree *BinaryTree) bool {
	treeInfo := getTreeInfo(tree)
	return treeInfo.isBalanced
}

func getTreeInfo(node *BinaryTree) treeInfo {
	if node == nil {
		return treeInfo{isBalanced: true, height: -1}
	}

	leftSubtreeInfo := getTreeInfo(node.Left)

	rightSubtreeInfo := getTreeInfo(node.Right)

	isBalanced := leftSubtreeInfo.isBalanced &&
		rightSubtreeInfo.isBalanced &&
		abs(leftSubtreeInfo.height-rightSubtreeInfo.height) <= 1

	height := max(leftSubtreeInfo.height, rightSubtreeInfo.height) + 1

	return treeInfo{
		isBalanced: isBalanced,
		height:     height,
	}
}


//      1
//    2     3
//  4   5      6
//    7   8      9
//

func heightBalancedBinaryTree(tree *BinaryTree) (balanced bool) {
	if tree == nil {
		return true
	}

	balanced = abs(Height(tree.Right) - Height(tree.Left)) <= 1
	if balanced {
		return heightBalancedBinaryTree(tree.Left) && heightBalancedBinaryTree(tree.Right)
	}

	return
}

func Height(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}

	if tree.Height != 0 {
		return tree.Height
	}

	tree.Height = max(Height(tree.Left), Height(tree.Right)) + 1
	return tree.Height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
