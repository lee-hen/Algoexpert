package rearrange_linked_list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newLinkedList(n int) *LinkedList { return &LinkedList{Value: n} }

func linkedListToArray(head *LinkedList) []int {
	array := make([]int, 0)
	current := head
	for current != nil {
		array = append(array, current.Value)
		current = current.Next
	}
	return array
}

func TestCase1(t *testing.T) {
	head := newLinkedList(3)
	head.Next = newLinkedList(0)
	head.Next.Next = newLinkedList(5)
	head.Next.Next.Next = newLinkedList(2)
	head.Next.Next.Next.Next = newLinkedList(1)
	head.Next.Next.Next.Next.Next = newLinkedList(4)
	result := RearrangeLinkedList(head, 3)
	array := linkedListToArray(result)

	expected := []int{0, 2, 1, 3, 5, 4}
	require.Equal(t, expected, array)
}
