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
