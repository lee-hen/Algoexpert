package boggle_board

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	board := [][]rune{
		{'t', 'h', 'i', 's', 'i', 's', 'a'},
		{'s', 'i', 'm', 'p', 'l', 'e', 'x'},
		{'b', 'x', 'x', 'x', 'x', 'e', 'b'},
		{'x', 'o', 'g', 'g', 'l', 'x', 'o'},
		{'x', 'x', 'x', 'D', 'T', 'r', 'a'},
		{'R', 'E', 'P', 'E', 'A', 'd', 'x'},
		{'x', 'x', 'x', 'x', 'x', 'x', 'x'},
		{'N', 'O', 'T', 'R', 'E', '-', 'P'},
		{'x', 'x', 'D', 'E', 'T', 'A', 'E'},
	}
	words := []string{"this", "is", "not", "a", "simple", "boggle", "board", "test", "REPEATED", "NOTRE-PEATED"}
	expected := []string{"this", "is", "a", "simple", "boggle", "board", "NOTRE-PEATED"}
	output := BoggleBoard(board, words)
	require.ElementsMatch(t, expected, output)
}
