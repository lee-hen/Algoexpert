package shift_linked_list

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func newLinkedList(n int) *LinkedList { return &LinkedList{Value: n} }

func linkedListToArray(head *LinkedList) []int {
	array := []int{}
	current := head
	for current != nil {
		array = append(array, current.Value)
		current = current.Next
	}
	return array
}

func TestCase1(t *testing.T) {
	head := newLinkedList(0)
	head.Next = newLinkedList(1)
	head.Next.Next = newLinkedList(2)
	head.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next = newLinkedList(4)
	head.Next.Next.Next.Next.Next = newLinkedList(5)
	result := ShiftLinkedList(head, 2)
	array := linkedListToArray(result)

	expected := []int{4, 5, 0, 1, 2, 3}
	require.Equal(t, expected, array)
}
