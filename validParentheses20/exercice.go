package validParentheses20

import "fmt"

func Run() {
	fmt.Println(isValid("()") == true)
	fmt.Println(isValid("()[]{}") == true)
	fmt.Println(isValid("(]") == false)
	fmt.Println(isValid("([])") == true)
}

func isValid(s string) bool {
	stack := make([]uint8, 0)

	for i := 0; i < len(s); i++ {
		if len(stack) > 0 {
			if s[i] == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if s[i] == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			if s[i] == ')' && stack[len(stack)-1] != '(' {
				return false
			}
		} else if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			return false
		}

		if s[i] == ')' || s[i] == '}' || s[i] == ']' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack) == 0
}
