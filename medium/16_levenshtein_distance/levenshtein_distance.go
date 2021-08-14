package levenshtein_distance
// important question

import "math"

// d -> a b c x
// delete d add a b c x         5
// add a b c do not replace d   4
// replace d with a add b c z   4

// "" x -> "" a
// "" x -> ""    del x
// "" x -> "" a insert

// 0 1 2 3 4
//"" a b c x  str2/str1
// 0 1 2 3 4 [""]         0
// 1 1 2 3 3 ["" x]       1
// 2 1 2 3 4 ["" x a]     2
// 3 2 1 2 3 ["" x a b]   3
// 4 3 2 1 2 ["" x a b c] 4

// not equal
// insert/deletion
// str1[row-1][col] + 1

// substitution
// str1[row-1][col-1] + 1

// delete/insertion
// str1[row][col-1] + 1

// equal
// do nothing => str1[row-1][col-1]

// x -> a         1
// x -> a b       2
// x -> a b c     3
// x -> a b c x   3

// x a -> a       1
// x a -> a b     2
// x a -> a b c   3
// x a -> a b c x 4

// x a b-> a       2
// x a b-> a b     1
// x a b-> a b c   2
// x a b-> a b c x 3

// x a b c-> a         3
// x a b c-> a b       2
// x a b c-> a b c     1
// x a b c-> a b c x   2

// LevenshteinDistance
// x, a, b, c →　a, b, c, x
// ?? why?
func LevenshteinDistance(a, b string) int {
	if b == "" {
		return len(a)
	}

	if a == "" {
		return len(b)
	}

	str1 := make([]rune, 0)
	str2 := make([]rune, 0)

	str1 = append(str1, []rune(" ")...)
	for _, ch := range a {
		str1 = append(str1, ch)
	}

	str2 = append(str2, []rune(" ")...)
	for _, ch := range b {
		str2 = append(str2, ch)
	}

	distances := make([][]int, 2, 2)
	distances[0] = make([]int, len(str2), len(str2))

	for idx := range str2 {
		distances[0][idx] = idx
	}

	distances[1] = make([]int, len(str2), len(str2))
	distances[1][0]++

	for row, ch1 := range str1 {
		if row == 0 {
			continue
		}
		for col, ch2 := range str2 {
			if col > 0 {
				if ch1 != ch2 {
					distances[1][col] = min(distances[0][col]+1, distances[0][col-1]+1, distances[1][col-1]+1)
				} else {
					distances[1][col] = distances[0][col-1]
				}
			}
		}

		for idx, d := range distances[1] {
			distances[0][idx] = d
		}
		distances[1][0]++
	}

	return distances[len(distances)-1][len(str2)-1]
}

func levenshteinDistance(a, b string) int {
	small, big := a, b
	if len(a) > len(b) {
		big, small = small, big
	}
	evenEdits := make([]int, len(small)+1)
	oddEdits := make([]int, len(small)+1)
	var previousEdits, currentEdits []int
	for i := range evenEdits {
		evenEdits[i] = i
		oddEdits[i] = math.MinInt32
	}
	for i := 1; i < len(big)+1; i++ {
		if i%2 == 1 {
			currentEdits, previousEdits = oddEdits, evenEdits
		} else {
			currentEdits, previousEdits = evenEdits, oddEdits
		}
		currentEdits[0] = i
		for j := 1; j < len(small)+1; j++ {
			if big[i-1] == small[j-1] {
				currentEdits[j] = previousEdits[j-1]
			} else {
				currentEdits[j] = 1 + min(previousEdits[j-1], previousEdits[j], currentEdits[j-1])
			}
		}
	}
	if len(big)%2 == 0 {
		return evenEdits[len(small)]
	}
	return oddEdits[len(small)]
}

func min(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num < curr {
			curr = num
		}
	}
	return curr
}
