package phone_number_mnemonics
// important question

import (
	"strconv"
)

//
//----- ----- -----
//|     |     |     |
//|  1  |  2  |  3  |
//|     | abc | def |
//----- ----- -----
//|     |     |     |
//|  4  |  5  |  6  |
//| ghi | jkl | mno |
//----- ----- -----
//|     |     |     |
//|  7  |  8  |  9  |
//| pqrs| tuv | wxyz|
//----- ----- -----
//|     |
//|  0  |
//|     |
//-----

func PhoneNumberMnemonics(phoneNumber string) []string {
	letters := []byte{
		byte('a'), byte('b'), byte('c'),
		byte('d'), byte('e'), byte('f'),
		byte('g'), byte('h'), byte('i'),
		byte('j'), byte('k'), byte('l'),
		byte('m'), byte('n'), byte('o'),
		byte('p'), byte('q'), byte('r'), byte('s'),
		byte('t'), byte('u'), byte('v'),
		byte('w'), byte('x'), byte('y'), byte('z'),
	}

	mnemonics := make(map[int][]byte)
	var j int
	for i := 0; i < 10; i++ {
		if i == 0  {
			mnemonics[i] = []byte{byte('0')}
		} else if i == 1 {
			mnemonics[i] = []byte{byte('1')}
		} else if i < 7 || i == 8 {
			mnemonics[i] = letters[j:j+3]
			j = j+3
		} else if i == 7 || i == 9 {
			mnemonics[i] = letters[j:j+4]
			j = j+4
		}
	}

	phoneNumberMnemonic := make([]byte, len(phoneNumber))
	return generatePhoneNumberMnemonics(0,phoneNumber, phoneNumberMnemonic, mnemonics)
}

func generatePhoneNumberMnemonics(i int, phoneNumber string, phoneNumberMnemonic []byte, mnemonics map[int][]byte) (phoneNumberMnemonics []string) {
	if i >= len(phoneNumber) {
		phoneNumberMnemonics = append(phoneNumberMnemonics, string(phoneNumberMnemonic))
		return
	}

	numberStr := []rune(phoneNumber)[i]
	number, _ := strconv.Atoi(string(numberStr))

	for _, letter := range mnemonics[number] {
		phoneNumberMnemonic[i] = letter
		phoneNumberMnemonics = append(phoneNumberMnemonics, generatePhoneNumberMnemonics(i+1, phoneNumber, phoneNumberMnemonic, mnemonics)...)

	}
	return
}
