package remove_kth_node_from_end

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// RemoveKthNodeFromEnd
// better solution
func RemoveKthNodeFromEnd(head *LinkedList, k int) {
	counter, first, second := 1, head, head
	for counter <= k {
		second = second.Next
		counter += 1
	}
	if second == nil {
		head.Value = head.Next.Value
		head.Next = head.Next.Next
		return
	}

	// second != nil is not corrected
	for second.Next != nil {
		second = second.Next
		first = first.Next
	}
	first.Next = first.Next.Next
}

// my solution
func removeKthNodeFromEnd(head *LinkedList, k int) {
	size := head.Size()
	if size <= 1 {
		return
	}

	current := head
	position := size - k
	if position == 0 {
		current.Value = current.Next.Value
		current.Next = current.Next.Next

		return
	}

	var idx = 1
	for position != 0 && current != nil && idx != position {
		current = current.Next
		idx++
	}

	removeNode := current.Next
	current.Next = removeNode.Next
	removeNode.Next = nil
}

func (ll *LinkedList) Size() int {
	current := ll
	var size int

	for current != nil {
		current = current.Next
		size++
	}
	return size
}
