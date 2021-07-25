package smallest_substring_containing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	bigString := "abcd$ef$axb$c$"
	smallString := "$$abf"
	expected := "f$axb$"
	require.Equal(t, expected,
		SmallestSubstringContaining(bigString, smallString),
	)
}

func TestCase2(t *testing.T) {
	bigString := "abcdefghijklmnopqrstuvwxyz"
	smallString := "aajjttwwxxzz"
	expected := ""
	require.Equal(t, expected,
		SmallestSubstringContaining(bigString, smallString),
	)
}

func TestCase3(t *testing.T) {
	bigString := "abzacdwejxjftghiwjtklmnopqrstuvwxyz"
	smallString := "aajjttwwxxzz"
	expected := "abzacdwejxjftghiwjtklmnopqrstuvwxyz"
	require.Equal(t,expected,
		SmallestSubstringContaining(bigString, smallString),
	)
}

func TestCase4(t *testing.T) {
	bigString := "a$fuu+afff+affaffa+a$Affab+a+a+$a$"
	smallString := "a+$aaAaaaa$++"
	expected := "affa+a$Affab+a+a+$a"
	require.Equal(t,
		expected, SmallestSubstringContaining(bigString, smallString),
	)
}

func TestCase5(t *testing.T) {
	bigString := "1456243561288281932363365412356789901!"
	smallString := "123!"
	expected := "2356789901!"
	require.Equal(t,
		expected, SmallestSubstringContaining(bigString, smallString),
	)
}

func TestCase6(t *testing.T) {
	bigString := "14562435612!88281932363365$412356789901"
	smallString := "$123!"
	expected := "!88281932363365$"
	require.Equal(t,
		expected, SmallestSubstringContaining(bigString, smallString),
	)
}

func TestCase7(t *testing.T) {
	bigString := "14562435612z!8828!193236!336!5$41!23!5!6789901#"
	smallString := "#!2z"
	expected := "z!8828!193236!336!5$41!23!5!6789901#"
	require.Equal(t,
		expected, SmallestSubstringContaining(bigString, smallString),
	)
}
