package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := []int{2, 3, 5, 5, 6, 8, 9}
	output := InsertionSort([]int{8, 5, 2, 9, 5, 6, 3})
	require.Equal(t, expected, output)
}
