package task_assignment

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	k := 3
	tasks := []int{1, 3, 5, 3, 1, 4}
	expected := [][]int{{4, 2}, {0, 5}, {3, 1}}
	actual := TaskAssignment(k, tasks)
	require.Equal(t, expected, actual)
}
