package main

import "fmt"

func main() {
	expected := true
	output := IsPalindrome("abcdcba")

	fmt.Println(expected == output)
}
