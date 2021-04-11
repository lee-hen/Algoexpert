package river_sizes

import (
	"reflect"
	"sort"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{1, 2, 2, 2, 5}
	input := [][]int{
		{1, 0, 0, 1, 0},
		{1, 0, 1, 0, 0},
		{0, 0, 1, 0, 1},
		{1, 0, 1, 0, 1},
		{1, 0, 1, 1, 0},
	}
	output := RiverSizes(input)
	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}
//  0   0  0  0  0
//0 {1, 0, 0, 1, 0},0
//0 {1, 0, 1, 0, 0},0
//0 {0, 0, 1, 0, 1},0
//0 {1, 0, 1, 0, 1},0
//0 {1, 0, 1, 1, 0},0
//   0  0  0  0  0