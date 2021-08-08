package reverse_alternating_k_nodes

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// ReverseAlternatingKNodes
// O(n) time | O(1) space - where n is the number of nodes in the Linked List
func ReverseAlternatingKNodes(head *LinkedList, k int) *LinkedList {
	var finalHead *LinkedList

	var previousNode *LinkedList
	currentNode := head

	for currentNode != nil {
		reversedGroupHead, nextNode := reverseKNodes(currentNode, k)
		// The `currentNode` is now the tail of the reversed
		// group, so we make it point to the `nextNode`.
		currentNode.Next = nextNode
		currentNode = nextNode

		if previousNode == nil {
			finalHead = reversedGroupHead
		} else {
			previousNode.Next = reversedGroupHead
		}

		var skippedNodesCount = 0
		for currentNode != nil && skippedNodesCount < k {
			previousNode = currentNode
			currentNode = currentNode.Next

			skippedNodesCount += 1
		}
	}

	return finalHead
}

func reverseKNodes(head *LinkedList, k int) (*LinkedList, *LinkedList) {
	reversedNodesCount := 0
	var previousNode *LinkedList
	var currentNode = head
	for currentNode != nil && reversedNodesCount < k {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode

		reversedNodesCount++
	}
	return previousNode, currentNode
}

// ReverseAlternatingKNodes1
// O(n) time | O(1) space - where n is the number of nodes in the Linked List
func ReverseAlternatingKNodes1(head *LinkedList, k int) *LinkedList {
	var finalHead *LinkedList

	isGroupToReverse := true
	runningK := 1

	var previousGroupTail *LinkedList = nil
	var currentGroupHead = head
	var currentNode = head

	for currentNode != nil {
		shouldReverse := isGroupToReverse && (runningK == k || currentNode.Next == nil)

		if !shouldReverse {
			if runningK == k {
				runningK = 1
				isGroupToReverse = true

				previousGroupTail = currentNode
				currentGroupHead = currentNode.Next
			} else {
				runningK += 1
			}

			currentNode = currentNode.Next
			continue
		}

		runningK = 1
		isGroupToReverse = false

		nextNode := currentNode.Next
		currentNode.Next = nil
		reversedGroupHead := reverseLinkedList(currentGroupHead)
		reversedGroupTail := currentGroupHead
		reversedGroupTail.Next = nextNode

		if previousGroupTail == nil {
			finalHead = reversedGroupHead
		} else {
			previousGroupTail.Next = reversedGroupHead
		}

		currentNode = nextNode
		currentGroupHead = nextNode
		previousGroupTail = reversedGroupTail
	}

	return finalHead
}

// my solution well can change to O(1) space

//        0    1    2    3    4    5    6                                         -> 15 - 16 - 17 - 18
// head = 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11 -> 12 -> 13 -> 14  // the head node with value 1
// k = 3

// pariOfNodes
// [1, 3] [6 9] [12 14(15)]

// 3 -> 2 -> 1 -> 4 -> 5 -> 6 -> 9 -> 8 -> 7 -> 10 -> 11 -> 12 -> 14 -> 13 // the new head node with value 3

func reverseAlternatingKNodes(head *LinkedList, k int) *LinkedList {
	pairOfNodes := make([][]*LinkedList, 1)
	pairOfNodes[0] = make([]*LinkedList, 2, 2)
	pairOfNodes[0][0] = head

	curr := head
	var lastNode *LinkedList
	var idx, prevIdx int
	for curr != nil {
		if idx == k-1 {
			pairOfNodes[0][1] = curr
			prevIdx = idx
		}

		if idx == prevIdx + k  {
			pairOfNode := make([]*LinkedList, 2, 2)
			pairOfNode[0] = curr
			pairOfNodes = append(pairOfNodes, pairOfNode)
		}

		if idx == prevIdx + 2*k {
			pairOfNodes[len(pairOfNodes)-1][1]= curr
			prevIdx = idx
		}

		idx++
		lastNode = curr
		curr = curr.Next
	}

	if pairOfNodes[len(pairOfNodes)-1][1] == nil {
		if pairOfNodes[len(pairOfNodes)-1][0] != lastNode {
			pairOfNodes[len(pairOfNodes)-1][1] = lastNode
		} else {
			pairOfNodes = pairOfNodes[:len(pairOfNodes)-1]
		}
	}

	return reversePairOfNodes(pairOfNodes)
}

func reversePairOfNodes(pairOfNodes [][]*LinkedList) *LinkedList {
	for i, pairOfNode := range pairOfNodes {
		var nextNode *LinkedList
		nextNode = pairOfNode[1].Next
		pairOfNode[1].Next = nil
		if i == 0 {
			reverseLinkedList(pairOfNode[0])
		} else {
			tail := pairOfNode[0].Next
			tempNext := nextNode
			nextNode = pairOfNode[1]
			reverseLinkedList(pairOfNode[0].Next)
			tail.Next = tempNext
		}
		pairOfNode[0].Next = nextNode
	}

	return pairOfNodes[0][1]
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
