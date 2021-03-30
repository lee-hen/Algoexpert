package branchsums

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
	var sums []int
	root.traverse(&sums, 0)
	return sums
}

// https://www.sohamkamani.com/golang/arrays-vs-slices/
func (tree *BinaryTree) traverse(sums *[]int, sum int) {
	if tree != nil {
		sum += tree.Value
		tree.Left.traverse(sums, sum)
		tree.Right.traverse(sums, sum)

		if tree.Left == nil && tree.Right == nil {
			*sums = append(*sums, sum)
		}
	}
}
