package balance_index

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	array := []int{0, 9, -8, 2, 7, 1, 11, -2, 1}
	actual := BalanceIndex(array)
	require.Equal(t, 5, actual)
}


func TestCase2(t *testing.T) {
	array := []int{0, 1, -1}
	actual := BalanceIndex(array)
	require.Equal(t, 0, actual)
}

func TestCase3(t *testing.T) {
	array := []int{-1, 1, 0}
	actual := BalanceIndex(array)
	require.Equal(t, 2, actual)
}

func TestCase4(t *testing.T) {
	array := []int{-22, 0, 9, -8, 2, 7, 1, 11, -2, 1, 1, 30}
	actual := BalanceIndex(array)
	require.Equal(t, 11, actual)
}

func TestCase5(t *testing.T) {
	array := []int{5, 5, -4, -9, -6, 6, -8, 1, 7, 3, -9}
	actual := BalanceIndex(array)
	require.Equal(t, 10, actual)
}
