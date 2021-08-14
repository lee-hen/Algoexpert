package waterfall_streams
// important question

// WaterfallStreams
// O(w^2 * h) time | O(w) space - where w and h
// are the width and height of the input array
func WaterfallStreams(array [][]float64, source int) []float64 {
	rowAbove := array[0]
	// We'll use -1 to represent water, since 1 is used for a block.
	rowAbove[source] = -1

	for row := 1; row < len(array); row++ {
		currentRow := array[row]
		for idx := range rowAbove {
			valueAbove := rowAbove[idx]

			hasWaterAbove := valueAbove < 0
			hasBlock := currentRow[idx] == 1

			if !hasWaterAbove {
				continue
			}

			if !hasBlock {
				// If there is no block in the current column, move the water down.
				currentRow[idx] += valueAbove
				continue
			}

			splitWater := valueAbove / 2

			// Move water right.
			var rightIdx = idx
			for rightIdx+1 < len(rowAbove) {
				rightIdx += 1
				if rowAbove[rightIdx] == 1.0 {
					break // if there is a block in the way
				}

				if currentRow[rightIdx] != 1.0 { // if there is no block below us
					currentRow[rightIdx] += splitWater
					break
				}
			}

			// Move water left.
			var leftIdx = idx
			for leftIdx-1 >= 0 {
				leftIdx -= 1
				if rowAbove[leftIdx] == 1.0 {
					break // if there is a block in the way
				}

				if currentRow[leftIdx] != 1.0 { // if there is no block below us
					currentRow[leftIdx] += splitWater
					break
				}
			}

		}

		rowAbove = currentRow
	}

	finalPercentages := make([]float64, 0, len(rowAbove))
	for _, num := range rowAbove {
		if num == 0 {
			finalPercentages = append(finalPercentages, num)
		} else {
			finalPercentages = append(finalPercentages, num*-100)
		}
	}
	return finalPercentages
}

// source =  3
// [0, 0, 0, 0, 0, 0, 0],
// [1, 0, 0, 0, 0, 0, 0],
// [0, 0, 1, 1, 1, 0, 0],
// [0, 0, 0, 0, 0, 0, 0],
// [1, 1, 1, 0, 0, 1, 0],
// [0, 0, 0, 0, 0, 0, 1],
// [0, 0, 0, 0, 0, 0, 0],

// answer [0, 0, 0, 25, 25, 0, 0]

// The water will flow as follows:
// [0, 0, 0, ., 0, 0, 0],
// [1, ., ., ., ., ., 0],
// [0, ., 1, 1, 1, ., 0],
// [., ., ., ., ., ., .],
// [1, 1, 1, ., ., 1, 0],
// [0, 0, 0, ., ., 0, 1],
// [0, 0, 0, ., ., 0, 0],

// WaterfallStreams
// my solution too complicated
func waterfallStreams(array [][]float64, source int) []float64 {
	array[0][source] = 100.0
	waterfallStreamsHelper(1, source, array)

	return array[len(array)-1]
}

func waterfallStreamsHelper(i int, j int, array [][]float64) {
	for i < len(array) {
		if stuckWaterfallStreams(i, j, array) {
			return
		}

		if array[i][j] == 1.0 {
			break
		}
		array[i][j] = array[i-1][j]
		i++
	}

	if i > len(array)-1 {
		return
	}

	fallLeft, fallRight := true, true
	left, right := j-1, j+1
	if left < 0 || array[i-1][left] == 1.0 {
		fallLeft = false
	}

	if right > len(array[i])-1 || array[i-1][right] == 1.0 {
		fallRight = false
	}

	for fallLeft && left >= 0 {
		if array[i][left] != 1.0 {
			break
		}
		left--
	}

	for fallRight && right < len(array[i]) {
		if array[i][right] != 1.0 {
			break
		}
		right++
	}

	if j > len(array[i])/2 {
		left, right = right, left
		fallLeft, fallRight = fallRight, fallLeft
	}

	if fallLeft && left >= 0 && left < len(array[i]) {
		if array[i-1][left] != 0.0 {
			array[i][left] = array[i-1][left]
		} else if i < len(array)-1 && array[i+1][left] != 1.0 {
			array[i][left] = 0.0
		}

		array[i][left] += array[i-1][j]/2
		if !stuckWaterfallStreams(i, left, array) {
			waterfallStreamsHelper(i+1, left, array)
		}
	}

	if fallRight && right >= 0 && right < len(array[i]) {
		if array[i-1][right] != 0.0 {
			array[i][right] = array[i-1][right]
		} else if i < len(array)-1 && array[i+1][right] != 1.0 {
			array[i][right] = 0.0
		}

		array[i][right] += array[i-1][j]/2
		if !stuckWaterfallStreams(i, right, array) {
			waterfallStreamsHelper(i+1, right, array)
		}
	}

	return
}

func stuckWaterfallStreams(i, j int, array [][]float64) bool {
	left, right := j-1, j+1

	if i < len(array)-1 {
		if left >= 0 && right <= len(array[i])-1 {
			if array[i][left] == 1.0 && array[i][right] == 1.0 && array[i+1][j] == 1.0 {
				return true
			}
		}
		if right <= len(array[i])-1 && left < 0 {
			if array[i+1][j] == 1.0 && array[i][right] == 1.0 {
				return true
			}
		}

		if right > len(array[i])-1 && left >= 0 {
			if array[i+1][j] == 1.0 && array[i][left] == 1.0 {
				return true
			}
		}
	}

	return false
}

