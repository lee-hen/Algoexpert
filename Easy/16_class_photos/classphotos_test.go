package classPhotos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	redShirtHeights := []int{9, 8, 11, 3, 4}
	blueShirtHeights := []int{10, 9, 12, 4, 5}
	expected := true
	actual := ClassPhotos(redShirtHeights, blueShirtHeights)
	require.Equal(t, expected, actual)
}
