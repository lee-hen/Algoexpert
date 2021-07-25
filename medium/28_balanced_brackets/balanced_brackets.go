package balanced_brackets

var opening = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
}

var closing = map[rune]bool{
	')': true,
	']': true,
	'}': true,
}

var matching = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

// BalancedBrackets
// ([])(){}(())()()
func BalancedBrackets(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		if _, found := opening[char]; found {
			stack = append(stack, char)
			continue
		}
		if _, found := closing[char]; found {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == matching[char] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}


// BalancedBrackets
// ([])(){}(())()()
// my solution
func balancedBrackets(s string) bool {
	var brackets []rune
	for _, r := range s {
		var removed bool
		if len(brackets) != 0 {
			open := string(brackets[len(brackets)-1])
			if string(r) == "]" && open == "[" {
				brackets = brackets[:len(brackets)-1]
				removed = true
			} else if string(r) == ")" && open == "(" {
				brackets = brackets[:len(brackets)-1]
				removed = true
			} else if string(r) == "}" && open == "{" {
				brackets = brackets[:len(brackets)-1]
				removed = true
			}
		}

		if removed == true {
			continue
		}

		if string(r) == "(" || string(r) == ")" ||
			string(r) == "["  || string(r) == "]"  ||
			string(r) == "{" || string(r) == "}" {
			brackets = append(brackets, r)
		}
	}

	return len(brackets) == 0
}

// if no other characters
//func isBalanced(str string) bool {
//	openParensStack := make([]rune, 0)
//	for _, char := range str {
//		if char == '(' || char == '[' || char == '{' {
//			openParensStack = append(openParensStack, char)
//		} else if len(openParensStack) > 0  {
//			openParensStack = openParensStack[:len(openParensStack)-1]
//		} else {
//			return false
//		}
//	}
//	return len(openParensStack) == 0
//}
