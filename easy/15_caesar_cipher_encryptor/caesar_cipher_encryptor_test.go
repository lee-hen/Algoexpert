package caesar_cipher_encryptor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCase1(t *testing.T) {
	expected := "zab"
	output := CaesarCipherEncryptor("xyz", 2)
	require.Equal(t, expected, output)
}
