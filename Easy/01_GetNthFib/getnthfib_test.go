package getnthfib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := 5
	output := GetNthFib(6)
	require.Equal(t, expected, output)
}
