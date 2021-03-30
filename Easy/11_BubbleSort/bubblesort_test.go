package bubblesort

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	output := BubbleSort([]int{8, 5, 2, 9, 5, 6, 3})
	fmt.Println(output)
	fmt.Println(expected)
}
