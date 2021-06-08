package find_loop

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func NewLinkedList(root int, children ...int) *LinkedList {
	ll := &LinkedList{root, nil}
	ll.Add(children...)
	return ll
}

func (ll *LinkedList) Add(values ...int) {
	current := ll
	for current.Next != nil {
		current = current.Next
	}
	for value := range values {
		current.Next = &LinkedList{value, nil}
		current = current.Next
	}
}

func (ll *LinkedList) GetNth(n int) *LinkedList {
	counter, current := 1, ll
	for counter < n {
		counter, current = counter+1, current.Next
	}
	return current
}

func TestCase1(t *testing.T) {
	ll := NewLinkedList(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	ll.GetNth(10).Next = ll.GetNth(5)
	output, expected := FindLoop(ll), ll.GetNth(5)
	require.Equal(t, expected, output)
}
