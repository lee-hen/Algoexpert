package shift_linked_list

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// 0 -> 1 -> 2 -> 3 -> 4 -> 5
// k = 2
// 4 -> 5 -> 0 -> 1 -> 2 -> 3

// ShiftLinkedList
// my solution is the better
func ShiftLinkedList(head *LinkedList, k int) *LinkedList {
	len := head.Length()
	k = k % len
	if  k < 0 {
		k += len
	}

	current := head
	tail := current.shift(k)

	for tail.Next != nil {
		current = current.Next
		tail = tail.Next
	}

	tail.Next = head
	head = current.Next
	current.Next = nil

	return head
}

func (ll *LinkedList) shift(k int) *LinkedList{
	current := ll
	var counter int
	for counter < k {
		counter++
		current = current.Next
	}

	return current
}

func (ll *LinkedList) Length() int{
	current := ll
	var counter int
	for current != nil {
		counter++
		current = current.Next
	}
	return counter
}

