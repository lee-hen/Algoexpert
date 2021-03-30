package classPhotos

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	redShirtHeights := []int{5, 8, 1, 3, 4}
	blueShirtHeights := []int{6, 9, 2, 4, 5}
	expected := true
	actual := ClassPhotos(redShirtHeights, blueShirtHeights)
	require.Equal(t, expected, actual)
}
