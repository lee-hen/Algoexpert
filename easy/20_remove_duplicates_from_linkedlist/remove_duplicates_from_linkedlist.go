package remove_duplicates_from_linkedlist

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func (linkedList *LinkedList) addMany(values []int) {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
}

func (linkedList *LinkedList) getValues() []int {
	values := make([]int, 0)
	current := linkedList
	for current != nil {
		values = append(values, current.Value)
		current = current.Next
	}
	return values
}

func (linkedList *LinkedList) RemoveDuplicatesFromLinkedList() {
	current := linkedList

	for current != nil && current.Next != nil {
		if current.Value == current.Next.Value {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
}
