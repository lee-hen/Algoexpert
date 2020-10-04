package main

import "fmt"

func main() {
	expected := "zab"
	output := CaesarCipherEncryptor("xyz", 2)

	fmt.Println(expected == output)
}
