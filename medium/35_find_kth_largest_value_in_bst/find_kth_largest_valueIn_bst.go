package find_kth_largest_value_in_bst

type BST struct {
	Value int
	Left  *BST
	Right *BST
}

type treeInfo struct {
	numberOfNodesVisited   int
	latestVisitedNodeValue int
}

// FindKthLargestValueInBst
// O(h + k) time | O(h) space - where h is the height of the tree and k is the input parameter
func FindKthLargestValueInBst(tree *BST, k int) int {
	treeInfo := treeInfo{0, -1}
	reverseInOrderTraverse(tree, k, &treeInfo)
	return treeInfo.latestVisitedNodeValue
}

func reverseInOrderTraverse(node *BST, k int, treeInfo *treeInfo) {
	if node == nil || treeInfo.numberOfNodesVisited >= k {
		return
	}

	reverseInOrderTraverse(node.Right, k, treeInfo)
	if treeInfo.numberOfNodesVisited < k {
		treeInfo.numberOfNodesVisited += 1
		treeInfo.latestVisitedNodeValue = node.Value
		reverseInOrderTraverse(node.Left, k, treeInfo)
	}
}

//         15
//      5       20
//   2    5  17    22
// 1   3
func findKthLargestValueInBst(tree *BST, k int) int {
	sortedNodeValues := make([]int, 0)
	inOrderTraverse(tree, &sortedNodeValues)
	return sortedNodeValues[len(sortedNodeValues)-k]
}

func inOrderTraverse(node *BST, sortedNodeValues *[]int) {
	if node == nil {
		return
	}
	inOrderTraverse(node.Left, sortedNodeValues)
	*sortedNodeValues = append(*sortedNodeValues, node.Value)
	inOrderTraverse(node.Right, sortedNodeValues)
}
