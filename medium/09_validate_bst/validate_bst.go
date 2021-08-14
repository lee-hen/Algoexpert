package validate_bst
// important question

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// ValidateBst
// Better solution
func (tree *BST) ValidateBst() bool {
	return tree.validateBst(math.MinInt32, math.MaxInt32)
}

func (tree *BST) validateBst(min, max int) bool {
	// tree.Value >= max -> parent.Left.Value(tree.Value) >= parent.Value(max)
	// tree.Value < min -> parent.Right.Value(tree.Value) < parent.Value(min)
	if tree.Value >= max || tree.Value < min {
		return false
	}
	if tree.Left != nil && !tree.Left.validateBst(min, tree.Value) {
		return false
	}
	if tree.Right != nil && !tree.Right.validateBst(tree.Value, max) {
		return false
	}
	return true
}

// ValidateBst1
// my solution
func (tree *BST) ValidateBst1() bool {
	return tree.validateBst1(nil) && tree.validateBst1(nil)
}

func (tree *BST) validateBst1(parent *BST) bool {
	current := tree
	if current != nil {
		left := current.Left
		right := current.Right
		if parent != nil && parent.Left == current && right != nil && right.getMaxValue() >= parent.Value {
			return false
		}

		if parent != nil && parent.Right == current && left != nil && left.getMinValue() < parent.Value {
			return false
		}

		if left != nil && right != nil {
			if left.Value > current.Value || right.Value < current.Value {
				return false
			}
			return left.validateBst1(current) && right.validateBst1(current)
		} else if right != nil {
			if right.Value < current.Value {
				return false
			}
			return right.validateBst1(current)
		} else if left != nil {
			if left.Value > current.Value {
				return false
			}
			return left.validateBst1(current)
		}
	}

	return true
}

func (tree *BST) getMaxValue() int {
	if tree.Right == nil {
		return tree.Value
	}
	return tree.Right.getMaxValue()
}

func (tree *BST) getMinValue() int {
	if tree.Left == nil {
		return tree.Value
	}
	return tree.Left.getMinValue()
}
