package productsum

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	input := SpecialArray{
		5, 2,
		SpecialArray{7, -1},
		3,
		SpecialArray{
			6,
			SpecialArray{
				-13, 8,
			},
			4,
		},
	}
	output := ProductSum(input)
	expected := 12
	fmt.Println(output)
	fmt.Println(expected == output)
}
