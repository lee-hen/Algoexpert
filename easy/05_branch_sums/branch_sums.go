package branch_sums
// important question

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

func NewBinaryTree(root int, values ...int) *BinaryTree {
	tree := &BinaryTree{Value: root}
	tree.Insert(values, 0)
	return tree
}

func (tree *BinaryTree) Insert(values []int, i int) *BinaryTree {
	if i >= len(values) {
		return tree
	}
	val := values[i]

	queue := []*BinaryTree{tree}
	for len(queue) > 0 {
		var current *BinaryTree
		current, queue = queue[0], queue[1:]
		if current.Left == nil {
			current.Left = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Left)

		if current.Right == nil {
			current.Right = &BinaryTree{Value: val}
			break
		}
		queue = append(queue, current.Right)
	}

	tree.Insert(values, i+1)
	return tree
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
