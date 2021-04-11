package tandem_bicycle


import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	redShirtSpeeds := []int{5, 5, 3, 9, 2}
	blueShirtSpeeds := []int{3, 6, 7, 2, 1}
	fastest := false
	expected := 32
	actual := TandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest)
	require.Equal(t, expected, actual)
}
