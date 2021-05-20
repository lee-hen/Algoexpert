package validate_three_nodes

type BST struct {
	Value, Height int

	Left  *BST
	Right *BST
}

func ValidateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	searchOne := nodeOne
	searchTwo := nodeThree

	for {
		foundThreeFromOne := searchOne == nodeThree
		foundOneFromThree := searchTwo == nodeOne
		foundNodeTwo := searchOne == nodeTwo || searchTwo == nodeTwo
		finishedSearching := searchOne == nil && searchTwo == nil

		if foundThreeFromOne || foundOneFromThree || foundNodeTwo || finishedSearching {
			break
		}

		if searchOne != nil {
			if searchOne.Value > nodeTwo.Value {
				searchOne = searchOne.Left
			} else {
				searchOne = searchOne.Right
			}
		}

		if searchTwo != nil {
			if searchTwo.Value > nodeTwo.Value {
				searchTwo = searchTwo.Left
			} else {
				searchTwo = searchTwo.Right
			}
		}
	}

	foundNodeFromOther := searchOne == nodeThree || searchTwo == nodeOne
	foundNodeTwo := searchOne == nodeTwo || searchTwo == nodeTwo

	if !foundNodeTwo || foundNodeFromOther {
		return false
	}

	if searchOne == nodeTwo {
		return searchForTarget(nodeTwo, nodeThree)
	}

	return searchForTarget(nodeTwo, nodeOne)
}

func searchForTarget(node *BST, target *BST) bool {
	currentNode := node
	for currentNode != nil && currentNode != target {
		if target.Value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return currentNode == target
}

func ValidateThreeNodes2(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	if isDescendant2(nodeTwo, nodeOne) {
		return isDescendant2(nodeThree, nodeTwo)
	}

	if isDescendant2(nodeTwo, nodeThree) {
		return isDescendant2(nodeOne, nodeTwo)
	}
	return false
}

func isDescendant2(node *BST, target *BST) bool {
	currentNode := node
	for currentNode != nil && currentNode != target {
		if target.Value < currentNode.Value {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}
	}
	return currentNode == target
}

func ValidateThreeNodes1(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	if isDescendant1(nodeTwo, nodeOne) {
		return isDescendant1(nodeThree, nodeTwo)
	}
	if isDescendant1(nodeTwo, nodeThree) {
		return isDescendant1(nodeOne, nodeTwo)
	}
	return false
}

func isDescendant1(node *BST, target *BST) bool {
	if node == nil {
		return false
	}
	if node == target {
		return true
	}
	if target.Value < node.Value {
		return isDescendant1(node.Left, target)
	}
	return isDescendant1(node.Right, target)
}

// ValidateThreeNodes
// it isn't guaranteed that nodeOne or nodeThree will be ancestors or descendents of nodeTwo
// nodeOne or nodeThree is an ancestor of nodeTwo and the other node is a descendent of nodeTwo
// ex nodeOne is an ancestor of nodeTwo then it needs to see if nodeThree is a descendent of nodeTwo
//         5
//     2         7
//   1    4    6    8
// 0   3
// nodeOne  7
// nodeTwo  2
// nodeThree 3
// my solution
func validateThreeNodes(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	// nodeOne ancestor nodeThree is descendent
	if !validateThreeNodesHelper(nodeOne, nodeTwo, nodeThree) {
		// nodeThree ancestor nodeOne is descendent
		return validateThreeNodesHelper(nodeThree, nodeTwo, nodeOne)
	}

	return true
}

func validateThreeNodesHelper(nodeOne *BST, nodeTwo *BST, nodeThree *BST) bool {
	nodeHeight1 := Height(nodeOne)
	nodeHeight2 := Height(nodeTwo)
	nodeHeight3 := Height(nodeThree)

	if !(nodeHeight1 > nodeHeight2 && nodeHeight2 > nodeHeight3) {
		return false
	}

	// left
	if nodeTwo.Value < nodeOne.Value {
		if nodeThree.Value < nodeOne.Value && nodeThree.Value >= nodeTwo.Value {
			return true
		}
		if nodeThree.Value < nodeTwo.Value {
			return true
		}

	} else { // right
		if nodeThree.Value >= nodeOne.Value && nodeThree.Value < nodeTwo.Value {
			return true
		}
		if nodeThree.Value >= nodeTwo.Value {
			return true
		}
	}
	return false
}

func Height(tree *BST) int {
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
