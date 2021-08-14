package get_nth_fib
// important question

func GetNthFib(n int) int {
	return help(n, map[int]int{1: 0, 2: 1})
}

func help(n int, memoize map[int]int) int {
	if val, ok := memoize[n]; ok {
		return val
	}

	memoize[n] = help(n-2, memoize) + help(n-1, memoize)
	return memoize[n]
}

func getNthFib3(n int) int {
	lastTwo := []int{0, 0}
	var nextFib int

	for i := 1; i <= n; i++ {
		if i == 2 {
			lastTwo[1] = 1
		}
		lastTwo = []int{lastTwo[1], nextFib}
		nextFib = lastTwo[1] + lastTwo[0]
	}

	return nextFib
}

// My solution
func getNthFib1(n int) int {
	if n == 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	return getNthFib1(n-1) + getNthFib1(n-2)
}

func getNthFib2(n int) int {
	var prevFirst int
	var prevLast int
	var thisValue int

	for i := 1; i <= n; i++ {
		if i == 2 {
			prevLast = 1
		}

		prevFirst = prevLast
		prevLast = thisValue

		thisValue = prevFirst + prevLast
	}

	return thisValue
}
