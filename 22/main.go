package main

import "fmt"

func main() {

	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))

}

func generateParenthesis(n int) []string {
	var ans []string
	var travel func(leftN int, rightN int, s string)
	travel = func(leftN int, rightN int, s string) {
		if leftN == rightN && rightN == n {
			// 如果 左括号等于右括号，并且数量等于 n
			ans = append(ans, s)
			return
		}

		if leftN > n || rightN > n {
			return
		}

		travel(leftN+1, rightN, s+"(")
		// 为了保证括号是有效的，那么必须是左括号比右括号是多的，可以纸上画一画
		if leftN > rightN {
			travel(leftN, rightN+1, s+")")
		}
	}
	travel(0, 0, "")
	return ans
}
