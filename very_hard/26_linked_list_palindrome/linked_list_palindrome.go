package linked_list_palindrome

type LinkedList struct {
	Value int
	Next  *LinkedList
}

type LinkedListInfo struct {
	outerNodesAreEqual bool
	leftNodeToCompare  *LinkedList
}

// LinkedListPalindrome
// O(n) time | O(1) space - where n is the number of nodes in the Linked List
func LinkedListPalindrome(head *LinkedList) bool {
	slowNode := head
	fastNode := head
	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
	}

	reversedSecondHalfNode := reverseLinkedList(slowNode)
	firstHalfNode := head

	for reversedSecondHalfNode != nil {
		if reversedSecondHalfNode.Value != firstHalfNode.Value {
			return false
		}
		reversedSecondHalfNode = reversedSecondHalfNode.Next
		firstHalfNode = firstHalfNode.Next
	}

	return true
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

//  0 -> 1 -> 2    2 <- 1 -> 0 -> 3
//                 c    n    t

//  0 -> 1    3 <- 1 <- 0
//            c    n    t

//  0 -> 1 -> 2    2 <- 1 <- 0
//  c                        t
// 同じく O(n) time | O(1) space
func linkedListPalindrome(head *LinkedList) bool {
	if head.Next == nil {
		return true
	}

	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Value == curr.Next.Value || curr.Next.Next != nil && curr.Value == curr.Next.Next.Value {
			temp := curr
			curr = curr.Next
			temp.Next = nil
			break
		}
		curr = curr.Next
	}

	if curr != nil {
		tempCurr := curr
		next := curr.Next
		var temp *LinkedList

		for next != nil {
			temp = next.Next
			next.Next = curr


			curr = next
			next = temp
		}
		tempCurr.Next = nil
	}

	tail := curr
	curr = head

	for curr != nil && tail != nil {
		if curr.Value == tail.Value {
			curr = curr.Next
			tail = tail.Next
		} else {
			break
		}
	}
	if curr == nil && tail == nil {
		return true
	}

	if curr == nil && tail != nil {
		return true
	}
	return false
}

// LinkedListPalindrome1
// O(n) time | O(n) space - where n is the number of nodes in the Linked List
func LinkedListPalindrome1(head *LinkedList) bool {
	isPalindromeResults := isPalindrome(head, head)
	return isPalindromeResults.outerNodesAreEqual
}

func isPalindrome(leftNode *LinkedList, rightNode *LinkedList) LinkedListInfo {
	if rightNode == nil {
		return LinkedListInfo{true, leftNode}
	}

	recursiveCallResults := isPalindrome(leftNode, rightNode.Next)
	leftNodeToCompare := recursiveCallResults.leftNodeToCompare
	outerNodesAreEqual := recursiveCallResults.outerNodesAreEqual

	recursiveIsEqual := outerNodesAreEqual && leftNodeToCompare.Value == rightNode.Value
	nextLeftNodeToCompare := leftNodeToCompare.Next

	return LinkedListInfo{recursiveIsEqual, nextLeftNodeToCompare}
}
