package shorten_path
// important question

import (
	"strings"
)

// ShortenPath
// O(n) time | O(n) space - where n is the length of the path
func ShortenPath(path string) string {
	startsWithSlash := path[0] == '/'
	splits := strings.Split(path, "/")
	tokens := make([]string, 0)
	for _, token := range splits {
		if isImportantToken(token) {
			tokens = append(tokens, token)
		}
	}

	stack := make([]string, 0)
	if startsWithSlash {
		stack = append(stack, "")
	}

	for _, token := range tokens {
		if token == ".." {
			if len(stack) == 0 || stack[len(stack)-1] == ".." {
				stack = append(stack, token)
			} else if stack[len(stack)-1] != "" {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, token)
		}
	}

	if len(stack) == 1 && stack[0] == "" {
		return "/"
	}
	return strings.Join(stack, "/")
}

func isImportantToken(token string) bool {
	return len(token) > 0 && token != "."
}


// my solution
// /foo/../test/../test/../foo//bar/./baz
// /foo/bar/baz

// /foo/
// /test/
// /test/
// /foo/bar/baz

//  0   1   2   3   4   5   6   7
// "/" "f" "o" "o" "/" "." "." "/" "t" "e" "s" "t" "/" "." "." "/" "t" "e" "s" "t" "/" "." "." "/" "f" "o" "o" "/" "/" "b" "a" "r" "/" "." "/" "b" "a" "z"

const root = "/"
const directorySeparator = "/"
const parentDirectory = ".."
const currentDirectory = "."

func shortenPath(path string) string {
	pathArr := strings.Split(path, "")
	directoryArr := make([]string, 0)
	for i := range pathArr {
		directoryArr = append(directoryArr, pathArr[i])
	}

	if directoryArr[len(directoryArr)-1] != directorySeparator {
		directoryArr = append(directoryArr, directorySeparator)
	}

	var symbol string
	stack := make([]string, 0)
	for _, str := range directoryArr {
		stack = append(stack, str)
		if str == directorySeparator || str == currentDirectory {
			symbol += str
		} else {
			symbol = ""
		}

		if symbol == currentDirectory + directorySeparator {
			stack = stack[:len(stack)-len(symbol)]
			symbol = ""
		}

		if symbol == directorySeparator + parentDirectory + directorySeparator {
			// "../../foo/../../
			if len(stack) >= len(symbol)+1 && stack[len(stack)-len(symbol)-1] == currentDirectory {
				// stack [. . / . . / f o o / . . /]  symbol /../
				// stack [. . / . . / . . /] 		  symbol /../
				goto jump
			}

			stack = stack[:len(stack)-len(symbol)]
			if len(stack) > 0 {
				top := stack[len(stack)-1]
				for top != directorySeparator {
					stack = stack[:len(stack)-1]
					if len(stack) == 0 {
						stack = append(stack, root)
						break
					}
					top = stack[len(stack)-1]
				}
			} else {
				stack = append(stack, root)

			}
			symbol = "/"
		}
		jump:

		if symbol == directorySeparator + currentDirectory + directorySeparator || symbol == directorySeparator + directorySeparator {
			stack = stack[:len(stack)-len(symbol)+1]
			symbol = "/"
		}
	}

	if len(stack) == 0 {
		return ""
	}

	stack = stack[:len(stack)-1]

	if len(stack) == 0 {
		if string(path[0]) == root {
			return root
		} else {
			return parentDirectory
		}
	}

	return strings.Join(stack, "")
}
