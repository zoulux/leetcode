package leetcode

import "testing"

func isValid(s string) bool {
	n := len(s)
	if n%2 == 1 {
		return false
	}

	stack := []byte{}

	pairs := map[byte]byte{')': '(', ']': '[', '}': '{'}

	for j := 0; j < len(s); j++ {
		i := s[j]
		// for _, i := range s {

		if len(stack) > 0 {

			if stack[len(stack)-1] == pairs[i] {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, i)
			}

		} else {
			stack = append(stack, i)
		}
	}
	return len(stack) == 0
}
func TestIsValid(t *testing.T) {

	t.Log(isValid("()"))
	t.Log(isValid("()[]{}"))
	t.Log(isValid("(]"))
	t.Log(isValid("([)]"))
	t.Log(isValid("{[]}"))
}
