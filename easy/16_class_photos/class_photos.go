package class_photos

import (
	"sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Sort(sort.Reverse(sort.IntSlice(blueShirtHeights)))
	sort.Sort(sort.Reverse(sort.IntSlice(redShirtHeights)))

	var backRow string
	if redShirtHeights[0] == blueShirtHeights[0] {
		return false
	} else if redShirtHeights[0] > blueShirtHeights[0] {
		backRow = "Red"
	} else {
		backRow = "Blue"
	}

	for i, redShirtHeight := range redShirtHeights {
		if backRow == "Red" {
			if blueShirtHeights[i] >= redShirtHeight {
				return false
			}
		} else {
			if blueShirtHeights[i] <= redShirtHeight {
				return false
			}
		}
	}

	return true
}
