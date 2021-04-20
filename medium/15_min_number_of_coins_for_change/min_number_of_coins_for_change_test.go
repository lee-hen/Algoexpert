package min_number_of_coins_for_change

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	input := []int{1, 5, 10}
	actual := MinNumberOfCoinsForChange(7, input)
	require.Equal(t, 3, actual)
}
