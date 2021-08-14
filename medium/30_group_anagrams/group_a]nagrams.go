package group_anagrams
// important question

import "sort"

func GroupAnagrams(words []string) [][]string {
	anagrams := map[string][]string{}
	for _, word := range words {
		sortedWord := sortWord(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}
	result := make([][]string, 0)
	for _, group := range anagrams {
		result = append(result, group)
	}
	return result
}

func GroupAnagrams1(words []string) [][]string {
	if len(words) == 0 {
		return [][]string{}
	}
	sortedWords := make([]string, 0)
	indices := make([]int, 0)

	for i, word := range words {
		sortedWords = append(sortedWords, sortWord(word))
		indices = append(indices, i)
	}

	sort.Slice(indices, func(i, j int) bool {
		return sortedWords[indices[i]] < sortedWords[indices[j]]
	})

	result := make([][]string, 0)
	currentAnagramGroup := make([]string, 0)
	currentAnagram := sortedWords[indices[0]]
	for _, index := range indices {
		word := words[index]
		sortedWord := sortedWords[index]
		if len(currentAnagramGroup) == 0 {
			currentAnagramGroup = append(currentAnagramGroup, word)
			currentAnagram = sortedWord
			continue
		}
		if sortedWord == currentAnagram {
			currentAnagramGroup = append(currentAnagramGroup, word)
			continue
		}
		result = append(result, currentAnagramGroup)
		currentAnagramGroup = []string{word}
		currentAnagram = sortedWord
	}

	result = append(result, currentAnagramGroup)
	return result
}

func sortWord(word string) string {
	wordBytes := []byte(word)
	sort.Slice(wordBytes, func(i, j int) bool {
		return wordBytes[i] < wordBytes[j]
	})
	return string(wordBytes)
}

func groupAnagrams(words []string) (anagrams [][]string) {
	wordsMap := make(map[int][]string)

	for _, word := range words {
		wordsMap[len(word)] = append(wordsMap[len(word)], word)
	}

	for _, wordsArr := range wordsMap {
		remainders := getAnagrams(wordsArr, &anagrams)
		for _, remainder := range remainders {
			anagrams = append(anagrams, []string{remainder})
		}
	}
	if anagrams == nil {
		return [][]string{}
	}

	return
}

func getAnagrams(wordsArr []string, anagrams * [][]string) (remainders []string) {
	if len(wordsArr) == 0 {
		return
	}

	word := wordsArr[len(wordsArr)-1]
	wordsArr = wordsArr[:len(wordsArr)-1]

	var anagram []string
	for _, wordStr := range wordsArr {
		if isAnagram(word, wordStr) {
			anagram = append(anagram, wordStr)
		}
	}

	if len(anagram) > 0 {
		anagram = append(anagram, word)
		*anagrams = append(*anagrams, anagram)
		wordsArr = difference(wordsArr, anagram[:len(anagram)-1])
	} else {
		remainders = append(remainders, word)
	}

	remainders = append(remainders, getAnagrams(wordsArr, anagrams)...)
	return
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string

	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}

func isAnagram(str1, str2 string) bool {
	rune1 := chars(str1)
	rune2 := chars(str2)

	sort.Sort(rune1)
	sort.Sort(rune2)

	return string(rune1) == string(rune2)
}

type chars []rune

func (c chars) Len() int {
	return len(c)
}

func (c chars) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c chars) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
