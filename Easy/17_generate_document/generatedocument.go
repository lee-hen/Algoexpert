package generate_document

func GenerateDocument(characters string, document string) bool {
	if document == "" {
		return true
	}

	charCount := make(map[rune]int)
 	for _, character := range characters {
		charCount[character] += 1
	}

	for _, character := range document {
		count, ok := charCount[character]
		if !ok {
			return false
		}
		charCount[character] = count - 1
		if charCount[character] < 0 {
			return false
		}
	}

	return true
}
