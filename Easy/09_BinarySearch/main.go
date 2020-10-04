package main

import "fmt"

func main() {
	expected := 3
	output := BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 100)

	fmt.Println(output)
	fmt.Println(expected == output)
}
