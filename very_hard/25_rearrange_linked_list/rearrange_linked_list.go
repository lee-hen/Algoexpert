package rearrange_linked_list
// important question

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// RearrangeLinkedList
// O(n) time | O(1) space - where n is the number of nodes in the Linked List
// This approach just like dynamic programming.
func RearrangeLinkedList(head *LinkedList, k int) *LinkedList {
	var smallerListHead, smallerListTail *LinkedList
	var equalListHead, equalListTail *LinkedList
	var greaterListHead, greaterListTail *LinkedList

	curr := head
	for curr != nil {
		if curr.Value < k {
			// previous smaller node will point to the curr node
			smallerListHead, smallerListTail = growLinkedList(smallerListHead, smallerListTail, curr)
		} else if curr.Value > k {
			// previous larger node will point to the curr node
			greaterListHead, greaterListTail = growLinkedList(greaterListHead, greaterListTail, curr)
		} else {
			// previous equal node will point to the curr node
			equalListHead, equalListTail = growLinkedList(equalListHead, equalListTail, curr)
		}

		prevNode := curr
		curr = curr.Next
		prevNode.Next = nil
	}

	firstHead, firstTail := connectLinkedLists(smallerListHead, smallerListTail, equalListHead, equalListTail)
	finalHead, _ := connectLinkedLists(firstHead, firstTail, greaterListHead, greaterListTail)
	return finalHead
}

// does not change the node.next
// ex.
// find the smaller node when newHead (smallerListHead) is nil then change the newHead point to the node
// and change newTail (smallerListTail) point the node and return -> caller will replace smallerListHead and smallerListTail and new's ones
// next time change the tail's (previous node) next pointer to the new node(curr)
// so the loop doesn't change the curr.next pointer so this implementation does require space only O(1)
func growLinkedList(head, tail, node *LinkedList) (*LinkedList, *LinkedList) {
	newHead, newTail := head, node
	if newHead == nil {
		newHead = node
	}
	if tail != nil {
		tail.Next = node
	}
	return newHead, newTail
}

func connectLinkedLists(headOne, tailOne, headTwo, tailTwo *LinkedList) (*LinkedList, *LinkedList) {
	newHead, newTail := headOne, tailTwo
	if newHead == nil {
		newHead = headTwo
	}
	if newTail == nil {
		newTail = tailOne
	}

	if tailOne != nil {
		tailOne.Next = headTwo
	}

	return newHead, newTail
}

// k=3
// 3 -> 0 -> 5 -> 2 -> 1 -> 4
// 0 -> 2 -> 1 -> 3 -> 5 -> 4


// k=3
// 0 -> 5 -> 2 -> 1 -> 4
// 0 -> 2 -> 1 -> 5 -> 4

func rearrangeLinkedList(head *LinkedList, k int) *LinkedList {
	smallerNodes := make([]*LinkedList, 0)
	largerNodes := make([]*LinkedList, 0)
	kthNodes := make([]*LinkedList, 0)

	curr := head
	for curr != nil {
		if curr.Value < k {
			smallerNodes = append(smallerNodes, curr)
		} else if curr.Value > k {
			largerNodes = append(largerNodes, curr)
		} else {
			kthNodes = append(kthNodes, curr)
		}
		curr = curr.Next
	}

	var startIdx int
	if len(smallerNodes) > 0 {
		head = smallerNodes[0]
		if len(kthNodes) != 0 {
			if head.Value > kthNodes[0].Value {
				head = kthNodes[0]
			} else {
				smallerNodes = append(smallerNodes, kthNodes...)
			}
		}

		if head == smallerNodes[0] {
			curr = head
			startIdx = 1
		} else {
			curr = kthNodes[len(kthNodes)-1]
			curr.Next = smallerNodes[0]
		}

		for i := startIdx; i < len(smallerNodes); i++ {
			node := smallerNodes[i]
			curr.Next = node
			curr = curr.Next
		}
		startIdx = 0
	} else if len(largerNodes) > 0 {
		head = largerNodes[0]
		if len(kthNodes) != 0 {
			head = kthNodes[0]
			curr = kthNodes[len(kthNodes)-1]
			curr.Next = largerNodes[0]
			startIdx = 0
		} else {
			curr = head
			startIdx = 1
		}
	}

	if curr != nil {
		for i := startIdx; i < len(largerNodes); i++ {
			node := largerNodes[i]
			curr.Next = node
			curr = curr.Next
		}

		curr.Next = nil
	}

	return head
}
