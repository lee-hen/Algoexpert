package merge_linked_lists

import (
	"github.com/stretchr/testify/require"

	"testing"
)

func NewLinkedList(val int, others ...int) *LinkedList {
	ll := &LinkedList{Value: val}
	current := ll
	for _, other := range others {
		current.Next = &LinkedList{Value: other}
		current = current.Next
	}
	return ll
}

func (ll *LinkedList) ToArray() []int {
	vals := []int{}
	current := ll
	for current != nil {
		vals = append(vals, current.Value)
		current = current.Next
	}
	return vals
}

func TestCase1(t *testing.T) {
	list1 := NewLinkedList(2, 6, 7, 8)
	list2 := NewLinkedList(1, 3, 4, 5, 9, 10)
	output := MergeLinkedLists(list1, list2)
	expectedNodes := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	require.Equal(t, output.ToArray(), expectedNodes)
}
