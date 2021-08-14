package move_element_to_end

// MoveElementToEnd
// better solution
func MoveElementToEnd(array []int, toMove int) []int {
	i, j := 0, len(array)-1
	for i < j {
		for i < j && array[j] == toMove {
			j--
		}
		if array[i] == toMove {
			array[i], array[j] = array[j], array[i]
		}
		i++
	}
	return array
}

func moveElementToEnd(array []int, toMove int) []int {
	movableIdx := - 1
	for i := range array {
		for j := len(array)-1; j > i; j-- {
			if toMove != array[j] {
				movableIdx = j
				break
			}
		}
		if movableIdx <= i {
			break
		}
		if array[i] == toMove {
			array[i], array[movableIdx] = array[movableIdx], array[i]
		}
	}
	return array
}
