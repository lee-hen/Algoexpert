package bst_construction
// important question

type BST struct {
	Value, Size int

	Left  *BST
	Right *BST
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func (tree *BST) Remove(value int) *BST {
	return remove(tree, value)
}

// Remove
// Better solution and easy to remember
func remove(tree *BST, value int) *BST {
	if tree == nil {
		return nil
	}
	if value < tree.Value {
		tree.Left = remove(tree.Left, value)
	} else if value > tree.Value {
		tree.Right = remove(tree.Right, value)
	} else {
		if tree.Right == nil && tree.Left == nil { // 左と右全部nilの場合、nilを設定する。
			tree = nil
		} else if tree.Right == nil && tree.Left != nil { // 右だけnilの場合、左の値を設定する
			tree.Value = tree.Left.Value
			tree.Right = tree.Left.Right // 先に左の右側を設定して
			tree.Left = tree.Left.Left   // 次に左を上書きする
		} else if tree.Left == nil && tree.Right != nil { // 左だけがnilの場合、右の値を設定する
			tree.Value = tree.Right.Value
			tree.Left = tree.Right.Left   // 先に右の左側を設定して
			tree.Right = tree.Right.Right // 次に右を上書きする
		} else { // 両方側がnilじゃない時
			// 右側最小の値を探し、削除のnodeにその値を設定する
			tree.Value = tree.Right.getMinNode().Value
			// 右側最小のNodeを削除する
			tree.Right = tree.Right.removeMin()
		}
	}

	if tree != nil {
		tree.Size = 1 + tree.Left.size() + tree.Right.size()
	}
	return tree
}

func (tree *BST) getMinNode() *BST {
	if tree.Left == nil {
		return tree
	}

	return tree.Left.getMinNode()
}

func (tree *BST) removeMin() *BST {
	if tree.Left == nil {
		return tree.Right
	}

	tree.Left = tree.Left.removeMin()
	tree.Size = 1 + tree.Left.size() + tree.Right.size()
	return tree
}

// Insert
// Better solution
func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = NewBST(value)
			tree.Left.Size = 1
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = NewBST(value)
			tree.Right.Size = 1
		} else {
			tree.Right.Insert(value)
		}
	}

	tree.Size = 1 + tree.Left.size() + tree.Right.size()
	return tree
}

// Contains
// Better solution
func (tree *BST) Contains(value int) bool {
	if value < tree.Value {
		if tree.Left == nil {
			return false
		} else {
			return tree.Left.Contains(value)
		}
	} else if value > tree.Value {
		if tree.Right == nil {
			return false
		} else {
			return tree.Right.Contains(value)
		}
	}
	return true
}

func (tree *BST) size() int {
	if tree == nil {
		return 0
	}
	return tree.Size
}

func (tree *BST) Rank(value int) int {
	if tree == nil {
		return 0
	}

	if value < tree.Value {
		return tree.Left.Rank(value)
	} else if value > tree.Value {
		return 1 + tree.Left.size() + tree.Right.Rank(value)
	} else {
		return tree.Left.size()
	}
}

// IDDFS
// https://en.wikipedia.org/wiki/Iterative_deepening_depth-first_search
func (tree *BST) IDDFS() []int {
	depth := tree.Depth()
	nums := make([]int, 0)
	for i := 0; i < depth; i++ {
		leaf := make([]int, 0)
		tree.DFS(i, &leaf)
		nums = append(nums, leaf...)
	}
	return nums
}

func (tree *BST) DFS(level int, arr *[]int) {
	current := tree
	if current == nil {
		return
	}

	if level == 0 {
		*arr = append(*arr, current.Value)
		return
	}

	current.Left.DFS(level-1, arr)
	current.Right.DFS(level-1, arr)
}

func (tree *BST) Height() int {
	if tree == nil {
		return -1
	}

	return max(tree.Left.Height(), tree.Right.Height()) + 1
}

func (tree *BST) Depth() int {
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

// ----------------------------------------------------------------

// Remove1
// Better solution
func (tree *BST) Remove1(value int) *BST {
	tree.remove1(value, nil)
	return tree
}

func (tree *BST) remove1(value int, parent *BST) {
	current := tree
	for current != nil {
		if value > current.Value {
			parent = current
			current = current.Right
		} else if value < current.Value {
			parent = current
			current = current.Left
		} else {
			if current.Right != nil && current.Left != nil {
				// set right min value then recursively remove that min value
				current.Value = getMinValue(current.Right)
				current.Right.remove1(current.Value, current)
			} else if parent == nil {
				if current.Left != nil { // root node only has left side
					current.Value = current.Left.Value
					// First set to right then left
					current.Right = current.Left.Right
					current.Left = current.Left.Left
				} else if current.Right != nil { // root node only has right side
					current.Value = current.Right.Value
					// First set to Left then right
					current.Left = current.Right.Left
					current.Right = current.Right.Right
				} else {

				}
			} else if parent.Left == current { // left side
				if current.Left != nil {
					parent.Left = current.Left
				} else { // when left is nil or right is nil or right is not nil
					parent.Left = current.Right
				}
			} else if parent.Right == current { // right side
				if current.Left != nil {
					parent.Right = current.Left
				} else { // when left is nil or right is nil or right is not nil
					parent.Right = current.Right
				}
			}
			break
		}
	}
}

func getMinValue(tree *BST) int {
	current := tree
	for current.Left != nil {
		current = current.Left
	}
	return current.Value
}

// Insert1
// Better solution1
func (tree *BST) Insert1(value int) *BST {
	currentNode := tree
	for currentNode != nil {
		if value >= currentNode.Value {
			if currentNode.Right == nil {
				currentNode.Right = &BST{
					Value: value,
				}
				break
			} else {
				currentNode = currentNode.Right
			}
		} else if value < currentNode.Value {
			if currentNode.Left == nil {
				currentNode.Left = &BST{
					Value: value,
				}
				break
			} else {
				currentNode = currentNode.Left
			}
		}
	}
	return tree
}

func (tree *BST) Contains1(value int) bool {
	currentNode := tree
	for currentNode != nil {
		if value == currentNode.Value {
			return true
		}
		if value > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			currentNode = currentNode.Left
		}
	}
	return false
}

// ----------------------------------------------------------------

func (tree *BST) Remove2(value int) *BST {
	tree.remove2(value, nil)
	return tree
}
func (tree *BST) remove2(value int, parent *BST) {
	if value < tree.Value {
		if tree.Left != nil {
			tree.Left.remove2(value, tree)
		}
	} else if value > tree.Value {
		if tree.Right != nil {
			tree.Right.remove2(value, tree)
		}
	} else {
		if tree.Left != nil && tree.Right != nil {
			tree.Value = tree.Right.getMinValue1()
			tree.Right.remove2(tree.Value, tree)
		} else if parent == nil {
			if tree.Left != nil {
				tree.Value = tree.Left.Value
				tree.Right = tree.Left.Right
				tree.Left = tree.Left.Left
			} else if tree.Right != nil {
				tree.Value = tree.Right.Value
				tree.Left = tree.Right.Left
				tree.Right = tree.Right.Right
			} else {
			}
		} else if parent.Left == tree {
			if tree.Left != nil {
				parent.Left = tree.Left
			} else {
				parent.Left = tree.Right
			}
		} else if parent.Right == tree {
			if tree.Left != nil {
				parent.Right = tree.Left
			} else {
				parent.Right = tree.Right
			}
		}
	}
}

func (tree *BST) getMinValue1() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue1()
}

// My solution
// consider if removed node has two children.
// consider if removed node has one child.
// consider if removed node has no children,
// consider if removed node is root node.
func (tree *BST) remove3(value int) *BST {
	parentNode := tree
	for parentNode != nil {
		var currentNode *BST
		if value == parentNode.Value {
			currentNode = parentNode
			if currentNode.Right != nil {
				minNodeParent := findMinParentNode(currentNode.Right)
				minNode := minNodeParent.Left
				if minNode != nil {
					minNodeParent.Left = minNode.Right
					currentNode.Value = minNode.Value
				} else {
					currentNode.Value = minNodeParent.Value
					currentNode.Right = minNodeParent.Right
				}

			} else if currentNode.Left != nil {
				currentNode.Value = currentNode.Left.Value
				currentNode.Left = currentNode.Left.Left
			}
			break
		}
		if value < parentNode.Value {
			currentNode = parentNode.Left
			if currentNode != nil && currentNode.Value == value {
				if currentNode.Right != nil {
					parentNode.Left = replaceCurrentToMinNode(currentNode)
				} else if currentNode.Left != nil {
					parentNode.Left = currentNode.Left
				} else {
					parentNode.Left = currentNode.Right
				}
				break
			}
		} else {
			currentNode = parentNode.Right
			if currentNode != nil && currentNode.Value == value {
				if currentNode.Right != nil {
					parentNode.Right = replaceCurrentToMinNode(currentNode)
				} else if currentNode.Left != nil {
					parentNode.Right = currentNode.Right
				} else {
					parentNode.Right = currentNode.Left
				}
			}
		}
		parentNode = currentNode
	}
	return tree
}

func replaceCurrentToMinNode(currentNode *BST) *BST {
	parentMinNode := findMinParentNode(currentNode.Right)
	minNode := parentMinNode.Left
	if minNode != nil {
		parentMinNode.Left = minNode.Right
		currentNode.Value = minNode.Value
	} else {
		parentMinNode.Left = currentNode.Left
	}
	return parentMinNode
}

func findMinParentNode(tree *BST) *BST {
	currentNode := tree
	for currentNode.Left != nil && currentNode.Left.Left != nil {
		currentNode = currentNode.Left
	}
	return currentNode
}
