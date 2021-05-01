package youngest_common_ancestor

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func (tree *AncestralTree) addAsAncestor(descendants ...*AncestralTree) {
	for _, descendant := range descendants {
		descendant.Ancestor = tree
	}
}

func getTrees() map[rune]*AncestralTree {
	trees := map[rune]*AncestralTree{}
	for _, r := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		trees[r] = &AncestralTree{Name: string(r)}
	}
	return trees
}

func TestCase1(t *testing.T) {
	trees := getTrees()
	trees['A'].addAsAncestor(trees['B'], trees['C'])
	trees['B'].addAsAncestor(trees['D'], trees['E'])
	trees['D'].addAsAncestor(trees['H'], trees['I'])
	trees['C'].addAsAncestor(trees['F'], trees['G'])
	yca := GetYoungestCommonAncestor(trees['A'], trees['E'], trees['I'])
	require.Equal(t, trees['B'], yca)
}
