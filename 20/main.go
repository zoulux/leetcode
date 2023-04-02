package main

func main() {

}

func isValid(s string) bool {

	stack := make([]rune, 0)

	//s 仅由括号 '(])[]{}' 组成

	for _, v := range s {
		switch v {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack = append(stack, v)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}

		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
