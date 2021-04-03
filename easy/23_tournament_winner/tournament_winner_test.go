package tournament_winner

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	expected := "Python"
	actual := TournamentWinner(competitions, results)
	require.Equal(t, expected, actual)
}