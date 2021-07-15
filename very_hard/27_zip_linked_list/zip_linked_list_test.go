package zip_linked_list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := addMany(&LinkedList{Value: 1}, []int{2, 3, 4, 5, 6})
	expected := []int{1, 6, 2, 5, 3, 4}
	actual := ZipLinkedList(input)
	result := getValues(actual)
	require.Equal(t, expected, result)
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
