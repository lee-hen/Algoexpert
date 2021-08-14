package valid_ip_addresses
// important question

import (
	"strconv"
	"strings"
)

func ValidIPAddresses(str string) []string {
	ipAddressesFound := make([]string, 0)
	for i := 1; i < min(len(str), 4); i++ {
		currentIPAddressParts := []string{"", "", "", ""}
		currentIPAddressParts[0] = str[:i]
		if !isValidPart(currentIPAddressParts[0]) {
			continue
		}
		for j := i + 1; j < i+min(len(str)-i, 4); j++ {
			currentIPAddressParts[1] = str[i:j]
			if !isValidPart(currentIPAddressParts[1]) {
				continue
			}
			for k := j + 1; k < j+min(len(str)-j, 4); k++ {
				currentIPAddressParts[2] = str[j:k]
				currentIPAddressParts[3] = str[k:]

				if isValidPart(currentIPAddressParts[2]) && isValidPart(currentIPAddressParts[3]) {
					ipAddressesFound = append(ipAddressesFound, strings.Join(currentIPAddressParts, "."))
				}
			}
		}
	}
	return ipAddressesFound
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func isValidPart(str string) bool {
	i, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if i > 255 {
		return false
	}
	return len(str) == len(strconv.Itoa(i))
}

// str length 12 or smaller
// insert .
// separated by . four positive integer between 0-255
// my solution
func validIPAddresses(str string) []string {
	if len(str) < 4 {
		return []string{}
	}
	ips := validIPAddressesHelper(1, 2, 3, []byte(str))
	if ips == nil {
		return []string{}
	}
	return ips
}

func validIPAddressesHelper(first, second, third int, bytes []byte) (ips []string) {
	if third > len(bytes)-1 {
		return
	}

	if second >= third {
		return
	}

	if first >= second {
		return
	}

	partFour := bytes[third:]
	partThree := bytes[second:third]
	partTwo := bytes[first:second]
	partOne := bytes[:first]

	if validPartOfIp(partOne) && validPartOfIp(partTwo) &&
		validPartOfIp(partThree) && validPartOfIp(partFour) {
		ip := make([]byte, 0)
		ip = append(ip, partOne...)
		ip = append(ip, '.')
		ip = append(ip, partTwo...)
		ip = append(ip, '.')
		ip = append(ip, partThree...)
		ip = append(ip, '.')
		ip = append(ip, partFour...)
		ips = append(ips, string(ip))
	}

	if first == 1 && second == 2 {
		ips = append(ips, validIPAddressesHelper(first, second, third+1, bytes)...)
	}
	if first == 1 {
		ips = append(ips, validIPAddressesHelper(first, second+1, third, bytes)...)
	}
	return append(ips, validIPAddressesHelper(first+1, second, third, bytes)...)
}

func validPartOfIp(section []byte) bool{
	if len(section) > 1 && section[0] == '0'{
		return false
	}

	number, _ := strconv.Atoi(string(section))
	if number < 0 || number > 255 {
		return false
	}

	return true
}
