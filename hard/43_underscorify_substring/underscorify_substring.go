package underscorify_substring
// important question

import (
	"strings"
)

type intervals []*interval

type interval struct {
	left  int
	right int
}

// UnderscorifySubstring
// Average case: O(n + m) | O(n) space - where n is the length
// of the main string and m is the length of the substring
func UnderscorifySubstring(str string, substring string) string {
	locations := getLocations(str, substring)
	locations = locations.collapse()
	return underscorify(str, locations)
}

func getLocations(str, substring string) intervals {
	result := intervals{}
	for start := 0; start < len(str); {
		nextIndex := strings.Index(str[start:], substring)
		if nextIndex == -1 {
			break
		}

		nextIndex += start
		result = append(result, &interval{nextIndex, nextIndex + len(substring)})
		start = nextIndex + 1
	}
	return result
}

func (array intervals) collapse() intervals {
	// If the array is empty, nothing to do
	if len(array) == 0 {
		return array
	}

	result := intervals{array[0]}
	previous := array[0]
	for i := 1; i < len(array); i++ {
		current := array[i]
		if current.left <= previous.right {
			// Collapse the two intervals
			previous.right = current.right
		} else {
			result = append(result, current)
			previous = current
		}
	}
	return result
}

func underscorify(str string, locations intervals) string {
	if len(locations) == 0 {
		return str
	}

	// We know the resulting string will have an additional 2*len(intervals)
	// characters
	result := make([]rune, len(str)+2*len(locations))
	resultIndex := 0
	locationIndex := 0
	for i, r := range str {
		location := locations[locationIndex]
		if i == location.left {
			result[resultIndex] = '_'
			resultIndex += 1
		} else if i == location.right {
			result[resultIndex] = '_'
			resultIndex += 1
			if locationIndex+1 < len(locations) {
				locationIndex += 1
			}
		}
		result[resultIndex] = r
		resultIndex += 1
	}

	if locations[locationIndex].right == len(str) {
		result[len(result)-1] = '_'
	}
	return string(result)
}

// underscorifySubstring
// my solution
// testthis is a testtest to see if testestest it works
// _test_this is a _testtest_ to see if _testestest_ it works
func underscorifySubstring(str string, substring string) string {
	strs := strings.Split(str, substring)
	sep := "_" + substring + "_"
	newStr := strings.Join(strs, sep)

	substrings := strings.Split(substring, "")

	// _abc__abc_ -> c__a -> ca
	// t__t -> tt
	sep1 := substrings[len(substrings)-1] + "__"  + substrings[0]

	var center string
	if len(substrings) > 1 {
		center = strings.Join(substrings[1:len(substrings)-1], "")
	}

	// t_es_t -> test
	sep2 := substrings[len(substrings)-1] + "_" + center + "_" + substrings[0]

	// t_est -> test_
	sep3 := substrings[len(substrings)-1] + "_" + center + substrings[0]
	newStrs := strings.Split(newStr, " ")

	for idx, word := range newStrs {
		if word == sep1 || word == sep2 || word == sep3 {
			continue
		}

		// sep1
		newStrs[idx] = strings.Replace(word, sep1, substrings[len(substrings)-1] + substrings[0], -1)

		// sep2
		if center == "" {
			// a__a -> aa
			newStrs[idx] = strings.Replace(newStrs[idx], sep2, substring+substring, -1)
		} else {
			newStrs[idx] = strings.Replace(newStrs[idx], sep2, substring, -1)
		}

		// sep3
		newStrs[idx] = strings.Replace(newStrs[idx], sep3, substring + "_", -1)
	}

	return strings.Join(newStrs, " ")
}
