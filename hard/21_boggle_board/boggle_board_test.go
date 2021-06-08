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

func TestCase2(t *testing.T) {
	board := [][]rune{
		{'f', 't', 'r', 'o', 'p', 'i', 'k', 'b', 'o'},
		{'r', 'w', 'l', 'p', 'e', 'u', 'e', 'a', 'b'},
		{'j', 'o', 't', 's', 'e', 'l', 'f', 'l', 'p'},
		{'s', 'z', 'u', 't', 'h', 'u', 'o', 'p', 'i'},
		{'k', 'a', 'e', 'g', 'n', 'd', 'r', 'g', 'a'},
		{'h', 'n', 'l', 's', 'a', 't', 'e', 't', 'x'},
	}

	words := []string{"frozen", "rotten", "teleport", "city", "zutgatz", "kappa", "before", "rope", "obligate", "annoying"}
	expected := []string{"rotten", "obligate", "before", "frozen", "kappa", "rope", "teleport"}
	output := BoggleBoard(board, words)
	require.ElementsMatch(t, expected, output)
}

func TestCase3(t *testing.T) {
	board := [][]rune{
		{'c', 'o', 'm'},
		{'r', 'p', 'l'},
		{'c', 'i', 't'},
		{'o', 'a', 'e'},
		{'f', 'o', 'd'},
		{'z', 'r', 'b'},
		{'g', 'i', 'a'},
		{'o', 'a', 'g'},
		{'f', 's', 'z'},
		{'t', 'e', 'i'},
		{'t', 'w', 'd'},
	}
	words := []string{"commerce", "complicated", "twisted", "zigzag", "comma", "foobar", "baz", "there"}
	expected := []string{"complicated", "foobar", "twisted", "zigzag"}

	output := BoggleBoard(board, words)
	require.ElementsMatch(t, expected, output)
}
