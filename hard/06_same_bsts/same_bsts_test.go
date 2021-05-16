package same_bsts

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	arrayOne := []int{10, 15, 8, 12, 94, 81, 5, 2, 11}
	arrayTwo := []int{10, 8, 5, 15, 2, 12, 11, 94, 81}
	require.Equal(t, SameBsts(arrayOne, arrayTwo), true)
}
