package node_swap

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// 0 -> 1 -> 2 -> 3 -> 4 -> 5
// 1 -> 0 -> 2 -> 3
//           c    n

//           n    c
//           3 -> 2 -> 4
//                |
//                p

// 1 -> 0 -> 3 -> 2 -> 5 -> 4

// 0 -> 1 -> 2 -> 3 -> 4
// 1 -> 0 -> 3 -> 2 -> 4

func NodeSwap(head *LinkedList) *LinkedList {
	tempNode := &LinkedList{Value: 0}
	tempNode.Next = head

	prevNode := tempNode
	for prevNode.Next != nil && prevNode.Next.Next != nil {
		firstNode := prevNode.Next
		secondNode := prevNode.Next.Next
		// prevNode -> firstNode -> secondNode -> x

		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		prevNode.Next = secondNode
		// prevNode -> secondNode -> firstNode -> x

		prevNode = firstNode
	}

	return tempNode.Next
}

// my solution
func nodeSwap(head *LinkedList) *LinkedList {
	if head.Next == nil {
		return head
	}

	curr, next := head, head.Next
	head = next
	var prev *LinkedList
	for curr != nil && next != nil {
		curr.Next = next.Next
		next.Next = curr
		if prev != nil {
			prev.Next = next
		}

		prev = curr
		curr = curr.Next
		if curr != nil {
			next = curr.Next
		}
	}
	return head
}

// NodeSwap1
// O(n) time | O(n) space - where n is the number of nodes in the Linked List
func NodeSwap1(head *LinkedList) *LinkedList {
	return recursiveNodeSwap(head)
}

func recursiveNodeSwap(head *LinkedList) *LinkedList {
	if head == nil || head.Next == nil {
		return head
	}

	nextNode := head.Next
	head.Next = recursiveNodeSwap(head.Next.Next)
	nextNode.Next = head
	return nextNode
}
