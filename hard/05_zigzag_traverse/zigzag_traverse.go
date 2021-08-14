package zigzag_traverse
// important question

func ZigzagTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}

	height := len(array) - 1
	width := len(array[0]) - 1

	result := make([]int, 0)
	row := 0
	col := 0

	goDiagonallyDown := true
	for !isOutOfBounds(row, col, height, width) {
		result = append(result, array[row][col])
		if goDiagonallyDown {
			if col == 0 || row == height {
				goDiagonallyDown = false
				if row == height {
					col++
				} else {
					row++
				}
			} else {
				row++
				col--
			}
		} else {
			if row == 0 || col == width {
				goDiagonallyDown = true
				if col == width {
					row++
				} else {
					col++
				}
			} else {
				row--
				col++
			}
		}
	}
	return result
}

func isOutOfBounds(row, col, height, width int) bool {
	return row < 0 || row > height || col < 0 || col > width
}

//   0  1   2   3
//0  1  3   4   10
//1  2  5   9   11
//2  6  8   12  15
//3  7  13  14  16

type Direction int

const (
	Down Direction = 1 // down
	DiagonallyUpwards  Direction = 2 // up right
	Right Direction = 3 // right
	DiagonallyDownward  Direction = 4 // down left
)

// ZigzagTraverse1
// 長さと幅さは異なる時、
// 長さは幅より長い
// 幅さは長さより長い
// 長さは1の時
// 幅さは1の時
func ZigzagTraverse1(slice [][]int) []int {
	zigzag := make([]int, 0)
	zigzagTraverseHelper(DiagonallyDownward,0,0, &zigzag, slice)
	return zigzag
}

func zigzagTraverseHelper(direction Direction, row, col int, zigzag *[]int, slice [][]int) {
	if len(*zigzag) == len(slice) * len(slice[0]) {
		return
	}

	*zigzag = append(*zigzag, slice[row][col])

	if col == 0 {
		if direction == Down {
			// 1 * 5
			// index out of bounds
			if col == len(slice[0])-1 {
				row++
				direction = Down
			} else {
				row--
				col++
				direction = DiagonallyUpwards
			}
		} else if direction == DiagonallyDownward {
			// 6 * 5
			// index out of bounds
			if row == len(slice)-1 {
				col++
				direction = Right
			} else {
				row++
				direction = Down
			}
		}
	} else if col == len(slice[0]) - 1 {
		// 5 * 6
		if direction == Right {
			col--
			row++
			direction = DiagonallyDownward
		} else if direction == DiagonallyUpwards {
			row++
			direction = Down
		} else if direction == Down {
			row++
			col--
			direction = DiagonallyDownward
		}
	} else if row == 0 && col > 0 && col < len(slice[0]) - 1 {
		 if direction == DiagonallyUpwards {
			col++
			direction = Right
		} else if direction == Right {
			 // 5 * 1
			 // index out of bounds
			 if row == len(slice)-1 {
				 col++
				 direction = Right
			 } else {
				 col--
				 row++
				 direction = DiagonallyDownward
			 }
		}
	} else if row == len(slice)-1 && col > 0 && col < len(slice[0]) - 1 {
		if direction == DiagonallyDownward {
			col++
			direction = Right
		} else if direction == Right {
			row--
			col++
			direction = DiagonallyUpwards
		}
	} else {
		if direction == DiagonallyDownward {
			col--
			row++
			direction = DiagonallyDownward
		} else if direction == DiagonallyUpwards {
			row--
			col++
			direction = DiagonallyUpwards
		}
	}
	zigzagTraverseHelper(direction, row, col, zigzag, slice)
}
