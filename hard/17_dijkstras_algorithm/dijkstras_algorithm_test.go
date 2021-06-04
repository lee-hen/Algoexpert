package dijkstras_algorithm

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	start := 0
	edges := [][][]int{
		{{1, 7}},
		{{2, 6}, {3, 20}, {4, 3}},
		{{3, 14}},
		{{4, 2}},
		{},
		{},
	}
	expected := []int{0, 7, 13, 27, 10, -1}
	actual := DijkstrasAlgorithm(start, edges)
	require.Equal(t, expected, actual)
}

func TestCase2(t *testing.T) {
	start := 1
	edges := [][][]int{
		{{1, 2}},
		{{0, 1}},
		{{3, 1}},
		{{2, 2}},
	}
	expected := []int{1, 0, -1, -1}
	actual := DijkstrasAlgorithm(start, edges)
	require.Equal(t, expected, actual)
}

func TestCase3(t *testing.T) {
	start := 3
	edges := [][][]int{
		{
			{1, 2},
			{3, 3},
			{4, 2},
		},
		{
			{0, 1},
			{6, 3},
		},
		{
			{3, 9},
		},
		{
			{0, 3},
			{1, 4},
			{4, 4},
			{8, 7},
		},
		{
			{0, 1},
			{10, 3},
		},
		{
			{7, 1},
			{8, 4},
		},
		{
			{8, 1},
		},
		{},
		{
			{7, 1},
		},
		{
			{10, 2},
		},
		{},
	}
	expected := []int{3, 4, -1, 0, 4, -1, 7, 8, 7, -1, 7}
	actual := DijkstrasAlgorithm(start, edges)
	require.Equal(t, expected, actual)
}
