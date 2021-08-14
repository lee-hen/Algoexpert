package right_sibling_tree
// important question

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// RightSiblingTree
// O(n) time | O(d) space - where n is the number of nodes in
// the Binary Tree and d is the depth (height) of the Binary Tree
func RightSiblingTree(root *BinaryTree) *BinaryTree {
	mutate(root, nil, false)
	return root
}

func mutate(node, parent *BinaryTree, isLeftChild bool) {
	if node == nil {
		return
	}

	left, right := node.Left, node.Right
	mutate(left, node, true)
	if parent == nil {
		node.Right = nil
	} else if isLeftChild {
		node.Right = parent.Right
	} else {
		if parent.Right == nil {
			node.Right = nil
		} else {
			node.Right = parent.Right.Left
		}
	}
	mutate(right, node, false)
}


//         1
//      /
//     2-----------3
//    /           /
//   4-----5-----6-----7
//  /           /     /
// 8---9    10-11    12-13
//            /
//          14

//
//                1
//           /          \
//         2             3
//       /   \         /   \
//      4     5       6     7
//    /   \    \    /     /   \
//   8     9   10  11    12   13
//                /
//               14

// my solution
// rightSiblingTree
func rightSiblingTree(root *BinaryTree) *BinaryTree {
	depth := root.Depth()

	for i := depth; i >= 0; i-- {
		last := root.dfs(i, nil)
		if last != nil {
			last.Right = nil
		}
	}

	return root
}

func (tree *BinaryTree) dfs(level int, previous *BinaryTree) *BinaryTree {
	if tree != nil {
		if level == 0 {
			if previous != nil {
				if previous.Right != nil {
					if previous.Right.Right == tree.Right {
						previous.Right.Right = nil
					}
				}

				if previous.Left != nil {
					if previous.Left.Right == tree.Left {
						previous.Left.Right = nil
					}
				}

				previous.Right = tree
			}

			return tree
		}

		previous = tree.Left.dfs(level-1, previous)
		previous = tree.Right.dfs(level-1, previous)
	}
	return previous
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
