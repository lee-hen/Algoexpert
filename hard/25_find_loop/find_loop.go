package find_loop
// important question

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// FindLoop
// better solution
func FindLoop(head *LinkedList) *LinkedList {
	first := head.Next
	second := head.Next.Next

	for first != second {
		first = first.Next

		second = second.Next.Next
	}

	first = head
	for first != second {
		first = first.Next
		second = second.Next
	}

	return first // or second
}

// my solution is the bad one
func findLoop1(head *LinkedList) *LinkedList {
	visited := make(map[*LinkedList]struct{})

	current := head
	for current != nil {
		if _, found := visited[current]; found {
			break
		}

		visited[current] = struct{}{}
		current = current.Next
	}

	return current
}
