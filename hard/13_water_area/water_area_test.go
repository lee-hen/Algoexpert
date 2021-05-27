package water_area

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := 48
	output := WaterArea([]int{0, 8, 0, 0, 5, 0, 0, 10, 0, 0, 1, 1, 0, 3})
	require.Equal(t, expected, output)
}
