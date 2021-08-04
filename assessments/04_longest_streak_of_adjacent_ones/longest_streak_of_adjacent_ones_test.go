package longest_streak_of_adjacent_ones

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	array := []int{1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 1, 0, 1}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, 8, actual)
}

func TestCase2(t *testing.T) {
	array := []int{0, 0, 0, 1, 1}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, 2, actual)
}

func TestCase3(t *testing.T) {
	array := []int{1, 1, 1, 0, 0}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, 3, actual)
}

func TestCase4(t *testing.T) {
	array := []int{1, 0, 1, 1, 1,0,1,1,1, 0}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, 5, actual)
}

func TestCase5(t *testing.T) {
	array := []int{}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, -1, actual)
}

func TestCase6(t *testing.T) {
	array := []int{0, 0 , 0, 0, 0, 0}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, 0, actual)
}

func TestCase7(t *testing.T) {
	array := []int{1, 1 , 1, 1, 1, 1}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, -1, actual)
}

func TestCase8(t *testing.T) {
	array := []int{0, 1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	actual := LongestStreakOfAdjacentOnes(array)
	require.Equal(t, 36, actual)
}
