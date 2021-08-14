package iterative_in_order_traversal
// important question

type BinaryTree struct {
	Value int

	Left   *BinaryTree
	Right  *BinaryTree
	Parent *BinaryTree
}

// IterativeInOrderTraversal
// O(n) time | O(1) space
func (tree *BinaryTree) IterativeInOrderTraversal(callback func(int)) {
	var previous, next *BinaryTree
	current := tree
	for current != nil {
		if previous == nil || previous == current.Parent {
			if current.Left != nil {
				next = current.Left
			} else {
				callback(current.Value)
				if current.Right != nil {
					next = current.Right
				} else {
					next = current.Parent
				}
			}
		} else if previous == current.Left {
			callback(current.Value)
			if current.Right != nil {
				next = current.Right
			} else {
				next = current.Parent
			}
		} else {
			next = current.Parent
		}
		previous, current = current, next
	}
}

//        1
//     /         \
//    2             3
//  /      \       /   \
// 4        5     6     7
//  \      /   \
//   9   10      11

// iterativeInOrderTraversal
// my solution
func (tree *BinaryTree) iterativeInOrderTraversal(callback func(int)) {
	current := tree

	var previous *BinaryTree
	for current != nil {
		if current.Left == nil || current.Left == previous {
			callback(current.Value)

			if current.Right != nil {
				previous = current
				current = current.Right
			} else if previous != nil && current.Parent == previous {
				current = previous.Parent
				if current == tree && previous == tree.Right {
					break
				}
			} else {
				previous = current
				current = current.Parent
			}
		} else if previous != nil && current.Right == previous {
			previous = current
			current = current.Parent
		} else {
			current = current.Left
		}
	}
}

// iterativeInOrderTraversal
// my other solution
func (tree *BinaryTree) iterativeInOrderTraversal1(callback func(int)) {
	current := tree
	visit := make(map[*BinaryTree]struct{})
	for current != nil {
		if current.Left == nil || visited(visit, current.Left)  {
			if !visited(visit, current) {
				callback(current.Value)
				visit[current] = struct{}{}

				if current.Right != nil {
					current = current.Right
				} else {
					current = current.Parent
				}
			} else {
				current = current.Parent
			}
		} else {
			current = current.Left
		}
	}
}

func visited(visit map[*BinaryTree]struct{}, node *BinaryTree) bool{
	if _, ok := visit[node]; ok {
		return true
	}
	return false
}
