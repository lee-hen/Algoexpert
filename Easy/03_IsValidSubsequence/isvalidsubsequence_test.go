package isvalidsubsequence

import (
	"testing"
)

func TestCase1(t *testing.T) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{22, 25, 6}
	IsValidSubsequence(array, sequence)
}
