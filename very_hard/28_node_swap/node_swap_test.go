package node_swap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	head := addMany(&LinkedList{Value: 0}, []int{1, 2, 3, 4, 5})
	expected := []int{1, 0, 3, 2, 5, 4}
	actual := NodeSwap(head)
	require.Equal(t, expected, actual.Values())
}

func TestCase2(t *testing.T) {
	head := addMany(&LinkedList{Value: 0}, []int{1, 2, 3, 4})
	expected := []int{1, 0, 3, 2, 4}
	actual := NodeSwap(head)
	require.Equal(t, expected, actual.Values())
}

func addMany(ll *LinkedList, values []int) *LinkedList {
	current := ll
	for _, val := range values {
		current.Next = &LinkedList{Value: val}
		current = current.Next
	}
	return ll
}

func (ll *LinkedList) Values() []int {
	values := []int{}
	current := ll
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}
