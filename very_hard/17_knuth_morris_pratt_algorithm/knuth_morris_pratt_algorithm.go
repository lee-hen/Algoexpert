package knuth_morris_pratt_algorithm
// important question

func KnuthMorrisPrattAlgorithm(str, substr string) bool {
	pattern := buildPattern(substr)
	return doesMatch(str, substr, pattern)
}

func buildPattern(substr string) []int {
	pattern := make([]int, len(substr))
	for i := range substr {
		pattern[i] = -1
	}
	i, j := 1, 0
	for i < len(substr) {
		if substr[i] == substr[j] {
			pattern[i] = j
			i++
			j++
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i++
		}
	}
	return pattern
}

func doesMatch(str, substr string, pattern []int) bool {
	i, j := 0, 0
	for i+len(substr)-j <= len(str) {
		if str[i] == substr[j] {
			if j == len(substr)-1 {
				return true
			}
			i++
			j++
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i++
		}
	}
	return false
}

// 0 1    2 3 4    5 6 7 8 9 10 11 12
// a e    f a e    d a e f a e  f  a
//   j -> j   j <- j            i
//          j                      i

//  0  1  2  3  4  5  6  7  8  9  10  11  12
// -1 -1 -1  0  1 -1  0  1  2  3   4  2    3

// 0 1 2 3 4    5 6 7 8 9 10   11 12
// a e f a e    d a e f a e    f  a
//         j -> j         j <- j
//                                j

// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24
// a e f a e f a e f a  e  d  a  e  f  a  e  d  a  e  f  a  e  f  a
//                                           i
//                                                                i

// true

func knuthMorrisPrattAlgorithm(str, substr string) bool {
	if len(str) < len(substr) {
		return false
	}

	traceIndices := make([]int, len(substr), len(substr))
	var j int
	traceIndices[j] = -1
	for i := 1; i < len(substr); i++ {
		if substr[i] == substr[j] {
			traceIndices[i] = j
			j++
		} else {
			var k = -1
			if j > 0 {
				k = i-1
				for traceIndices[k] != -1 {
					k = traceIndices[k]
					if substr[k+1] == substr[i] {
						j = k+1
						traceIndices[i] = j
						break
					}
				}
			}

			if k == -1  {
				j = 0
				traceIndices[i] = -1
			}
		}
	}

	j = 0
	for i := 0; i < len(str); i++ {
		if str[i] == substr[j] {
			if j == len(substr)-1 {
				return true
			}
			j++
		} else {
			var k = -1
			if j > 0 {
				k = j-1
				if traceIndices[k] != -1 {
					k = traceIndices[k]
				}
			}

			if k != -1 && substr[k+1] == str[i] {
				j = k+1
				j++
			} else {
				j = 0
				if substr[j] == str[i] {
					j++
				}
			}
		}
	}

	return false
}
