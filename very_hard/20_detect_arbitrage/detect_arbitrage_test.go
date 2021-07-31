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

func TestCase2(t *testing.T) {
	input := [][]float64{
		{1.0, 1.27, 0.718},
		{0.74,  1.0,  0.56},
		{1.39, 1.77,   1.0},
	}
	expected := false
	actual := DetectArbitrage(input)
	require.Equal(t, expected, actual)
}
