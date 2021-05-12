package valid_ip_addresses

// str length 12 or smaller
// insert .

// separated by . four positive integer between 0-255

//0.0.0.0
//255.255.255.255

// 0. 10. 2. 3

// ValidIPAddresses
// 1921680
//"1.9.216.80",
//"1.92.16.80",
//"1.92.168.0",
//"19.2.16.80",
//"19.2.168.0",
//"19.21.6.80",
//"19.21.68.0",
//"19.216.8.0",
//"192.1.6.80",
//"192.1.68.0",
//"192.16.8.0",
func ValidIPAddresses(str string) []string {
	if len(str) < 4 {
		return []string{}
	}
	return validIPAddressesHelper(0, []byte{}, []byte(str))
}

func validIPAddressesHelper(idx int, chars, bytes []byte) (ips []string) {
	for i := idx; i < len(bytes); i++ {
		chars = append(chars, bytes[i])
		validIPAddressesHelper(idx+1, chars, bytes)
	}

	return
}
