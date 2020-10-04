package main

import "fmt"

func NewBST(value int) *BST {
	return &BST{Value: value}
}

func main() {

	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)

	expected := 13
	actual := root.FindClosestValue(12)
	fmt.Println(actual)
	fmt.Println(actual == expected)
}
