package breadth_first_search

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func NewNode(name string) *Node {
	return &Node{
		Name:     name,
		Children: []*Node{},
	}
}

func (n *Node) AddChildren(names ...string) *Node {
	for _, name := range names {
		child := Node{Name: name}
		n.Children = append(n.Children, &child)
	}
	return n
}

func TestCase1(t *testing.T) {
	var graph = NewNode("A").AddChildren("B", "C", "D")
	graph.Children[0].AddChildren("E").AddChildren("F")
	graph.Children[2].AddChildren("G").AddChildren("H")
	graph.Children[0].Children[1].AddChildren("I").AddChildren("J")
	graph.Children[2].Children[0].AddChildren("K")
	output := graph.BreadthFirstSearch([]string{})
	expected := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}
	require.Equal(t, expected, output)
}
