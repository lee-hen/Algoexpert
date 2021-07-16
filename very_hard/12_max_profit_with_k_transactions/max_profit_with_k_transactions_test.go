package max_profit_with_k_transactions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	require.Equal(t, 93, MaxProfitWithKTransactions([]int{5, 11, 3, 50, 60, 90}, 2))
}

func TestCase2(t *testing.T) {
	require.Equal(t, 95, MaxProfitWithKTransactions([]int{5, 11, 90, 50, 60, 3}, 2))
}

func TestCase3(t *testing.T) {
	require.Equal(t, 95, MaxProfitWithKTransactions([]int{9, 11, 5, 90, 50, 60, 3}, 2))
}

func TestCase4(t *testing.T) {
	require.Equal(t, 62, MaxProfitWithKTransactions([]int{1, 25, 24, 23, 12, 36, 14, 40, 31, 41, 5}, 2))
}

func TestCase5(t *testing.T) {
	require.Equal(t, 106, MaxProfitWithKTransactions([]int{50, 25, 12, 4, 3, 10, 1, 100}, 2))
}
