package invert_binary_tree

type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) InvertBinaryTree() {
	current := tree
	if current != nil {
		current.Right, current.Left = current.Left, current.Right
		if current.Left != nil {
			current.Left.InvertBinaryTree()
		}
		if current.Right != nil {
			current.Right.InvertBinaryTree()
		}
	}
}
