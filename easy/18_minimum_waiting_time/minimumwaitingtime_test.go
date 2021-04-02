package minimum_waiting_time

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T)  {
	queries := []int{3, 2, 1, 2, 6}
	expected := 17
	actual := MinimumWaitingTime(queries)
	require.Equal(t, expected, actual)
}
