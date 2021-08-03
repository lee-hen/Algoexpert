package glob_matching

// GlobMatching
// O(n * m) time | O(n * m) space - where n is the length
// of the fileName and m is the length of the pattern.
func GlobMatching(fileName, pattern string) bool {
	matchTable := initializeMatchTable(fileName, pattern)
	for i := 1; i <= len(fileName); i++ {
		for j := 1; j <= len(pattern); j++ {
			if pattern[j-1] == '*' {
				matchTable[i][j] = matchTable[i][j-1] || matchTable[i-1][j]
			} else if pattern[j-1] == '?' || pattern[j-1] == fileName[i-1] {
				matchTable[i][j] = matchTable[i-1][j-1]
			}
		}
	}
	return matchTable[len(fileName)][len(pattern)]
}

func initializeMatchTable(filename, pattern string) [][]bool {
	matchTable := make([][]bool, len(filename)+1)
	for i := 0; i < len(filename)+1; i++ {
		matchTable[i] = make([]bool, len(pattern)+1)
	}

	matchTable[0][0] = true
	for j := 1; j <= len(pattern); j++ {
		if pattern[j-1] == '*' {
			matchTable[0][j] = matchTable[0][j-1]
		}
	}

	return matchTable
}

func globMatching(fileName, pattern string) bool {
	patternMatching := make([][]int, len(fileName)+1, len(fileName)+1)
	for i := range patternMatching {
		patternMatching[i] = make([]int, len(pattern)+1, len(pattern)+1)
	}
	patternMatching[0][0] = 1

	for j := range pattern {
		if pattern[j] == '*' {
			patternMatching[0][j+1] = patternMatching[0][j]
		}
	}

	for i := range fileName {
		for j := range pattern {
			if fileName[i] == pattern[j] || pattern[j] == '?' {
				patternMatching[i+1][j+1] = patternMatching[i][j]
			} else if pattern[j] == '*' {
				if patternMatching[i][j+1] == 1 || patternMatching[i+1][j] == 1 {
					patternMatching[i+1][j+1] = 1
				}
			} else  {
				patternMatching[i+1][j+1] = 0
			}
		}
	}
	return patternMatching[len(fileName)][len(pattern)] == 1
}
