package caesarcipherencryptor

import (
	"fmt"
	"testing"
)

func TestCase1(t *testing.T) {
	expected := "zab"
	output := CaesarCipherEncryptor("xyz", 2)

	fmt.Println(expected == output)
}
