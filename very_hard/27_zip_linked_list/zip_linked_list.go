package zip_linked_list

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// 1 -> 2 -> 3 -> 4 -> 5
//           s
//                     f

// 1 -> 2 -> 3
// 5 -> 4 -> 3

// 0    1    2    3    4    5
// 1 -> 2 -> 3 -> 4 -> 5 -> 6
//                s
//                               f

//         p1
// 1   2   3 -> 4
// | / |  /
// 6   5   4
//         p2

func ZipLinkedList(head *LinkedList) *LinkedList {
	slowNode := head
	fastNode := head
	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}
	secondHalfHead := slowNode
	//	secondHalfHead := slowNode.Next
	//	slowNode.Next = nil

	secondHalfNode := reverseLinkedList(secondHalfHead)
	firstHalfNode := head
	interweaveLinkedLists(firstHalfNode, secondHalfNode)
	return head
}

func reverseLinkedList(head *LinkedList) *LinkedList {
	var previousNode *LinkedList = nil
	var currentNode = head
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

func interweaveLinkedLists(p1, p2 *LinkedList) {
	n1, n2 := p1.Next, p2.Next
	for p1 != p2 && p1.Next != p2 {
		p1.Next = p2
		p2.Next = n1

		p1, p2 = n1, n2
		n1, n2 = p1.Next, p2.Next
	}

	//for p1 != nil && p2 != nil {
	//	p1Next := p1.Next
	//	p2Next := p2.Next
	//
	//	p1.Next = p2
	//	p2.Next = p1Next
	//
	//	p1 = p1Next
	//	p2 = p2Next
	//}
}
