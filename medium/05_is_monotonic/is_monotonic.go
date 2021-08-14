package is_monotonic
// important question

// IsMonotonic
// Better solution
func IsMonotonic(array []int) bool {
	if len(array) <= 2 {
		return true
	}
	direction := array[1] - array[0]
	for i := 2; i < len(array); i++ {
		if direction == 0 {
			direction = array[i] - array[i-1]
			continue
		}
		if breaksDirection(direction, array[i-1], array[i]) {
			return false
		}
	}
	return true
}
func breaksDirection(direction, previousInt, currentInt int) bool {
	difference := currentInt - previousInt
	if direction > 0 {
		return difference < 0
	}
	return difference > 0
}

func isMonotonic(array []int) bool {
	isIncreasing := false
	isDecreasing := false

	for i, arr := range array {
		if i > len(array) - 2 {
			break
		}
		if arr < array[i+1] {
			isIncreasing = true
		}

		if arr > array[i+1] {
			isDecreasing = true
		}
	}
	return !isIncreasing || !isDecreasing
}

func isMonotonic1(array []int) bool {
	for i, arr := range array {
		if i > len(array)-2 {
			break
		}
		if arr == array[i+1] {
			array = append(array[:i], array[i+1:]...)
		}
	}

	if len(array) < 3 {
		return true
	}

	for i,arr := range array {
		if i > len(array)-3 {
			break
		}
		if arr < array[i+1] && array[i+1] > array[i+2] ||
			arr > array[i+1] && array[i+1] < array[i+2] {
			return false
		}

	}
	return true
}
