package run_length_encoding

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "AAAAAAAAAAAABfBCCCCDDssdddddddddd"
	expected := "9A3A1B1f1B4C2D2s9d1d"
	actual := RunLengthEncoding(input)
	require.Equal(t, expected, actual)
}
