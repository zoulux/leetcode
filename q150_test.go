package leetcode

import (
	"strconv"
	"testing"
)

func evalRPN(tokens []string) int {

	stack := []int{}

	for _, token := range tokens {

		if token == "+" || token == "-" || token == "*" || token == "/" {
			num2 := stack[len(stack)-1]
			num1 := stack[len(stack)-2]

			stack = stack[:len(stack)-2]

			ret := 0

			switch token {
			case "+":
				ret = num1 + num2
			case "-":
				ret = num1 - num2
			case "*":
				ret = num1 * num2
			case "/":
				ret = num1 / num2
			}
			stack = append(stack, ret)
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}

func TestEvalRPN(t *testing.T) {
	// t.Log(evalRPN([]string{"2", "1", "+", "3", "*"}))
	t.Log(evalRPN([]string{"4", "13", "5", "/", "+"}))
	t.Log(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
