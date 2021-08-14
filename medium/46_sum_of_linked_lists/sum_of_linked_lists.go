package sum_of_linked_lists
// important question

type LinkedList struct {
	Value int
	Next  *LinkedList
}

func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	carry := 0
	nodeOne := linkedListOne
	nodeTwo := linkedListTwo

	newLinkedListHeadPointer := &LinkedList{Value: 0}
	currentNode := newLinkedListHeadPointer
	for nodeOne != nil || nodeTwo != nil || carry != 0 {
		var valueOne, valueTwo int
		if nodeOne != nil {
			valueOne = nodeOne.Value
		}
		if nodeTwo != nil {
			valueTwo = nodeTwo.Value
		}

		sumOfValues := valueOne + valueTwo + carry
		newNode := &LinkedList{Value: sumOfValues % 10}

		currentNode.Next = newNode
		currentNode = newNode

		carry = sumOfValues / 10
		if nodeOne != nil {
			nodeOne = nodeOne.Next
		}
		if nodeTwo != nil {
			nodeTwo = nodeTwo.Next
		}
	}
	return newLinkedListHeadPointer.Next
}

func SumOfLinkedLists1(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	sliceOne := traverseLinkedList(linkedListOne)
	sliceTwo := traverseLinkedList(linkedListTwo)
	number := sliceToInt(ReverseSlice(sliceOne)) + sliceToInt(ReverseSlice(sliceTwo))

	reversedSlice := make([]int, 0)
	intToSlice(number, &reversedSlice)

	result := LinkedList{}
	result.insert(reversedSlice)

	return result.Next
}

func traverseLinkedList(l *LinkedList) []int {
	slice := make([]int, 0)
	curr := l
	for curr != nil {
		slice = append(slice, curr.Value)
		curr = curr.Next

	}
	return slice
}

func ReverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func intToSlice(n int, s *[]int) {
	if n != 0 {
		i := n % 10
		*s = append(*s, i) // reverse order output
		// *s = append([]int{i}, *s...) // origin order output
		intToSlice(n/10, s)
	}

	if len(*s) == 0 {
		*s = append(*s, 0)
	}
	return
}

func (linkedList *LinkedList) insert(values []int) *LinkedList {
	current := linkedList
	for current.Next != nil {
		current = current.Next
	}
	for _, value := range values {
		current.Next = &LinkedList{Value: value}
		current = current.Next
	}
	return linkedList
}
