package merge_linked_lists

type LinkedList struct {
	Value int
	Next  *LinkedList
}

//                      p1    n
// 1  3 -> 4 -> 5      9 -> 10
// v /            \   /
// 2                 6    7 -> 8
//                        p2

// MergeLinkedLists
//  O(n + m) time | O(1) space
// simple loop (my) solution is better
func MergeLinkedLists(headOne *LinkedList, headTwo *LinkedList) *LinkedList {
	var mergedList *LinkedList

	ptrOne, ptrTwo := headOne, headTwo
	if headOne.Value > headTwo.Value {
		ptrOne, ptrTwo = headTwo, headOne
	}
	mergedList = ptrOne

	next := ptrOne.Next
	for next != nil {
		if ptrTwo == nil {
			break
		}

		if ptrTwo.Value < next.Value {
			ptrOne.Next = ptrTwo
			ptrTwo = ptrTwo.Next

			ptrOne = ptrOne.Next
			ptrOne.Next = next
		} else {
			next = next.Next
			ptrOne = ptrOne.Next
		}
	}

	if ptrTwo != nil {
		ptrOne.Next = ptrTwo
	}

	// recursiveMerge(ptrOne, ptrTwo, nil)
	return mergedList
}

// O(n + m) time | O(n + m) space - where n is the number of nodes in the first
// Linked List and m is the number of nodes in the second Linked List
func recursiveMerge(p1, p2, p1Prev *LinkedList) {
	if p1 == nil {
		p1Prev.Next = p2
		return
	}

	if p2 == nil {
		return
	}

	if p1.Value < p2.Value {
		recursiveMerge(p1.Next, p2, p1)
		return
	}

	if p1Prev != nil {
		p1Prev.Next = p2
	}

	newP2 := p2.Next
	p2.Next = p1

	recursiveMerge(p1, newP2, p2)
}
