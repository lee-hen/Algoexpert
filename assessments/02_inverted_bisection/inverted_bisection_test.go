package inverted_bisection

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	test := &LinkedList{Value: 0}
	test.Next = &LinkedList{Value: 1}
	test.Next.Next = &LinkedList{Value: 2}
	test.Next.Next.Next = &LinkedList{Value: 3}
	test.Next.Next.Next.Next = &LinkedList{Value: 4}
	test.Next.Next.Next.Next.Next = &LinkedList{Value: 5}

	output := InvertedBisection(test)

	actual := output.ToArray()
	expected := []int{2, 1, 0, 5, 4, 3}
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	test := &LinkedList{Value: 0}
	test.Next = &LinkedList{Value: 1}
	test.Next.Next = &LinkedList{Value: 2}
	test.Next.Next.Next = &LinkedList{Value: 3}
	test.Next.Next.Next.Next = &LinkedList{Value: 4}
	test.Next.Next.Next.Next.Next = &LinkedList{Value: 5}
	test.Next.Next.Next.Next.Next.Next = &LinkedList{Value: 6}

	output := InvertedBisection(test)

	actual := output.ToArray()
	expected := []int{2, 1, 0, 3, 6, 5, 4}
	require.Equal(t, expected, actual)
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
