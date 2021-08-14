package find_closest_value
// important question

import (
	"fmt"
)

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func (tree *BST) FindClosestValue(target int) int {
	closest := tree.Value
	for currentNode := tree; currentNode != nil; {
		if Abs(target-closest) > Abs(currentNode.Value-target) {
			closest = currentNode.Value
		}

		if target < currentNode.Value {
			currentNode = currentNode.Left
		} else if target > currentNode.Value {
			currentNode = currentNode.Right
		} else {
			break
		}
	}

	return closest
}

func (tree *BST) FindClosestValue2(target int) int {
	return tree.findClosestValue(target, tree.Value)
}

func (tree *BST) findClosestValue(target int, closest int) int {
	if Abs(target-closest) > Abs(tree.Value-target) {
		closest = tree.Value
	}

	if target < tree.Value && tree.Left != nil {
		return tree.Left.findClosestValue(target, closest)
	} else if target > tree.Value && tree.Right != nil {
		return tree.Right.findClosestValue(target, closest)
	} else {
		return closest
	}
}

// My solution
func (tree *BST) findClosestValue1(target int) int {
	fmt.Println(tree.Value)

	if target < tree.Value {

		if tree.Left == nil {
			return tree.Value
		}

		if Abs(tree.Left.Value-target) > Abs(tree.Value-target) {
			if tree.Left.Right == nil {
				return tree.Value
			} else {
				if Abs(tree.Left.Right.Value-target) > Abs(tree.Value-target) {
					return tree.Value
				}
			}
		}

		return tree.Left.findClosestValue1(target)
	} else {

		if target == tree.Value {
			return tree.Value
		}

		if tree.Right == nil {
			return tree.Value
		}

		if Abs(tree.Right.Value-target) > Abs(tree.Value-target) {
			if tree.Right.Left == nil {
				return tree.Value
			} else {
				if Abs(tree.Right.Left.Value-target) > Abs(tree.Value-target) {
					return tree.Value
				}
			}
		}

		return tree.Right.findClosestValue1(target)
	}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
