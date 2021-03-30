package ispalindrome

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := true
	output := IsPalindrome("abcdcba")

	fmt.Println(expected == output)
}
