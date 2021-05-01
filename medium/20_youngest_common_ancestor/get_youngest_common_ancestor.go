package youngest_common_ancestor

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

// GetYoungestCommonAncestor
// better solution is the same as this
func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	if descendantOne.Ancestor == nil {
		return descendantOne
	}

	if descendantTwo.Ancestor == nil {
		return descendantTwo
	}

	depth1 := depth(descendantOne, top)
	depth2 := depth(descendantTwo, top)

	var ancestor *AncestralTree
	if diff := depth1-depth2; diff > 0 {
		ancestor = findCommonAncestor(descendantOne, descendantTwo, diff)
	} else {
		ancestor = findCommonAncestor(descendantTwo, descendantOne, diff*-1)
	}

	return ancestor
}

func findCommonAncestor(current, other *AncestralTree, depth int) *AncestralTree {
	diff := depth
	for diff > 0 {
		current = current.Ancestor
		diff--
	}

	for current != nil && current != other {
		current = current.Ancestor
		other = other.Ancestor
	}
	return current
}

func depth(current, top *AncestralTree) int {
	var counter int
	for current.Ancestor != top {
		current = current.Ancestor
		counter++
	}
	return counter
}
