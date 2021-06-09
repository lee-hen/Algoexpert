package reverse_linked_list

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// ReverseLinkedList
// my solution is the better
func ReverseLinkedList(head *LinkedList) *LinkedList {
	prev := head
	current := prev.Next
	prev.Next = nil

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
