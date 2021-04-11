package move_element_to_end

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestCase1(t *testing.T) {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	expected := []int{1, 3, 4, 2, 2, 2, 2, 2}
	output := MoveElementToEnd(array, toMove)
	sort.Ints(output[0:3])
	require.Equal(t, expected, output)
}