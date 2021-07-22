package count_inversions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	input := []int{8, 5, 2, 9, 5, 6, 3}
	expected := 12
	actual := CountInversions(input)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	input := []int{2, 3, 3, 1, 9, 5, 6}
	expected := 5
	actual := CountInversions(input)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, -1}
	expected := 6
	actual := CountInversions(input)
	require.Equal(t, expected, actual)
}

func TestCase4(t *testing.T) {
	input := []int{54, 1, 2, 3, 4}
	expected := 4
	actual := CountInversions(input)
	require.Equal(t, expected, actual)
}

func TestCase5(t *testing.T) {
	input := []int{1, 10, 2, 8, 3, 7, 4, 6, 5}
	expected := 16
	actual := CountInversions(input)
	require.Equal(t, expected, actual)
}
