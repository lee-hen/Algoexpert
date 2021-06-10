package generate_document

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "Algoexpert is the Best!"
	expected := true
	actual := GenerateDocument(characters, document)
	require.Equal(t, expected, actual)
}
