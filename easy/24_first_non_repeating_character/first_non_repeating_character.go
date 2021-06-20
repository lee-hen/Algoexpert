package first_non_repeating_character

func FirstNonRepeatingCharacter(str string) int {

	nonRepeatedChar := make(map[int32]int)
	for _, s := range str {
		nonRepeatedChar[s]++
	}

	for idx, s := range str {
		if nonRepeatedChar[s] == 1 {
			return idx
		}
	}

	return 0
}
