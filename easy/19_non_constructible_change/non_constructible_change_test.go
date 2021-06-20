package non_constructible_change

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{5, 7, 1, 1, 2, 3, 22}
	expected := 20
	actual := NonConstructibleChange(input)
	require.Equal(t, expected, actual)
}