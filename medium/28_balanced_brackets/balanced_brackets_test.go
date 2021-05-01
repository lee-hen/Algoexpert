package balanced_brackets

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func  TestCase1(t *testing.T) {
	expected := true
	output := BalancedBrackets("([])(){}(())()()")
	require.Equal(t, expected, output)
}
