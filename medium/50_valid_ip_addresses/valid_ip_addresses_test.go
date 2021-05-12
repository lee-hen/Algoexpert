package valid_ip_addresses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	input := "1921680"
	expected := []string{
		"1.9.216.80",
		"1.92.16.80",
		"1.92.168.0",
		"19.2.16.80",
		"19.2.168.0",
		"19.21.6.80",
		"19.21.68.0",
		"19.216.8.0",
		"192.1.6.80",
		"192.1.68.0",
		"192.16.8.0",
	}
	actual := ValidIPAddresses(input)
	require.ElementsMatch(t, expected, actual)
}
