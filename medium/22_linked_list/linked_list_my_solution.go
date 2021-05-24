package linked_list

func (ll *DoublyLinkedList) setHead(node *Node) {
	if ll.Head == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		if !ll.isStandAlone(node) {
			node = ll.getRemovedNodeAndCutBindings(node)
		}

		current := ll.Head
		current.Prev = node
		node.Next = current

		ll.resetHead(node)
		ll.resetTail(node)
	}
}

func (ll *DoublyLinkedList) setTail(node *Node) {
	if ll.Tail == nil {
		ll.Head = node
		ll.Tail = node
	} else {
		if !ll.isStandAlone(node) {
			node = ll.getRemovedNodeAndCutBindings(node)
		}

		current := ll.Tail
		current.Next = node
		node.Prev = current

		ll.resetHead(node)
		ll.resetTail(node)
	}
}

func (ll *DoublyLinkedList) insertAfter(node, nodeToInsert *Node) {
	if node == nil {
		ll.SetTail(nodeToInsert)
		return
	}

	if !ll.isStandAlone(nodeToInsert) {
		nodeToInsert = ll.getRemovedNodeAndCutBindings(nodeToInsert)
	}

	current := node
	next := current.Next

	nodeToInsert.Prev = current
	current.Next = nodeToInsert

	if next != nil {
		nodeToInsert.Next = next
		next.Prev = nodeToInsert
	} else {
		ll.Tail = nodeToInsert
	}
}


func (ll *DoublyLinkedList) insertBefore(node, nodeToInsert *Node) {
	if node == nil {
		ll.SetHead(nodeToInsert)
		return
	}
	if !ll.isStandAlone(nodeToInsert) {
		nodeToInsert = ll.getRemovedNodeAndCutBindings(nodeToInsert)
	}

	current := node
	previous := current.Prev

	current.Prev = nodeToInsert
	nodeToInsert.Next = current

	if previous != nil {
		nodeToInsert.Prev = previous
		previous.Next = nodeToInsert
	} else {
		ll.Head = nodeToInsert
	}
}

func (ll *DoublyLinkedList) insertAtPosition(position int, nodeToInsert *Node) {
	node := ll.Head

	var idx = 1
	for node != nil {
		if idx == position {
			break
		}

		node = node.Next
		idx++
	}
	ll.InsertBefore(node, nodeToInsert)
}

func (ll *DoublyLinkedList) removeNodesWithValue(value int) {
	current := ll.Head
	for current != nil {
		if value == current.Value {
			ll.Remove(current)
		}
		current = current.Next
	}
}

func (ll *DoublyLinkedList) remove(node *Node) {
	current := node
	previous := current.Prev
	next := current.Next

	if previous != nil {
		previous.Next = next
	} else {
		ll.Head = current.Next
	}

	if next != nil {
		next.Prev = previous
	} else {
		ll.Tail = current.Prev
	}
}

func (ll *DoublyLinkedList) containsNodeWithValue(value int) bool {
	current := ll.Head
	for current != nil {
		if value == current.Value {
			return true
		}
		current = current.Next
	}
	return false
}


func (ll *DoublyLinkedList) isStandAlone (node *Node) bool {
	current := ll.Head

	for current != nil {
		if current == node {
			return false
		}
		current = current.Next
	}
	return true
}

func  (ll *DoublyLinkedList) getRemovedNodeAndCutBindings (node *Node) *Node{
	ll.Remove(node)
	node.Prev = nil
	node.Next = nil
	return node
}

func (ll *DoublyLinkedList) resetTail(node *Node) {
	current := node
	for current.Next != nil {
		current = current.Next
	}

	ll.Tail = current
}

func (ll *DoublyLinkedList) resetHead(node *Node) {
	current := node
	for current.Prev != nil {
		current = current.Prev
	}

	ll.Head = current
}
