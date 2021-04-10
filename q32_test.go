package leetcode

import "testing"

func longestValidParentheses(s string) int {

	n := len(s)
	dp := make([]int, n)
	max := 0

	for i := 1; i < n; i++ {
		c := s[i]
		lc := s[i-1]
		if c == ')' {
			if lc == '(' {
				if i >= 2 {
					dp[i] = 2 + dp[i-2]
				} else {
					dp[i] = 2
				}

			} else if lc == ')' {
				if i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
					dp[i] = 2 + dp[i-1]
					if i-dp[i-1]-2 >= 0 {
						dp[i] = dp[i] + dp[i-dp[i-1]-2]
					}
				}
			}
		}

		if dp[i] > max {
			max = dp[i]
		}

	}
	return max
}

// "(()"
// ")()())"
// ""

func TestLongestValidParentheses(t *testing.T) {
	// t.Log(longestValidParentheses("(()"))
	// t.Log(longestValidParentheses(")()()"))
	// t.Log(longestValidParentheses(""))
	// t.Log(longestValidParentheses("()(()"))
	// t.Log(longestValidParentheses(")()())"))
	t.Log(longestValidParentheses("()(())"))

}
