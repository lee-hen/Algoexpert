package minimum_characters_for_words

import "math"

func MinimumCharactersForWords(words []string) []string {
	maximumCharacterFrequencies := map[rune]int{}
	for _, word := range words {
		characterFrequencies := countCharacterFrequencies(word)
		updateMaximumFrequencies(characterFrequencies, maximumCharacterFrequencies)
	}
	return makeArrayFromCharacterFrequencies(maximumCharacterFrequencies)
}

func countCharacterFrequencies(str string) map[rune]int {
	characterFrequencies := map[rune]int{}
	for _, character := range str {
		characterFrequencies[character] += 1
	}
	return characterFrequencies
}

func updateMaximumFrequencies(frequencies map[rune]int, maximumFrequencies map[rune]int) {
	for character, frequency := range frequencies {
		if maxFrequency, found := maximumFrequencies[character]; found {
			maximumFrequencies[character] = max(frequency, maxFrequency)
		} else {
			maximumFrequencies[character] = frequency
		}
	}
}

func makeArrayFromCharacterFrequencies(characterFrequencies map[rune]int) []string {
	characters := make([]string, 0)
	for character, frequency := range characterFrequencies {
		for i := 0; i < frequency; i++ {
			characters = append(characters, string(character))
		}
	}
	return characters
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// "this", "that", "did", "deed", "them!", "a"
// "t", "h", "i", "s", "t",  "a", "d", "d", "e", "e", "m", "!"

func minimumCharactersForWords(words []string) []string {
	characters := make([]map[byte]int, len(words), len(words))
	timesToSeen := make(map[byte]int)

	for i, word := range words {
		characters[i] = make(map[byte]int)
		for _, char := range []byte(word) {
			if _, saw := timesToSeen[char]; !saw {
				timesToSeen[char] = math.MinInt32
			}

			characters[i][char]++
		}
	}

	for _, character := range characters {
		for char, times := range character {
			if timesToSeen[char] < times {
				timesToSeen[char] = times
			}
		}
	}

	result := make([]string, 0)
	for char, times := range timesToSeen{
		for i := 0; i < times; i++ {
			result = append(result, string(char))
		}
	}

	return result
}
