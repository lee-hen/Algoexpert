package reverse_linked_list

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
	values := make([]int, 0)
	current := ll
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func TestCase1(t *testing.T) {
	input := NewLinkedList(0, 1, 2, 3, 4, 5)
	output := ReverseLinkedList(input)
	expected := []int{5, 4, 3, 2, 1, 0}
	require.NotNil(t, output)
	require.Equal(t, expected, output.ToArray())
}
