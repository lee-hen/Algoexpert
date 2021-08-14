package find_successor
// important question

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

func FindSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	if node.Right != nil {
		return getLeftmostChild(node.Right)
	}
	return getRightmostParent(node)
}

func getLeftmostChild(node *BinaryTree) *BinaryTree {
	var currentNode = node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}

func getRightmostParent(node *BinaryTree) *BinaryTree {
	var currentNode = node
	for currentNode.Parent != nil && currentNode.Parent.Right == currentNode {
		currentNode = currentNode.Parent
	}
	return currentNode.Parent
}

func FindSuccessor1(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	visited := map[*BinaryTree]struct{}{}
	// https://github.com/lee-hen/Algoexpert/blob/master/medium/10_bst_traversal/bst_traversal.go
	return tree.inOrderTraverse(node, visited)
}

//           0
//       1
//    2     3
//  	  4
//      5
//    6
//  7
// 2 1 7 6 5 4 3
func (tree *BinaryTree) inOrderTraverse(node *BinaryTree, visited map[*BinaryTree]struct{}) (successor *BinaryTree) {
	if tree == nil {
		return successor
	}

	successor = tree.Left.inOrderTraverse(node, visited)
	if successor != nil {
		return
	}
	visited[tree] = struct{}{}
	if tree == node {
		if tree.Right != nil {
			successor = tree.Right
			for successor.Left != nil {
				successor = successor.Left
			}
		} else {
			successor = tree.Parent
			_, found := visited[successor]
			for successor != nil && found {
				successor = successor.Parent
				_, found = visited[successor]
			}
		}
		return
	}

	successor = tree.Right.inOrderTraverse(node, visited)
	if successor != nil {
		return
	}

	return
}

func findSuccessor(tree *BinaryTree, node *BinaryTree) *BinaryTree {
	inOrderTraversalOrder := make([]*BinaryTree, 0)
	getInOrderTraversalOrder(tree, &inOrderTraversalOrder)
	for idx, currentNode := range inOrderTraversalOrder {
		if currentNode != node {
			continue
		}
		if idx == len(inOrderTraversalOrder)-1 {
			return nil
		}
		return inOrderTraversalOrder[idx+1]
	}
	return nil
}

func getInOrderTraversalOrder(node *BinaryTree, order *[]*BinaryTree) {
	if node == nil {
		return
	}
	getInOrderTraversalOrder(node.Left, order)
	*order = append(*order, node)
	getInOrderTraversalOrder(node.Right, order)
}
