package longest_common_subsequence
// important question

// LongestCommonSubsequence
// O(nm*min(n, m)) time | O(nm*min(n, m)) space
// ?? why?
func LongestCommonSubsequence(s1 string, s2 string) []byte {
	lcs := make([][][]byte, len(s2)+1)
	for i := range lcs {
		lcs[i] = make([][]byte, len(s1)+1)
	}

	for i := 1; i < len(lcs); i++ {
		for j := 1; j < len(lcs[i]); j++ {
			if s2[i-1] == s1[j-1] {
				tmp := make([]byte, len(lcs[i-1][j-1]))
				copy(tmp, lcs[i-1][j-1])
				lcs[i][j] = append(tmp, s2[i-1])
			} else {
				if len(lcs[i-1][j]) < len(lcs[i][j-1]) {
					lcs[i][j] = lcs[i][j-1]
				} else {
					lcs[i][j] = lcs[i-1][j]
				}
			}
		}
	}

	return lcs[len(s2)][len(s1)]
}

// Z X V V Y Z W        s1
// X K Y K Z P W        s2

//  0   0   1   2   3   4   5   6
//      X   K   Y   K   Z   P   W    s2/s1
//      0   0   0   0   0   0   0       0
//  0   0   0   0   0   1   1   1    Z  0
//  0   1   1   1   1   1   1   1    X  1
//  0   1   1   1   1   1   1   1    V  2
//  0   1   1   1   1   1   1   1    V  3
//  0   1   1   2   2   2   2   2    Y  4
//  0   1   1   2   2   3   3   3    Z  5
//  0   1   1   2   2   3   3   4    W  6

// 0 1 2 i
// A B C D       str(i)
// B A D A       str(j)
// 0 1 2 j

// when str(i) != str(j)
// we need to compare [A B C D] [B A D] LCS to [A B C] [B A D A] LCS
// so the answer is either LCS(i)(j-1) or LCS(i-1)(j)

// 0 1 2 i
// A B C D       str(i)
// A C B D       str(j)
// 0 1 2 j

// when str(i) == str(j)
// when need to pick [A B C] [A C B] LCS which is 2
// LCS[i][j] = LCS[i-1][j-1]+1

func longestCommonSubsequence(s1 string, s2 string) []byte {
	chars1 := []byte(s1)
	chars2 := []byte(s2)

	commonSubsequenceLens := make([][]int, len(chars1)+1, len(chars1)+1)
	for i := 0; i < len(commonSubsequenceLens); i++ {
		commonSubsequenceLens[i] = make([]int, len(chars2)+1, len(chars2)+1)
	}

	for i := 1; i <= len(chars1); i++  {
		char1 := chars1[i-1]
		for j := 1; j <= len(chars2); j++ {
			char2 := chars2[j-1]
			if char1 == char2 {
				commonSubsequenceLens[i][j] = commonSubsequenceLens[i-1][j-1]+1
			} else {
				commonSubsequenceLens[i][j] = max(commonSubsequenceLens[i-1][j], commonSubsequenceLens[i][j-1])
			}
		}
	}

	return buildLongestCommonSubsequences(commonSubsequenceLens, chars2)
}

func buildLongestCommonSubsequences(commonSubsequenceLens [][]int, chars2 []byte) []byte {
	longestCommonSubsequences := make([]byte, 0)

	i := len(commonSubsequenceLens)-1
	j := len(commonSubsequenceLens[len(commonSubsequenceLens)-1])-1

	for i > 0 && j > 0 {
		if commonSubsequenceLens[i][j] == commonSubsequenceLens[i-1][j] {
			i--
		} else if commonSubsequenceLens[i][j] == commonSubsequenceLens[i][j-1] {
			j--
		} else {
			longestCommonSubsequences = append(longestCommonSubsequences, chars2[j-1])
			i--
			j--
		}
	}

	ReverseSlice(longestCommonSubsequences)
	return longestCommonSubsequences
}

func ReverseSlice(slice []byte) {
	start := 0
	end := len(slice) -1

	for start < end {
		slice[start], slice[end] = slice[end], slice[start]
		start += 1
		end -= 1
	}
}

func max(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}
