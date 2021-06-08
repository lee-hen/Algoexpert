package laptop_rentals

import (
	"github.com/stretchr/testify/require"

	"testing"
)

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
