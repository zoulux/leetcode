package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
	fmt.Println(minRemoveToMakeValid("a)b(c)d"))
	fmt.Println(minRemoveToMakeValid("))(("))

}

func minRemoveToMakeValid(s string) string {

	bits := make([]bool, len(s))

	var stack []int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			stack = append(stack, i)
			bits[i] = true
		case ')':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
				bits[i] = true
			}
		default:
			bits[i] = true
		}
	}

	for _, v := range stack {
		bits[v] = false
	}

	var ans strings.Builder
	for i := 0; i < len(s); i++ {
		if bits[i] {
			ans.WriteByte(s[i])
		}
	}

	return ans.String()
}
