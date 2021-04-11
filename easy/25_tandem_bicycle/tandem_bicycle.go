package tandem_bicycle

import (
	"math"
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)

	if fastest {
		sort.Sort(sort.Reverse(sort.IntSlice(blueShirtSpeeds)))
	} else {
		sort.Ints(blueShirtSpeeds)
	}

	var totalSpeed int
	for idx := 0; idx < len(redShirtSpeeds); idx++ {
		redShirtSpeed := redShirtSpeeds[idx]
		blueShirtSpeed := blueShirtSpeeds[idx]


		totalSpeed += int(math.Max(float64(redShirtSpeed), float64(blueShirtSpeed)))

	}
	return totalSpeed
}


//  5, 5, 3, 9, 2
//	3, 6, 7, 2, 1
// My not correct solution
func tandemBicycle1(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	sort.Ints(redShirtSpeeds)
	sort.Sort(sort.Reverse(sort.IntSlice(blueShirtSpeeds)))

	var totalSpeed int
	for idx := 0; idx < len(redShirtSpeeds); idx++ {
		redShirtSpeed := redShirtSpeeds[idx]
		blueShirtSpeed := blueShirtSpeeds[idx]

		if fastest {
			totalSpeed += int(math.Max(float64(redShirtSpeed), float64(blueShirtSpeed)))
		} else {
			totalSpeed += int(math.Min(float64(redShirtSpeed), float64(blueShirtSpeed)))
		}
	}
	return totalSpeed
}
