package continuous_median

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	var handler = NewContinuousMedianHandler()
	handler.Insert(5)
	handler.Insert(10)
	require.Equal(t, 7.5, handler.GetMedian())
	handler.Insert(100)
	require.Equal(t, 10.0, handler.GetMedian())
}

func TestCase2(t *testing.T) {
	var handler = NewContinuousMedianHandler()
	handler.Insert(5)
	handler.Insert(10)
	handler.Insert(100)
	handler.Insert(200)
	handler.Insert(6)
	require.Equal(t, 10.0, handler.GetMedian())
}
