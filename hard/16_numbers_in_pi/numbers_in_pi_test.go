package numbers_in_pi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	numbers := []string{"314159265358979323846", "26433", "8", "3279", "314159265", "35897932384626433832", "79"}
	require.Equal(t, 2, NumbersInPi("3141592653589793238462643383279", numbers))
}

func TestCase2(t *testing.T) {
	numbers := []string{"3141", "5", "31", "2", "4159", "9", "42"}
	require.Equal(t,2, NumbersInPi("3141592", numbers))
}

func TestCase3(t *testing.T) {
	numbers := []string{"314159265358979323846", "327", "26433", "8", "3279", "9", "314159265", "35897932384626433832", "79"}
	require.Equal(t,2, NumbersInPi("3141592653589793238462643383279", numbers))
}

func TestCase4(t *testing.T) {
	numbers := []string{"3", "1", "4", "592", "65", "55", "35", "8", "9793", "2384626", "83279"}
	require.Equal(t,13, NumbersInPi("3141592653589793238462643383279", numbers))
}

// 3141592653589793238462643383279
// 3141592653589793238462643383279