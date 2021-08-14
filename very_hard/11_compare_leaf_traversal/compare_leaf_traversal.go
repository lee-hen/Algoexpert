package compare_leaf_traversal
// important question

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

// CompareLeafTraversal
// O(n + m) time | O(h1 + h2) space - where n is the number of nodes in the first
// Binary Tree, m is the number in the second, h1 is the height of the first Binary
// Tree, and h2 is the height of the second
func CompareLeafTraversal(tree1 *BinaryTree, tree2 *BinaryTree) bool {
	tree1TraversalStack := []*BinaryTree{tree1}
	tree2TraversalStack := []*BinaryTree{tree2}

	for len(tree1TraversalStack) > 0 && len(tree2TraversalStack) > 0 {
		tree1Leaf := getNextLeafNode(&tree1TraversalStack)
		tree2Leaf := getNextLeafNode(&tree2TraversalStack)

		if tree1Leaf.Value != tree2Leaf.Value {
			return false
		}
	}

	return len(tree1TraversalStack) == 0 && len(tree2TraversalStack) == 0
}

func getNextLeafNode(traversalStack *[]*BinaryTree) *BinaryTree {
	var currentNode *BinaryTree
	currentNode, *traversalStack = (*traversalStack)[len(*traversalStack)-1], (*traversalStack)[:len(*traversalStack)-1]

	for !isLeafNode(currentNode) {
		if currentNode.Right != nil {
			*traversalStack = append(*traversalStack, currentNode.Right)
		}

		// We purposely add the left node to the stack after the
		// right node so that it gets popped off the stack first.
		if currentNode.Left != nil {
			*traversalStack = append(*traversalStack, currentNode.Left)
		}
		currentNode, *traversalStack = (*traversalStack)[len(*traversalStack)-1], (*traversalStack)[:len(*traversalStack)-1]
	}

	return currentNode
}


type TreePair struct {
	First  *BinaryTree
	Second *BinaryTree
}

// CompareLeafTraversal1
// O(n + m) time | O(max(h1, h2)) space - where n is the number of nodes in the first
// Binary Tree, m is the number in the second, h1 is the height of the first Binary
// Tree, and h2 is the height of the second
func CompareLeafTraversal1(tree1 *BinaryTree, tree2 *BinaryTree) bool {
	tree1LeafNodesLinkedList := connectLeafNodes(tree1, nil, nil).First
	tree2LeafNodesLinkedList := connectLeafNodes(tree2, nil, nil).First

	list1CurrentNode := tree1LeafNodesLinkedList
	list2CurrentNode := tree2LeafNodesLinkedList

	for list1CurrentNode != nil && list2CurrentNode != nil {
		if list1CurrentNode.Value != list2CurrentNode.Value {
			return false
		}

		list1CurrentNode = list1CurrentNode.Right
		list2CurrentNode = list2CurrentNode.Right
	}

	return list1CurrentNode == nil && list2CurrentNode == nil
}

func connectLeafNodes(currentNode *BinaryTree, head *BinaryTree, previousNode *BinaryTree) TreePair {
	if currentNode == nil {
		return TreePair{head, previousNode}
	}
	newHead := head
	newPreviousNode := previousNode

	if isLeafNode(currentNode) {
		if previousNode == nil {
			newHead = currentNode
		} else {
			previousNode.Right = currentNode
		}

		newPreviousNode = currentNode
	}

	leftPair := connectLeafNodes(currentNode.Left, newHead, newPreviousNode)
	leftHead, leftPreviousNode := leftPair.First, leftPair.Second
	return connectLeafNodes(currentNode.Right, leftHead, leftPreviousNode)
}

func isLeafNode(node *BinaryTree) bool {
	return node.Left == nil && node.Right == nil
}

// compareLeafTraversal
// naive approach
// space complexity O((n+1)/2 + (m+1)/2)
func compareLeafTraversal(tree1 *BinaryTree, tree2 *BinaryTree) bool {
	leaf1 := make([]int, 0)
	tree1.dfs(&leaf1)
	leaf2 := make([]int, 0)
	tree2.dfs(&leaf2)

	if len(leaf1) != len(leaf2) {
		return false
	}

	for i := 0; i < len(leaf1); i++ {
		if leaf1[i] != leaf2[i] {
			return false
		}
	}

	return true
}

func (tree *BinaryTree) dfs(arr *[]int) {
	current := tree
	if current == nil {
		return
	}

	if isLeafNode(current) {
		*arr = append(*arr, current.Value)
		return
	}

	current.Left.dfs(arr)
	current.Right.dfs(arr)
	return
}
