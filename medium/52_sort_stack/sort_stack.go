package sort_stack
// important question

func SortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}

	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	SortStack(stack)
	insertInSortedOrder(&stack, top)

	return stack
}

func insertInSortedOrder(stack *[]int, value int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= value {
		*stack = append(*stack, value)
		return
	}

	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	insertInSortedOrder(stack, value)
	*stack = append(*stack, top)
}

func sortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}

	sortStackHelper(&stack)
	return stack
}

func sortStackHelper(stack *[]int) {
	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]

	if len(*stack) == 0 {
		*stack = append(*stack, top)
		return
	}

	sortStackHelper(stack)

	if (*stack)[len(*stack)-1] <= top {
		*stack = append(*stack, top)
	} else {
		value := top

		top = (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]

		*stack = append(*stack, value)
		*stack = append(*stack, top)
		sortStackHelper(stack)
	}
}
