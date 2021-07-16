package detect_arbitrage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := [][]float64{
		{1.0, 0.8631, 0.5903},
		{1.1586, 1.0, 0.6849},
		{1.6939, 1.46, 1.0},
	}
	expected := true
	actual := DetectArbitrage(input)
	require.Equal(t, expected, actual)
}
