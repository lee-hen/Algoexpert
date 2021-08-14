package line_through_points
// important question

import (
	"fmt"
	"math"
)

// LineThroughPoints
// O(n^2) time | O(n) space - where n is the number of points
func LineThroughPoints(points [][]int) int {
	maxNumberOfPointsOnLine := 1

	for idx1, p1 := range points {
		slopes := map[string]int{}
		for idx2 := idx1 + 1; idx2 < len(points); idx2++ {
			p2 := points[idx2]
			rise, run := getSlopeOfLineBetweenPoints(p1, p2)
			slopeKey := createHashtableKeyForRational(rise, run)
			if slopes[slopeKey] == 0 {
				slopes[slopeKey] = 1
			}
			slopes[slopeKey] += 1
		}

		currentMaxNumberOfPointsOnLine := maxSlope(slopes)
		maxNumberOfPointsOnLine = max(maxNumberOfPointsOnLine, currentMaxNumberOfPointsOnLine)
	}
	return maxNumberOfPointsOnLine
}

func getSlopeOfLineBetweenPoints(p1, p2 []int) (int, int) {
	p1x, p1y := p1[0], p1[1]
	p2x, p2y := p2[0], p2[1]

	if p1x == p2x {
		return 1, 0
	}

	var xDiff = p1x - p2x
	var yDiff = p1y - p2y
	gcd := getGreatestCommonDivisor(abs(xDiff), abs(yDiff))
	xDiff = xDiff / gcd
	yDiff = yDiff / gcd
	if xDiff < 0 {
		xDiff *= -1
		yDiff *= -1
	}
	return yDiff, xDiff
}

func getGreatestCommonDivisor(num1, num2 int) int {
	a := num1
	b := num2
	for {
		if a == 0 {
			return b
		}
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
}

func createHashtableKeyForRational(numerator int, denominator int) string {
	return fmt.Sprintf("%d:%d", numerator, denominator)
}

func maxSlope(slopes map[string]int) int {
	currentMax := 0
	for _, slope := range slopes {
		currentMax = max(slope, currentMax)
	}
	return currentMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type LinePoint struct {
	x int
	y int
}

func lineThroughPoints(points [][]int) int {
	if len(points) == 1 {
		return 1
	}

	slopPoints := make(map[float64]map[LinePoint]struct{})

	for i := 0; i < len(points); i++ {
		var j = len(points)-1
		for j > i {

			point1 := LinePoint{
				points[i][0],
				points[i][1],
			}
			point2 := LinePoint{
				points[j][0],
				points[j][1],
			}

			slop := calculateSlop(point1, point2)

			if _, ok := slopPoints[slop]; !ok {
				slopPoints[slop] = make(map[LinePoint]struct{})
			}
			slopPoints[slop][point1] = struct{}{}
			slopPoints[slop][point2] = struct{}{}

			j--
		}
	}

	maxCount := math.MinInt32
	for k, v := range slopPoints {
		obliqueCuts := make(map[float64]int)
		for p := range v {
			b := float64(p.y) - k * float64(p.x)
			if k == math.MinInt32 {
				b = float64(p.x)
			}
			obliqueCuts[b]++
		}

		if len(obliqueCuts) == 1 {
			if len(v) > maxCount {
				maxCount = len(v)
			}
		} else {
			for _, c := range obliqueCuts {
				if c > maxCount {
					maxCount = c
				}
			}
		}
	}

	return maxCount
}

func calculateSlop(a, b LinePoint) float64 {
	if b.x == a.x {
		return math.MinInt32
	}

	return float64(b.y - a.y) / float64(b.x - a.x)
}
