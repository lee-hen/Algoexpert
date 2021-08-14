package binary_tree_diameter
// important question

type BinaryTree struct {
	Value, Height int

	Left  *BinaryTree
	Right *BinaryTree
}

type TreeInfo struct {
	diameter int
	height   int
}

// BinaryTreeDiameter
// well my solution is prefer
func BinaryTreeDiameter(tree *BinaryTree) int {
	return getTreeInfo(tree).diameter
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{0, 0}
	}

	leftTreeInfo := getTreeInfo(tree.Left)
	rightTreeInfo := getTreeInfo(tree.Right)

	longestPathThroughRoot := leftTreeInfo.height + rightTreeInfo.height
	maxDiameterSoFar := max(leftTreeInfo.diameter, rightTreeInfo.diameter)
	currentDiameter := max(longestPathThroughRoot, maxDiameterSoFar)
	currentHeight := 1 + max(leftTreeInfo.height, rightTreeInfo.height)

	return TreeInfo{currentDiameter, currentHeight}
}

// BinaryTreeDiameter
// my solution
// Easy to understand.
func binaryTreeDiameter(tree *BinaryTree) int {
	if tree == nil {
		return 0
	}

	currentDiameter := Height(tree.Left) + Height(tree.Right)
	return max(max(currentDiameter, binaryTreeDiameter(tree.Left)), binaryTreeDiameter(tree.Right))
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
