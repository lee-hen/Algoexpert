package maximize_expression

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := []int{3, 6, 1, -3, 2, 7}
	expected := 4
	actual := MaximizeExpression(input)
	require.Equal(t, expected, actual)
}
