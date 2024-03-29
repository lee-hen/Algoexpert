package knapsack_problem

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []interface{}{10, []int{0, 2}}
	output := KnapsackProblem([][]int{{6, 7}, {1, 2}, {4, 3}, {5, 6} }, 10)
	require.Equal(t, expected, output)
}

func TestCase2(t *testing.T) {
	expected := []interface{}{10, []int{1, 3}}
	output := KnapsackProblem([][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}, 10)
	require.Equal(t, expected, output)
}

func TestCase3(t *testing.T) {
	expected := []interface{}{10, []int{0, 1, 2}}
	output := KnapsackProblem([][]int{{1, 2}, {4, 3}, {5, 6}, {6, 7}}, 11)
	require.Equal(t, expected, output)
}
