package caesar_cipher_encryptor

func CaesarCipherEncryptor(str string, key int) string {
	key = key % 26
	var buff []rune
	for i := range str {
		charCode := int(str[i]) + key
		if charCode > 122 {
			charCode = 96 + charCode%122
		}
		buff = append(buff, rune(charCode))
	}

	return string(buff)
}

// caesarCipherEncryptor
// My solution
func caesarCipherEncryptor(str string, key int) string {
	key = key % 26
	var buff []byte
	for i := range str {
		char := int(str[i]) + key
		diff := char - 122
		if diff > 0 {
			char = 96 + diff
		}
		buff = append(buff, byte(char))
	}

	return string(buff)
}
