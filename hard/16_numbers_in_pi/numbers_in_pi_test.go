package numbers_in_pi

import (
	"github.com/stretchr/testify/require"
	"testing"
)

const PI = "3141592653589793238462643383279"

func TestCase1(t *testing.T) {
	numbers := []string{"314159265358979323846", "26433", "8", "3279", "314159265", "35897932384626433832", "79"}
	require.Equal(t, NumbersInPi(PI, numbers), 2)
}
