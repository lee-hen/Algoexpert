package remove_duplicates_from_linkedlist

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	actual := &LinkedList{Value: 1}
	expected := &LinkedList{Value: 1}

	actual.addMany([]int{1, 3, 4, 4, 4, 5, 6, 6})
	expected.addMany([]int{3, 4, 5, 6})

	actual.RemoveDuplicatesFromLinkedList()
	require.Equal(t, expected.getValues(), actual.getValues())
}
