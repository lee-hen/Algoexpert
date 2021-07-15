package max_profit_with_k_transactions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	require.Equal(t, 93, MaxProfitWithKTransactions([]int{5, 11, 3, 50, 60, 90}, 2))
}
