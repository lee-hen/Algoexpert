package non_attacking_queens

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := 4
	expected := 2
	actual := NonAttackingQueens(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := 5
	expected := 10
	actual := NonAttackingQueens(input)
	require.Equal(t, expected, actual)
}
