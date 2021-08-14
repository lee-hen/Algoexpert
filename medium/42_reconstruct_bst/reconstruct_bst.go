package reconstruct_bst
// important question

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

type treeInfo struct {
	rootIdx int
}

func ReconstructBst(preOrderTraversalValues []int) *BST {
	treeInfo := &treeInfo{rootIdx: 0}
	return reconstructBstFromRange(
		math.MinInt32,
		math.MaxInt32,
		preOrderTraversalValues,
		treeInfo,
	)
}

func reconstructBstFromRange(lowerBound, upperBound int, preOrderTraversalValues []int, currentSubtreeInfo *treeInfo) *BST {
	if currentSubtreeInfo.rootIdx == len(preOrderTraversalValues) {
		return nil
	}


	rootValue := preOrderTraversalValues[currentSubtreeInfo.rootIdx]
	if rootValue < lowerBound || rootValue >= upperBound {
		return nil
	}

	currentSubtreeInfo.rootIdx += 1

	leftSubtree := reconstructBstFromRange(
		lowerBound,
		rootValue,
		preOrderTraversalValues,
		currentSubtreeInfo,
	)
	rightSubtree := reconstructBstFromRange(
		rootValue,
		upperBound,
		preOrderTraversalValues,
		currentSubtreeInfo,
	)
	return &BST{Value: rootValue, Left: leftSubtree, Right: rightSubtree}
}

// 10, 4, 2, 1, 5, 17, 19, 18

//        10
//      /    \
//     4      17
//   /   \      \
//  2     5     19
// /           /
// 1          18
// visit left right

// ReconstructBst1
// Why I didn't figure it out
func ReconstructBst1(preOrderTraversalValues []int) *BST {
	if len(preOrderTraversalValues) == 0 {
		return nil
	}

	currentValue := preOrderTraversalValues[0]
	rightSubtreeRootIdx := len(preOrderTraversalValues)

	for idx := 1; idx < len(preOrderTraversalValues); idx++ {
		value := preOrderTraversalValues[idx]
		if value >= currentValue {
			rightSubtreeRootIdx = idx
			break
		}
	}

	leftSubtree := ReconstructBst(preOrderTraversalValues[1:rightSubtreeRootIdx])
    rightSubtree := ReconstructBst(preOrderTraversalValues[rightSubtreeRootIdx:])
	return &BST{
		Value: currentValue,
		Left: leftSubtree,
		Right: rightSubtree,
	}
}

// my solution
func reconstructBst(preOrderTraversalValues []int) *BST {
	bst := &BST{Value: preOrderTraversalValues[0]}
	for i := 1; i < len(preOrderTraversalValues); i++ {
		value := preOrderTraversalValues[i]
		bst.insert(value)
	}

	return bst
}

func (tree *BST) insert(value int)  *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST {
				Value: value,
			}
		} else {
			tree.Left.insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.insert(value)
		}
	}
	return tree
}
