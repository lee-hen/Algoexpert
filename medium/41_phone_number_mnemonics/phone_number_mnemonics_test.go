package phone_number_mnemonics

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	phoneNumber := "1905"
	expected := []string{
		"1w0j",
		"1w0k",
		"1w0l",
		"1x0j",
		"1x0k",
		"1x0l",
		"1y0j",
		"1y0k",
		"1y0l",
		"1z0j",
		"1z0k",
		"1z0l",
	}
	actual := PhoneNumberMnemonics(phoneNumber)
	require.Equal(t, expected, actual)
}

