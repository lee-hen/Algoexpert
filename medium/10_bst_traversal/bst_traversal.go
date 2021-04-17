package bst_traversal

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	current := tree
	if current.Left != nil {
		array = current.Left.InOrderTraverse(array)
	}
	if current != nil {
		array = append(array, current.Value)
	}

	if current.Right != nil {
		array = current.Right.InOrderTraverse(array)
	}
	return array
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	current := tree
	if current != nil {
		array = append(array, current.Value)
	}
	if current.Left != nil {
		array = current.Left.PreOrderTraverse(array)
	}
	if current.Right != nil {
		array = current.Right.PreOrderTraverse(array)
	}
	return array
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	current := tree
	if current.Left != nil {
		array = current.Left.PostOrderTraverse(array)
	}
	if current.Right != nil {
		array = current.Right.PostOrderTraverse(array)
	}
	if current != nil {
		array = append(array, current.Value)
	}
	return array
}
