package rectangle_mania

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	coords := [][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}, {2, 1}, {2, 0}, {3, 1}, {3, 0}}
	require.Equal(t, RectangleMania(coords), 6)
}
