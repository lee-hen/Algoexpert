package first_duplicate_value

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	input := []int{2, 1, 5, 2, 3, 3, 4}
	expected := 2
	actual := FirstDuplicateValue(input)
	require.Equal(t, expected, actual)
}
