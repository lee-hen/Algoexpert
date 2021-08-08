package subtrees_within_range
// important question

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// SubtreesWithinRange
// Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of
// nodes in the BST and h is the height of the BST
func SubtreesWithinRange(tree *BST, targetRange []int) int {
	answer := 0
	isTreeWithinRange(tree, targetRange, &answer)
	return answer
}

func isTreeWithinRange(tree *BST, targetRange []int, answer *int) bool {
	if tree == nil {
		return true
	}

	leftTreeWithinRange := isTreeWithinRange(tree.Left, targetRange, answer)
	rightTreeWithinRange := isTreeWithinRange(tree.Right, targetRange, answer)
	nodeInRange := tree.Value >= targetRange[0] && tree.Value <= targetRange[1]
	treeWithinRange := leftTreeWithinRange && rightTreeWithinRange && nodeInRange

	if treeWithinRange {
		*answer++
	}

	return treeWithinRange
}

type treeInfo struct {
	maxValue               int
	minValue               int
	numSubtreesWithinRange int
}

// SubtreesWithinRange
// Average case: when the tree is balanced
// O(n) time | O(h) space - where n is the number of
// nodes in the BST and h is the height of the BST
func subtreesWithinRange1(tree *BST, targetRange []int) int {
	return getTreeInfo(tree, targetRange).numSubtreesWithinRange
}

func getTreeInfo(tree *BST, targetRange []int) treeInfo {
	numSubtreesWithinRange := 0
	maxValue := tree.Value
	minValue := tree.Value

	if tree.Left != nil {
		leftSubtreeInfo := getTreeInfo(tree.Left, targetRange)
		minValue = leftSubtreeInfo.minValue
		numSubtreesWithinRange += leftSubtreeInfo.numSubtreesWithinRange
	}

	if tree.Right != nil {
		rightSubtreeInfo := getTreeInfo(tree.Right, targetRange)
		maxValue = rightSubtreeInfo.maxValue
		numSubtreesWithinRange += rightSubtreeInfo.numSubtreesWithinRange
	}

	if minValue >= targetRange[0] && maxValue <= targetRange[1] {
		numSubtreesWithinRange++
	}

	return treeInfo{
		maxValue:               maxValue,
		minValue:               minValue,
		numSubtreesWithinRange: numSubtreesWithinRange,
	}
}

// tree =   10
//        /     \
//       5      15
//     /   \   /   \
//    2     8 13    22
//  /      / \  \     \
// 1      5   9  14     23
//   x  y
//targetRange = [5, 15]

// 5
// The 5 subtrees are:
//   8    5    9    13    14
//  / \               \
// 5   9               14

// my solution

func subtreesWithinRange(tree *BST, targetRange []int) int {
	var inRangeSubTrees int
	tree.isBstInRange(targetRange, false, &inRangeSubTrees)
	return inRangeSubTrees
}

func (tree *BST) isBstInRange(targetRange []int, inRange bool, inRangeSubTrees *int) bool {
	if tree == nil {
		return true
	}

	leftInRange := tree.Left.isBstInRange(targetRange, inRange, inRangeSubTrees)
	rightInRange := tree.Right.isBstInRange(targetRange, inRange, inRangeSubTrees)

	if tree.Value >= targetRange[0] && tree.Value <= targetRange[1] {
		inRange = true
	} else {
		inRange = false
	}

	if inRange && leftInRange && rightInRange {
		*inRangeSubTrees++
	}

	return inRange && leftInRange && rightInRange
}
