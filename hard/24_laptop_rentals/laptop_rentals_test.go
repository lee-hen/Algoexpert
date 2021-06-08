package laptop_rentals

import (
	"github.com/stretchr/testify/require"

	"testing"
)

//[0, 2]
//[1, 4]
//[4, 6]
//[0, 4]
//[7, 8]
//[9, 11]
//[3, 10]

// [0, 2] [3, 10]                   1
// [0, 4] [4, 6] [7, 8] [9, 11]     1
// [1, 4]                           1

//                              [0, 2]
//               [0, 4]                             [1, 4]
//      [3, 10]         [4, 6]             [7, 8]              [9, 11]
func TestCase1(t *testing.T) {
	input := [][]int{
		{0, 2},
		{1, 4},
		{4, 6},
		{0, 4},
		{7, 8},
		{9, 11},
		{3, 10},
	}
	expected := 3
	actual := LaptopRentals(input)
	require.Equal(t, expected, actual)
}


//[0, 5],
//[2, 4],
//[4, 7],
//[5, 7],
//[9, 20],
//[3, 15],
//[6, 10]

//                              [0, 5]
//               [2, 4]                             [3, 15]
//      [5, 7]          [9, 20]             [4, 7]              [6, 10]


// [0, 5]  [5, 7] [9, 20]
// [2, 4] [4, 7]
// [3, 15]
// [6, 10]
func TestCase2(t *testing.T) {
	input := [][]int{
		{0, 5},
		{2, 4},
		{4, 7},
		{5, 7},
		{9, 20},
		{3, 15},
		{6, 10},
	}
	expected := 4
	actual := LaptopRentals(input)
	require.Equal(t, expected, actual)
}
