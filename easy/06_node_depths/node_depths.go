package node_depths

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return root.CalculateNodeDepths(0)
}

func (node *BinaryTree) CalculateNodeDepths(depth int) int {
	if node == nil {
		return 0
	}

	return depth + node.Left.CalculateNodeDepths(depth+1) + node.Right.CalculateNodeDepths(depth+1)
}

type Level struct {
	Root  *BinaryTree
	Depth int
}

func NodeDepths2(root *BinaryTree) int {
	stack := []Level{{Root: root, Depth: 0}}
	var top Level
	var depths int
	for len(stack) > 0 {
		top, stack = stack[len(stack)-1], stack[:len(stack)-1]
		node, depth := top.Root, top.Depth
		if node == nil {
			continue
		}
		depths += depth
		stack = append(stack, Level{Root: node.Left, Depth: depth + 1})
		stack = append(stack, Level{Root: node.Right, Depth: depth + 1})
	}
	return depths
}

// My solution
func nodeDepths1(root *BinaryTree) int {
	var depths int
	root.calculateNodeDepths(-1, &depths)
	return depths
}

func (node *BinaryTree) calculateNodeDepths(depth int, depths *int) {
	if node != nil {
		depth += 1
		node.Left.calculateNodeDepths(depth, depths)
		node.Right.calculateNodeDepths(depth, depths)
		*depths += depth
	}
}
