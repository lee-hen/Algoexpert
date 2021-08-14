package all_kinds_of_node_depths
// important question

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

// AllKindsOfNodeDepths
// Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree
func AllKindsOfNodeDepths(root *BinaryTree) int {
	return allKindsOfNodeDepthsHelper(root, 0)
}

func allKindsOfNodeDepthsHelper(root *BinaryTree, depth int) int {
	if root == nil {
		return 0
	}

	// Formula to calculate 1 + 2 + 3 + ... + depth - 1 + depth
	depthSum := (depth * (depth + 1)) / 2
	return depthSum + allKindsOfNodeDepthsHelper(root.Left, depth+1) + allKindsOfNodeDepthsHelper(root.Right, depth+1)
}

// AllKindsOfNodeDepths2
// Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree
func AllKindsOfNodeDepths2(root *BinaryTree) int {
	return dfsHelper(root, 0, 0)
}

func dfsHelper(root *BinaryTree, depthSum, depth int) int {
	if root == nil {
		return 0
	}

	depthSum += depth
	return depthSum + dfsHelper(root.Left, depthSum, depth+1) + dfsHelper(root.Right, depthSum, depth+1)
}

// AllKindsOfNodeDepths1
// Average case: when the tree is balanced
// O(nlog(n)) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree
func AllKindsOfNodeDepths1(root *BinaryTree) int {
	if root == nil {
		return 0
	}
	return AllKindsOfNodeDepths(root.Left) + AllKindsOfNodeDepths(root.Right) + nodeDepths(root, 0)
}

func nodeDepths(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}
	return depth + nodeDepths(node.Left, depth+1) + nodeDepths(node.Right, depth+1)
}

type TreeInfo struct {
	NumNodesInTree int
	SumOfDepths    int
	SumOfAllDepths int
}

// allKindsOfNodeDepths
// Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of nodes in
// the Binary Tree and h is the height of the Binary Tree
func allKindsOfNodeDepths(root *BinaryTree) int {
	return getTreeInfo(root).SumOfAllDepths
}

func getTreeInfo(tree *BinaryTree) TreeInfo {
	if tree == nil {
		return TreeInfo{}
	}

	leftInfo, rightInfo := getTreeInfo(tree.Left), getTreeInfo(tree.Right)

	sumOfLeftDepths := leftInfo.SumOfDepths + leftInfo.NumNodesInTree
	sumOfRightDepths := rightInfo.SumOfDepths + rightInfo.NumNodesInTree

	numNodesInTree := 1 + leftInfo.NumNodesInTree + rightInfo.NumNodesInTree
	sumOfDepths := sumOfLeftDepths + sumOfRightDepths
	sumOfAllDepths := sumOfDepths + leftInfo.SumOfAllDepths + rightInfo.SumOfAllDepths

	return TreeInfo{NumNodesInTree: numNodesInTree, SumOfDepths: sumOfDepths, SumOfAllDepths: sumOfAllDepths}
}

// 26

//           1         level    0)        node 1:   16
//        /     \
//       2       3     level    1)        node 2:   6     node 3:　 2
//     /   \   /   \
//    4     5 6     7  level    2)        node 4:   2     node 5:   0   node 6:  0   node 7:　 0
//  /   \
// 8     9             level    3)        node 8:   0     node 9:　　0

func allKindsOfNodeDepths2(root *BinaryTree) int {
	depth := root.Depth()

	return root.dfs2(0, depth, depth)
}

func (tree *BinaryTree) dfs2(depthSum, depth, level int) int {
	if tree == nil {
		return 0
	}

	depthSum += depth - level
	depthLeftSum := tree.Left.dfs2(depthSum, depth, level-1)
	depthRightSum := tree.Right.dfs2(depthSum, depth, level-1)
	return depthSum + depthLeftSum + depthRightSum
}

// allKindsOfNodeDepths1
// naive approach
func allKindsOfNodeDepths1(root *BinaryTree) int {
	depth := root.Depth()
	return traverseBinaryTree(root, depth)
}

func traverseBinaryTree(node *BinaryTree, depth int) int {
	if node == nil {
		return 0
	}

	//sumDepth := node.iddfs(depth)
	sumDepth := node.dfs(depth, depth, 0)
	sumDepth += traverseBinaryTree(node.Left, depth-1)
	sumDepth += traverseBinaryTree(node.Right, depth-1)

	return sumDepth
}

func (tree *BinaryTree) dfs(level, depth, sum int) int {
	if tree == nil {
		return sum
	}

	sum += depth - level
	sum = tree.Left.dfs(level-1, depth, sum)
	sum = tree.Right.dfs(level-1, depth, sum)
	return sum
}

func (tree *BinaryTree) Depth() int {
	if tree == nil {
		return 0
	}

	return max(tree.Left.Depth(), tree.Right.Depth()) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// iddfs approach O(b^d)
//func (tree *BinaryTree) iddfs(depth int) int {
//	var sumDepth int
//	for i := 0; i < depth; i++ {
//		sum := tree.dfs(i)
//		sumDepth += sum * i
//	}
//	return sumDepth
//}
//
//func (tree *BinaryTree) dfs(level int) int {
//	if tree == nil {
//		return 0
//	}
//
//	if level == 0 {
//		return 1
//	}
//
//	leftSum := tree.Left.dfs(level-1)
//	rightSum := tree.Right.dfs(level-1)
//
//	return leftSum + rightSum
//}
