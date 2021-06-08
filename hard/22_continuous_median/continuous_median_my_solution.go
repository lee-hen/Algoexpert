package continuous_median

func (handler *ContinuousMedianHandler) insert(number int) {
	bigNumbers := make([]int, 0)

	for len(handler.numbers) > 1 && handler.numbers[len(handler.numbers)-1] > number {
		bigger := handler.numbers[len(handler.numbers)-1]
		handler.numbers = handler.numbers[:len(handler.numbers)-1]
		bigNumbers = append(bigNumbers, bigger)
	}

	handler.numbers = append(handler.numbers, number)

	for len(bigNumbers) > 0 {
		bigger := bigNumbers[len(bigNumbers)-1]

		handler.numbers = append(handler.numbers, bigger)
		bigNumbers = bigNumbers[:len(bigNumbers)-1]
	}

	handler.updateMedian()
}

func (handler *ContinuousMedianHandler) updateMedian() {
	if len(handler.numbers) == 1 {
		handler.Median = float64(handler.numbers[0])
	} else if len(handler.numbers)%2 == 0 {
		middleIdx := len(handler.numbers) / 2
		handler.Median = float64(handler.numbers[middleIdx]+handler.numbers[middleIdx-1]) / 2.0
	} else {
		handler.Median = float64(handler.numbers[len(handler.numbers)/2])
	}
}
