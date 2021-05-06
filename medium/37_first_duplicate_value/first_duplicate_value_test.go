package first_duplicate_value

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	input := []int{7, 6, 5, 3, 6, 4, 3, 5, 2}
	expected := 6
	actual := FirstDuplicateValue(input)
	require.Equal(t, expected, actual)
}
