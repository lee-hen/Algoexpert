package sum_of_linked_lists

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	ll1 := addMany(&LinkedList{Value: 2}, []int{4, 7, 1})
	ll2 := addMany(&LinkedList{Value: 9}, []int{4, 5})
	expected := addMany(&LinkedList{Value: 1}, []int{9, 2, 2})
	actual := SumOfLinkedLists(ll1, ll2)
	require.Equal(t, expected, actual)
}


func TestCase2(t *testing.T) {
	ll1 := addMany(&LinkedList{}, []int{0})
	ll2 := addMany(&LinkedList{}, []int{0})
	expected := &LinkedList{Value: 0}
	actual := SumOfLinkedLists(ll1.Next, ll2.Next)
	require.Equal(t, expected, actual)
}

func addMany(linkedList *LinkedList, values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
	return linkedList
}

func getValues(linkedList *LinkedList) []int {
	values := []int{}
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}
