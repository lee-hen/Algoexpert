package valid_starting_city

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	distances := []int{5, 25, 15, 10, 15}
	fuel := []int{1, 2, 1, 0, 3}
	mpg := 10
	expected := 4
	actual := ValidStartingCity(distances, fuel, mpg)
	require.Equal(t, expected, actual)
}
