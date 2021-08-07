package inverted_bisection

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// InvertedBisection
// O(n) time | O(1) space - where n is the number of nodes in the Linked List
func InvertedBisection(head *LinkedList) *LinkedList {
	length := getLinkedListLength(head)
	if length <= 3 {
		return head
	}

	firstHalfTail := getMiddleNode(head, length)
	middleNode := firstHalfTail.Next
	secondHalfHead := middleNode.Next
	if length%2 == 0 {
		middleNode = nil
		secondHalfHead = firstHalfTail.Next
	}

	firstHalfTail.Next = nil
	reverseLinkedList(head)
	reversedSecondHalfHead := reverseLinkedList(secondHalfHead)

	if middleNode == nil {
		head.Next = reversedSecondHalfHead
	} else {
		head.Next = middleNode
		middleNode.Next = reversedSecondHalfHead
	}
	return firstHalfTail
}

func getLinkedListLength(head *LinkedList) int {
	length := 0
	currentNode := head
	for currentNode != nil {
		currentNode = currentNode.Next
		length++
	}
	return length
}

func getMiddleNode(head *LinkedList, length int) *LinkedList {
	halfLength := length / 2
	currentPosition := 1
	currentNode := head
	for currentPosition != halfLength {
		currentNode = currentNode.Next
		currentPosition++
	}
	return currentNode
}



// 0 -> 1 -> 2 -> 3 -> 4 -> 5
// 2 -> 1 -> 0 -> 5 -> 4 -> 3

// 0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6
// 2 -> 1 -> 0 -> 3 -> 6 -> 5 -> 4
// my solution
func invertedBisection(head *LinkedList) *LinkedList {
	slowNode := head
	fastNode := head

	var lastFirstHalfNode, secondHalfHeadNode  *LinkedList
	for fastNode != nil && fastNode.Next != nil {
		lastFirstHalfNode = slowNode
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	var middleNode *LinkedList
	if lastFirstHalfNode != nil {
		if fastNode == nil {
			secondHalfHeadNode = slowNode
		} else {
			middleNode = lastFirstHalfNode.Next
			secondHalfHeadNode = slowNode.Next
			slowNode.Next = nil
		}
		lastFirstHalfNode.Next = nil
	}

	secondHalfNode := reverseLinkedList(secondHalfHeadNode)
	firstHalfNode := reverseLinkedList(head)
	if middleNode != nil {
		head.Next = middleNode
		middleNode.Next = secondHalfNode
	} else {
		head.Next = secondHalfNode
	}
	return firstHalfNode
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
