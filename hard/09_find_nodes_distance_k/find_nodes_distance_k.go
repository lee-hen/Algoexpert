package find_nodes_distance_k
// important question

type BinaryTree struct {
	Value, distance int

	Left, Right, parent *BinaryTree
}

// FindNodesDistanceK
// k=2
// target 3
// output 2 7 8      11
// all node values will be unique
func FindNodesDistanceK(tree *BinaryTree, target int, k int) []int {
	nodesDistanceK := make([]int, 0)
	findDistanceFromNodeToTarget(tree, target, k, &nodesDistanceK)
	return nodesDistanceK
}

func findDistanceFromNodeToTarget(node *BinaryTree, target int, k int, nodesDistanceK *[]int) int {
	if node == nil {
		return -1 // either not in the root's left subtree or it's right subtree
	}

	if node.Value == target {
		addSubtreeNodesAtDistanceK(node, 0, k, nodesDistanceK) // traverse children nodes distanceK from target node
		return 1 // return distance(1) from target to it's parent node
	}

	leftDistance := findDistanceFromNodeToTarget(node.Left, target, k, nodesDistanceK)
	rightDistance := findDistanceFromNodeToTarget(node.Right, target, k, nodesDistanceK)

	if leftDistance == k || rightDistance == k {
		*nodesDistanceK = append(*nodesDistanceK, node.Value)
	}

	// after call findDistanceFromNodeToTarget(node.Left, target, k, nodesDistanceK)
	// all node the call return is the target node's parent at root's left subtree
	if leftDistance != -1 { // target node is at the left subtree
		addSubtreeNodesAtDistanceK(node.Right, leftDistance+1, k, nodesDistanceK) // traverse parent node's right subtree,  because of node.Right so the distance needs to plus 1
		return leftDistance + 1 // the same thing about parent's parent node
	}

	// after call findDistanceFromNodeToTarget(node.Right, target, k, nodesDistanceK)
	// all node the call return is the target node's parent at root's right subtree
	if rightDistance != -1 { // target node is at the right subtree
		addSubtreeNodesAtDistanceK(node.Left, rightDistance+1, k, nodesDistanceK) // traverse parent node's left subtree
		return rightDistance + 1 // the same thing about parent's parent node
	}

	return -1 // this is not evaluate
}

func addSubtreeNodesAtDistanceK(node *BinaryTree, distance int, k int, nodesDistanceK *[]int) {
	if node == nil {
		return
	}

	if distance == k {
		*nodesDistanceK = append(*nodesDistanceK, node.Value)
	} else {
		addSubtreeNodesAtDistanceK(node.Left, distance+1, k, nodesDistanceK)
		addSubtreeNodesAtDistanceK(node.Right, distance+1, k, nodesDistanceK)
	}
}

func FindNodesDistanceK1(tree *BinaryTree, target int, k int) []int {
	nodesToParents := map[int]*BinaryTree{}
	populateNodesToParents(tree, nodesToParents, nil)
	targetNode := getNodeFromValue(target, tree, nodesToParents)

	return breadthFirstSearchForNodesDistanceK(targetNode, nodesToParents, k)
}

func breadthFirstSearchForNodesDistanceK(targetNode *BinaryTree, nodesToParents map[int]*BinaryTree, k int) []int {
	type item struct {
		node     *BinaryTree
		distance int
	}

	queue := []item{{node: targetNode, distance: 0}}
	seen := map[int]bool{targetNode.Value: true}

	var currentItem item
	for len(queue) > 0 {
		currentItem, queue = queue[0], queue[1:]
		currentNode, distanceFromTarget := currentItem.node, currentItem.distance

		if distanceFromTarget == k {
			nodesDistanceK := make([]int, 0)

			for _, i := range queue {
				nodesDistanceK = append(nodesDistanceK, i.node.Value)
			}

			nodesDistanceK = append(nodesDistanceK, currentNode.Value)
			return nodesDistanceK
		}

		connectedNodes := []*BinaryTree{currentNode.Left, currentNode.Right, nodesToParents[currentNode.Value]}
		for _, node := range connectedNodes {
			if node == nil {
				continue
			}

			if seen[node.Value] {
				continue
			}

			seen[node.Value] = true
			queue = append(queue, item{node: node, distance: distanceFromTarget + 1})
		}
	}

	return []int{}
}

func getNodeFromValue(value int, tree *BinaryTree, nodesToParents map[int]*BinaryTree) *BinaryTree {
	if tree.Value == value {
		return tree
	}

	nodeParent := nodesToParents[value]
	if nodeParent.Left != nil && nodeParent.Left.Value == value {
		return nodeParent.Left
	}
	return nodeParent.Right
}

func populateNodesToParents(node *BinaryTree, nodesToParents map[int]*BinaryTree, parent *BinaryTree) {
	if node == nil {
		return
	}

	nodesToParents[node.Value] = parent
	populateNodesToParents(node.Left, nodesToParents, node)
	populateNodesToParents(node.Right, nodesToParents, node)
}

// FindNodesDistanceK
// my solution
func findNodesDistanceK(tree *BinaryTree, target int, k int) []int {
	targetNode := &BinaryTree{}
	nodesDistanceK := make([]int, 0)
	findTargetNode(tree, &targetNode, target, k, &nodesDistanceK)

	visited := make(map[*BinaryTree]bool)
	traverseParentFromTargetNode(targetNode.parent, target, k, visited, &nodesDistanceK)
	traverseChildrenFromTargetNode(targetNode, k, &nodesDistanceK)
	return nodesDistanceK
}

func findTargetNode(tree *BinaryTree, targetNode **BinaryTree, target, k int, nodesDistanceK *[]int)  {
	if tree == nil {
		return
	}

	if tree.Left != nil {
		tree.Left.parent = tree
	}

	if tree.Right != nil {
		tree.Right.parent = tree
	}

	if tree.Value == target {
		*targetNode = tree
		return
	}

	findTargetNode(tree.Left, targetNode, target, k, nodesDistanceK)
	findTargetNode(tree.Right, targetNode, target, k, nodesDistanceK)
}

func traverseParentFromTargetNode(tree *BinaryTree, target, k int, visited map[*BinaryTree]bool, nodesDistanceK *[]int)  {
	if tree == nil || tree.Value == target {
		return
	}
	if visited[tree] {
		return
	}

	visited[tree] = true

	if tree.Right != nil && tree.Right.Value == target {
		tree.distance++
	}

	if tree.Left != nil && tree.Left.Value == target {
		tree.distance++
	}

	if tree.parent != nil && tree.parent.distance != 0 {
		tree.distance = tree.parent.distance+1
	}

	if tree.Right != nil && tree.Right.distance !=0 {
		tree.distance = tree.Right.distance+1
	}

	if tree.Left != nil && tree.Left.distance !=0 {
		tree.distance = tree.Left.distance+1
	}

	if k < tree.distance  {
		return
	}

	if tree.distance == k {
		*nodesDistanceK = append(*nodesDistanceK, tree.Value)
	}

	traverseParentFromTargetNode(tree.Left, target, k, visited, nodesDistanceK)
	traverseParentFromTargetNode(tree.Right, target, k, visited, nodesDistanceK)
	traverseParentFromTargetNode(tree.parent, target, k, visited, nodesDistanceK)
}

func traverseChildrenFromTargetNode(tree *BinaryTree, k int, nodesDistanceK *[]int) {
	if tree.Left != nil {
		current := tree
		for current.distance != k-1 {
			current.Left.distance = current.distance+1
			current = current.Left
		}

		if current != nil && current.Left != nil {
			*nodesDistanceK = append(*nodesDistanceK, current.Left.Value)
		}
		if current != nil && current.Right != nil {
			*nodesDistanceK = append(*nodesDistanceK, current.Right.Value)
		}
	}

	if tree.Right != nil {
		current := tree
		for current.distance != k-1 {
			current.Right.distance = current.distance+1
			current = current.Right
		}
		if current != nil && current.Left != nil {
			*nodesDistanceK = append(*nodesDistanceK, current.Left.Value)
		}
		if current != nil && current.Right != nil {
			*nodesDistanceK = append(*nodesDistanceK, current.Right.Value)
		}
	}

	if k == 1 && len(*nodesDistanceK) > 0 {
		*nodesDistanceK = (*nodesDistanceK)[:len(*nodesDistanceK)-2]
	}
}
