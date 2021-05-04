package bst_traversal

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(arr []int) []int {
	current := tree
	if current == nil {
		return arr
	}

	arr = current.Left.InOrderTraverse(arr)
	arr = append(arr, current.Value)
	arr = current.Right.InOrderTraverse(arr)
	return arr
}

func (tree *BST) PreOrderTraverse(arr []int) []int {
	current := tree
	if current == nil {
		return arr
	}

	arr = append(arr, current.Value)
	arr = current.Left.PreOrderTraverse(arr)
	arr = current.Right.PreOrderTraverse(arr)
	return arr
}

func (tree *BST) PostOrderTraverse(arr []int) []int {
	current := tree
	if current == nil {
		return arr
	}

	arr = current.Left.PostOrderTraverse(arr)
	arr = current.Right.PostOrderTraverse(arr)
	arr = append(arr, current.Value)
	return arr
}
