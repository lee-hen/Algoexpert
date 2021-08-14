package generate_div_tags
// important question

import (
	"strings"
)

const openingTag = "<div>"
const closingTag = "</div>"

// GenerateDivTags
// O((2n)!/((n!((n + 1)!)))) time | O((2n)!/((n!((n + 1)!)))) space -
// where n is the input number
func GenerateDivTags(numberOfTags int) []string {
	matchedDivTags := make([]string, 0)
	generateDivTagsFromPrefix(numberOfTags, numberOfTags, "", &matchedDivTags)
	return matchedDivTags
}

func generateDivTagsFromPrefix(openingTagsNeeded int, closingTagsNeeded int, prefix string, result *[]string) {
	if openingTagsNeeded > 0 {
		newPrefix := prefix + openingTag
		generateDivTagsFromPrefix(openingTagsNeeded-1, closingTagsNeeded, newPrefix, result)
	}

	// when openingTagsNeeded is 1 there are two open tags have been added to the prefix
	// then the recursive function will add one closing tag and add one open tag
	// finally add two closing tag and so on
	// similar of closingTagsNeeded when closingTagsNeeded is 2 there are 1 closing tag has been added to the prefix
	// because the openingTagsNeeded is 1 so the same process as openingTagsNeeded
	if openingTagsNeeded < closingTagsNeeded {
		newPrefix := prefix + closingTag
		generateDivTagsFromPrefix(openingTagsNeeded, closingTagsNeeded-1, newPrefix, result)
	}

	if closingTagsNeeded == 0 {
		*result = append(*result, prefix)
	}
}

// 3
// "<div> <div> <div> </div> </div> </div>"
// "<div> <div> </div> <div> </div> </div>"
// "<div> <div> </div> </div> <div> </div>"
// "<div> </div> <div> <div> </div> </div>"
// "<div> </div> <div> </div> <div> </div>"
//

func generateDivTags(numberOfTags int) []string {
	return generateDivTagsHelper(numberOfTags, []string{})
}

func generateDivTagsHelper(numberOfTags int, partialWithExtraOpeningTag []string) []string {
	finalDivTags := make([]string, 0)

	if len(partialWithExtraOpeningTag) == numberOfTags * 2 {
		finalDivTag := strings.Join(partialWithExtraOpeningTag, "")
		finalDivTags = append(finalDivTags, finalDivTag)
		return finalDivTags
	}

	if numberOfOpenTags(partialWithExtraOpeningTag) < numberOfTags {
		partialWithExtraOpeningTag = append(partialWithExtraOpeningTag, openingTag)
		finalDivTags = append(finalDivTags, generateDivTagsHelper(numberOfTags, partialWithExtraOpeningTag)...)

		for numberOfClosingTags(partialWithExtraOpeningTag) < numberOfOpenTags(partialWithExtraOpeningTag) {
			partialWithExtraOpeningTag = append(partialWithExtraOpeningTag, closingTag)
			finalDivTags = append(finalDivTags, generateDivTagsHelper(numberOfTags, partialWithExtraOpeningTag)...)
		}
	}

	return finalDivTags
}

func numberOfOpenTags(partialWithExtraOpeningTag []string) int {
	var numberOfOpenTags int
	for _, tag := range partialWithExtraOpeningTag {
		if tag == openingTag {
			numberOfOpenTags++
		}
	}
	return numberOfOpenTags
}

func numberOfClosingTags(partialWithExtraOpeningTag []string) int {
	var numberOfClosingTags int
	for _, tag := range partialWithExtraOpeningTag {
		if tag == closingTag {
			numberOfClosingTags++
		}
	}
	return numberOfClosingTags
}
