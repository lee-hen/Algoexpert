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

func TestCase2(t *testing.T) {
	head := newLinkedList(0)
	head.Next = newLinkedList(3)
	head.Next.Next = newLinkedList(2)
	head.Next.Next.Next = newLinkedList(1)
	head.Next.Next.Next.Next = newLinkedList(4)
	head.Next.Next.Next.Next.Next = newLinkedList(5)
	head.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-1)
	head.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-2)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(6)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(7)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(2)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-9000)

	result := RearrangeLinkedList(head, -9000)
	array := linkedListToArray(result)

	expected := []int{-9000, 0, 3, 2, 1, 4, 5, 3, -1, -2, 3, 6, 7, 3, 2}
	require.Equal(t, expected, array)
}

func TestCase3(t *testing.T) {
	head := newLinkedList(0)
	head.Next = newLinkedList(3)
	head.Next.Next = newLinkedList(2)
	head.Next.Next.Next = newLinkedList(1)
	head.Next.Next.Next.Next = newLinkedList(4)
	head.Next.Next.Next.Next.Next = newLinkedList(5)
	head.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-1)
	head.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-2)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(6)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(7)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(2)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-9000)
	result := RearrangeLinkedList(head, 2)
	array := linkedListToArray(result)

	expected := []int{0, 1, -1, -2, -9000, 2, 2, 3, 4, 5, 3, 3, 6, 7, 3}
	require.Equal(t, expected, array)
}

func TestCase4(t *testing.T) {
	head := newLinkedList(0)
	head.Next = newLinkedList(3)
	head.Next.Next = newLinkedList(2)
	head.Next.Next.Next = newLinkedList(1)
	head.Next.Next.Next.Next = newLinkedList(4)
	head.Next.Next.Next.Next.Next = newLinkedList(5)
	head.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-1)
	head.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-2)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(6)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(7)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(2)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(-9000)
	result := RearrangeLinkedList(head, 3)
	array := linkedListToArray(result)

	expected := []int{0, 2, 1, -1, -2, 2, -9000, 3, 3, 3, 3, 4, 5, 6, 7}
	require.Equal(t, expected, array)
}

func TestCase5(t *testing.T) {
	head := newLinkedList(0)
	head.Next = newLinkedList(2)
	head.Next.Next = newLinkedList(3)
	head.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next = newLinkedList(3)
	head.Next.Next.Next.Next.Next.Next = newLinkedList(4)
	head.Next.Next.Next.Next.Next.Next.Next = newLinkedList(5)
	head.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(6)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(7)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(8)
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = newLinkedList(9)
	result := RearrangeLinkedList(head, 3)
	array := linkedListToArray(result)

	expected := []int{0, 2, 3, 3, 3, 3, 4, 5, 6, 7, 8, 9}
	require.Equal(t, expected, array)
}

