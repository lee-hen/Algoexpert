package staircase_traversal

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	stairs := 4
	maxSteps := 2
	expected := 5
	actual := StaircaseTraversal(stairs, maxSteps)
	require.Equal(t, expected, actual)
}
