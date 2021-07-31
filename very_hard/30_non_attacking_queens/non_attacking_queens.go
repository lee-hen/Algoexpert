package non_attacking_queens

// NonAttackingQueens
// Upper Bound: O(n!) time | O(n) space - where n is the input number
func NonAttackingQueens(n int) int {
	blockedColumns := map[int]bool{}
	blockedUpDiagonals := map[int]bool{}
	blockedDownDiagonals := map[int]bool{}
	return getNumberOfNonAttackingQueenPlacements(0, blockedColumns,
		blockedUpDiagonals, blockedDownDiagonals, n)
}

func getNumberOfNonAttackingQueenPlacements(
	row int,
	blockedColumns map[int]bool,
	blockedUpDiagonals map[int]bool,
	blockedDownDiagonals map[int]bool,
	boardSize int,
) int {
	if row == boardSize {
		return 1
	}

	validPlacements := 0
	for col := 0; col < boardSize; col++ {
		if IsNonAttackingPlacement(row, col, blockedColumns, blockedUpDiagonals, blockedDownDiagonals) {
			placeQueen(row, col, blockedColumns, blockedUpDiagonals, blockedDownDiagonals)
			validPlacements += getNumberOfNonAttackingQueenPlacements(row+1,
				blockedColumns, blockedUpDiagonals, blockedDownDiagonals, boardSize)
			removeQueen(row, col, blockedColumns, blockedUpDiagonals, blockedDownDiagonals)
		}
	}
	return validPlacements
}

// IsNonAttackingPlacement
// This is always an O(1)-time operation.
func IsNonAttackingPlacement(row, col int,
	blockedColumns map[int]bool,
	blockedUpDiagonals map[int]bool,
	blockedDownDiagonals map[int]bool,
) bool {
	if blockedColumns[col] {
		return false
	}
	if blockedUpDiagonals[row+col] {
		return false
	}
	if blockedDownDiagonals[row-col] {
		return false
	}
	return true
}

func placeQueen(row, col int,
	blockedColumns map[int]bool,
	blockedUpDiagonals map[int]bool,
	blockedDownDiagonals map[int]bool,
) {
	blockedColumns[col] = true
	blockedUpDiagonals[row+col] = true
	blockedDownDiagonals[row-col] = true
}

func removeQueen(row, col int,
	blockedColumns map[int]bool,
	blockedUpDiagonals map[int]bool,
	blockedDownDiagonals map[int]bool,
) {
	blockedColumns[col] = false
	blockedUpDiagonals[row+col] = false
	blockedDownDiagonals[row-col] = false
}

// NonAttackingQueens1
// Lower Bound: O(n!) time | O(n) space - where n is the input number
func NonAttackingQueens1(n int) int {
	// Each index of `columnPlacements` represents a row of the chessboard,
	// and the value at each index is the column (on the relevant row) where
	// a queen is currently placed.
	columnPlacements := make([]int, n)
	return getNumberOfNonAttackingQueenPlacements1(0, columnPlacements, n)
}

func getNumberOfNonAttackingQueenPlacements1(
	row int, columnPlacements []int, boardSize int) int {
	if row == boardSize {
		return 1
	}

	validPlacements := 0
	for col := 0; col < boardSize; col++ {
		if isNonAttackingPlacement(row, col, columnPlacements) {
			columnPlacements[row] = col
			validPlacements += getNumberOfNonAttackingQueenPlacements1(row+1,
				columnPlacements, boardSize)
		}
	}
	return validPlacements
}

// As `row` tends to `n`, this becomes an O(n)-time operation.
func isNonAttackingPlacement(row, col int, columnPlacements []int) bool {
	for previousRow := 0; previousRow < row; previousRow++ {
		columnToCheck := columnPlacements[previousRow]
		sameColumn := columnToCheck == col
		onDiagonal := abs(columnToCheck-col) == row-previousRow
		if sameColumn || onDiagonal {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type Direction int

const (
	None Direction = iota - 1
	Up
	LeftUp
	RightUp
)

// nonAttackingQueens
// my solution.
func nonAttackingQueens(n int) int {
	placements := make([]int, n, n)
	for i := range placements {
		placements[i] = -1
	}

	return tryPlaceAttackingQueens(0, n, placements)
}

func tryPlaceAttackingQueens(i, n int, placements []int) int {
	if i == n {
		return 1
	}

	var result int
	for j := 0; j < n; j++ {
		if validToPlaceQueen(i, j, n, placements) {
			placements[i] = j
			result += tryPlaceAttackingQueens(i+1, n, placements)
		}
	}

	return result
}

func validToPlaceQueen(i, j, n int, placements []int) bool {
	// check row up, weather could place Queens in the current position.
	return validWithDirection(i, j, n, Up, placements) &&
		validWithDirection(i, j, n, LeftUp, placements) &&
		validWithDirection(i, j, n, RightUp, placements)
}

func validWithDirection(i, j, n int, direction Direction, placements []int) bool {
	next := nextPosition(direction, i, j)

	for next[0] >= 0 && next[0] < n && next[1] >= 0 && next[1] < n {
		if placements[next[0]] == next[1] {
			return false
		}
		next = nextPosition(direction, next[0], next[1])
	}

	return true
}

func nextPosition(direction Direction, i, j int) []int {
	if direction == LeftUp {
		i, j = i-1, j-1
	} else if direction == RightUp {
		i, j = i-1, j+1
	} else if direction == Up {
		i = i - 1
	}
	return []int{i, j}
}
