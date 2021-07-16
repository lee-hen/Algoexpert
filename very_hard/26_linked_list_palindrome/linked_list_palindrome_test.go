package linked_list_palindrome

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func  TestCase1(t *testing.T) {
	head := &LinkedList{Value: 0}
	head.Next = &LinkedList{Value: 1}
	head.Next.Next = &LinkedList{Value: 2}
	head.Next.Next.Next = &LinkedList{Value: 2}
	head.Next.Next.Next.Next = &LinkedList{Value: 1}
	head.Next.Next.Next.Next.Next = &LinkedList{Value: 0}
	expected := true
	actual := LinkedListPalindrome(head)
	require.Equal(t, expected, actual)
}
