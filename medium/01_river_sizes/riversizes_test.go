package river_sizes

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := []int{2, 1, 21, 5, 2, 1}
	input := [][]int{
		{1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 1, 0},
		{1, 0, 1, 0, 0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 1, 0, 1, 1, 0, 1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1, 1, 0, 0, 0, 1, 0, 0},
		{1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 0, 1},
	}
	output := RiverSizes(input)
	fmt.Println("0-------------------------------")
	t.Log(output)
	t.Log(expected)

	sort.Ints(output)
	if !reflect.DeepEqual(expected, output) {
		t.Fail()
	}
}
