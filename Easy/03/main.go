package main

import "fmt"

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{22, 25, 6}
	fmt.Println(IsValidSubsequence(array, sequence))
}
