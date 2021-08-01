package number_of_binary_tree_topologies

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := 5
	output := NumberOfBinaryTreeTopologies(3)
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := 14
	output := NumberOfBinaryTreeTopologies(4)
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := 42
	output := NumberOfBinaryTreeTopologies(5)
	require.Equal(t, expected, output)
}
