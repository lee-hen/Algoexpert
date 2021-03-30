package binarysearch

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 3
	output := BinarySearch([]int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}, 100)

	fmt.Println(output)
	fmt.Println(expected == output)
}
