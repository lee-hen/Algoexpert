package min_height_bst
// important question

// MinHeightBST
// better solution
func MinHeightBST(array []int) *BST {
	return constructMinHeightBst(array, 0, len(array)-1)
}
func constructMinHeightBst(array []int, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}
	midIdx := (startIdx + endIdx) / 2
	bst := &BST{Value: array[midIdx]}
	bst.Left = constructMinHeightBst(array, startIdx, midIdx-1)
	bst.Right = constructMinHeightBst(array, midIdx+1, endIdx)
	return bst
}

// MinHeightBST3
// better solution
func MinHeightBST3(array []int) *BST {
	return constructMinHeightBst1(array, nil, 0, len(array)-1)
}
func constructMinHeightBst1(array []int, bst *BST, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}
	midIdx := (startIdx + endIdx) / 2
	newBstNode := &BST{Value: array[midIdx]}
	if bst == nil {
		bst = newBstNode
	} else {
		if array[midIdx] < bst.Value {
			bst.Left = newBstNode
			bst = bst.Left
		} else {
			bst.Right = newBstNode
			bst = bst.Right
		}
	}
	constructMinHeightBst1(array, bst, startIdx, midIdx-1)
	constructMinHeightBst1(array, bst, midIdx+1, endIdx)
	return bst
}


// MinHeightBST2
// easy to understand
func MinHeightBST2(array []int) *BST {
	tree := &BST{Value: array[len(array)/2]}
	if len(array) == 1 {
		return tree
	}
	tree.insertBST(0, len(array)/2-1, array)
	tree.insertBST(len(array)/2+1, len(array)-1, array)
	return tree
}

func (tree *BST) insertBST(left, right int, array []int)  {
	middle := (left + right) / 2
	if left <= right {
		tree.Insert(array[middle])
		tree.insertBST(middle+1, right, array)
		tree.insertBST(left, middle-1, array)
	}
}

// MinHeightBST1
// my solution
func MinHeightBST1(array []int) *BST {
	tree := &BST{Value: array[len(array)/2]}
	if len(array) == 1 {
		return tree
	}
	tree.constructMinHeightBST(0, len(array)/2, array)
	tree.constructMinHeightBST(len(array)/2, len(array), array)
	return tree
}

func (tree *BST) constructMinHeightBST(left, right int, array []int)  {
	middle := (left + right) / 2
	if middle == 0 {
		tree.Insert(array[middle])
	}
	if left != middle {
		tree.Insert(array[middle])
		tree.constructMinHeightBST(middle, right, array)
		tree.constructMinHeightBST(left, middle, array)
	}
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
