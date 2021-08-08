package reverse_alternating_k_nodes

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	ll := &LinkedList{Value: 1}
	addMany(ll, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14})
	k := 3
	expected := []int{3, 2, 1, 4, 5, 6, 9, 8, 7, 10, 11, 12, 14, 13}
	actual := reverseAlternatingKNodes(ll, k).Values()
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	ll := &LinkedList{Value: 1}
	addMany(ll, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18})
	k := 3
	expected := []int{3, 2, 1, 4, 5, 6, 9, 8, 7, 10, 11, 12, 15, 14, 13, 16, 17, 18}
	actual := reverseAlternatingKNodes(ll, k).Values()
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	ll := &LinkedList{Value: 1}
	addMany(ll, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14})
	k := 20
	expected := []int{14, 13, 12, 11, 10, 9, 8,  7, 6, 5, 4, 3, 2, 1}
	actual := reverseAlternatingKNodes(ll, k).Values()
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	ll := &LinkedList{Value: 1}
	addMany(ll, []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14})
	k := 9
	expected := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 11, 12, 13, 14}
	actual := reverseAlternatingKNodes(ll, k).Values()
	require.Equal(t, expected, actual)
}

func addMany(ll *LinkedList, values []int) {
	current := ll
	for _, val := range values {
		current.Next = &LinkedList{Value: val}
		current = current.Next
	}
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
