package minimum_area_rectangle
// important question

import (
	"fmt"
	"math"
	"sort"
)

// MinimumAreaRectangle
// O(n^2) time | O(n) space - where n is the number of points
func MinimumAreaRectangle(points [][]int) int {
	pointSet := createPointSet(points)
	var minimumAreaFound = math.MaxInt32

	for currentIdx := range points {
		p2x, p2y := points[currentIdx][0], points[currentIdx][1]
		for previousIdx := 0; previousIdx < currentIdx; previousIdx++ {
			p1x, p1y := points[previousIdx][0], points[previousIdx][1]
			pointsShareValue := p1x == p2x || p1y == p2y

			if pointsShareValue {
				continue
			}

			// If (p1x, p2y) and (p2x, p1y), exist we've found a rectangle.
			point1OnOppositeDiagonalExists := pointSet[convertPointToString(p1x, p2y)]
			point2OnOppositeDiagonalExists := pointSet[convertPointToString(p2x, p1y)]
			oppositeDiagonalExists := point1OnOppositeDiagonalExists && point2OnOppositeDiagonalExists
			if oppositeDiagonalExists {
				currentArea := abs(p2x-p1x) * abs(p2y-p1y)
				minimumAreaFound = min(minimumAreaFound, currentArea)
			}
		}
	}
	if minimumAreaFound != math.MaxInt32 {
		return minimumAreaFound
	}
	return 0
}

func createPointSet(points [][]int) map[string]bool {
	pointSet := map[string]bool{}
	for _, point := range points {
		x, y := point[0], point[1]
		pointString := convertPointToString(x, y)
		pointSet[pointString] = true
	}
	return pointSet
}

func convertPointToString(x, y int) string {
	return fmt.Sprintf("%d:%d", x, y)
}


// MinimumAreaRectangle1
// O(n^2) time | O(n) space - where n is the number of points
func MinimumAreaRectangle1(points [][]int) int {
	columns := initializeColumns(points)
	var minimumAreaFound = math.MaxInt32
	edgesParallelToYAxis := map[string]int{}

	sortedColumns := make([]int, 0)
	for k := range columns {
		sortedColumns = append(sortedColumns, k)
	}
	sort.Ints(sortedColumns)

	for _, x := range sortedColumns {
		yValuesInCurrentColumn := columns[x]
		sort.Ints(yValuesInCurrentColumn)

		for currentIdx := range yValuesInCurrentColumn {
			y2 := yValuesInCurrentColumn[currentIdx]

			for previousIdx := 0; previousIdx < currentIdx; previousIdx++ {
				y1 := yValuesInCurrentColumn[previousIdx]
				pointString := fmt.Sprintf("%d:%d", y1, y2)

				if _, found := edgesParallelToYAxis[pointString]; found {
					currentArea := (x - edgesParallelToYAxis[pointString]) * (y2 - y1)
					minimumAreaFound = min(minimumAreaFound, currentArea)
				}

				edgesParallelToYAxis[pointString] = x
			}
		}
	}

	if minimumAreaFound != math.MaxInt32 {
		return minimumAreaFound
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func initializeColumns(points [][]int) map[int][]int {
	columns := map[int][]int{}
	for _, point := range points {
		x, y := point[0], point[1]
		columns[x] = append(columns[x], y)
	}

	return columns
}

// minimumAreaRectangle
// my solution
func minimumAreaRectangle(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	parallelPoints := make([][][]int, 0)
	var i int
	for i < len(points) {
		j := i + 1
		for j < len(points) {
			if points[i][0] != points[j][0] {
				break
			}
			j++
		}

		if len(points[i:j]) != 1{
			parallelPoints = append(parallelPoints, points[i:j])
		}

		i = j
	}

	if len(parallelPoints) == 0 {
		return 0
	}

	sort.Slice(parallelPoints[0], func(i, j int) bool {
		return parallelPoints[0][i][1] < parallelPoints[0][j][1]
	})

	var j = 1
	var areas []int
	for j < len(parallelPoints) {
		sort.Slice(parallelPoints[j], func(m, n int) bool {
			return parallelPoints[j][m][1] < parallelPoints[j][n][1]
		})

		area := calculateRectangleArea(parallelPoints[j-1], parallelPoints[j])
		areas = append(areas, area)
		j++
	}

	return minOfSlice(areas)
}

func calculateRectangleArea(mPoints, nPoints [][]int) int {
	if len(mPoints) > len(nPoints) {
		mPoints, nPoints = nPoints, mPoints
	}

	var i, j int
	rectanglePoints := make([][]int, 0)

	for i < len(mPoints) {
		y := mPoints[i][1]

		var k = j
		for k < len(nPoints) {
			if y == nPoints[k][1] {
				j = k+1
				rectanglePoints = append(rectanglePoints, nPoints[k])
				rectanglePoints = append(rectanglePoints, mPoints[i])
				break
			}
			k++
		}

		i++
	}

	return areaOfRectangle(rectanglePoints)
}

func areaOfRectangle(points [][]int) int {
	var i int
	minArea := math.MaxInt32
	for i+3 < len(points) {
		width := points[i+1][0] - points[i+2][0]
		height := points[i+1][1] - points[i+2][1]

		area := abs(height) * abs(width)
		if area < minArea {
			minArea = area
		}
		i+=2
	}
	return minArea
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func minOfSlice(values []int) int {
	min := math.MaxInt32
	for _, value := range values {
		if value < min {
			min = value
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}
