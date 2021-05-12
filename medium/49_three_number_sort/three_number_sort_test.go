package three_number_sort

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := []int{0, 1, -1}
	expected := []int{0, 0, 0, 1, 1, 1, -1, -1}
	actual := ThreeNumberSort(array, order)
	require.Equal(t, expected, actual)
}
