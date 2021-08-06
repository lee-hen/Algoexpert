package max_subsequence_dot_product

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	arrayOne := []int{4, 7, 9, -6, 6}
	arrayTwo := []int{5, 1, -1, -3, -2, -10}
	actual := MaxSubsequenceDotProduct(arrayOne, arrayTwo)
	require.Equal(t, 105, actual)
}
