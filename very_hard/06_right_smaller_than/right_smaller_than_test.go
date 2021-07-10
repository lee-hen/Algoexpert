package right_smaller_than

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	array := []int{8, 5, 11, -1, 3, 4, 2}
	expected := []int{5, 4, 4, 0, 1, 1, 0}
	actual := RightSmallerThan(array)
	require.Equal(t, expected, actual)
}
